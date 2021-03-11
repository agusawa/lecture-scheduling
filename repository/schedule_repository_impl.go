package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"lecture-scheduling/entity"
	"lecture-scheduling/exception"
	"time"
)

func NewScheduleRepository(connection *sql.DB) ScheduleRepository {
	return &scheduleRepositoryImpl{
		connection: connection,
	}
}

type scheduleRepositoryImpl struct {
	connection *sql.DB
}

func (repository *scheduleRepositoryImpl) Add(schedule *entity.Schedule) {
	query := "INSERT INTO schedules (code, name, start_time, end_time, lecturer_name, day) VALUES (?, ?, ?, ?, ?, ?)"

	statement, err := repository.connection.Prepare(query)
	exception.PanicIfNeeded(err)
	defer statement.Close()

	result, err := statement.Exec(schedule.Code, schedule.Name, schedule.StartTime, schedule.EndTime, schedule.LecturerName, schedule.Day)
	exception.PanicIfNeeded(err)

	id, err := result.LastInsertId()
	exception.PanicIfNeeded(err)

	fmt.Println(id)
	schedule.Id = int(id)
}

func (repository *scheduleRepositoryImpl) FindAll() (schedules []entity.Schedule) {
	query := "SELECT id, code, name, start_time, end_time, lecturer_name, day FROM schedules ORDER BY day, start_time"

	rows, err := repository.connection.Query(query)
	exception.PanicIfNeeded(err)
	defer rows.Close()
	repository.serializeRows(rows, &schedules)

	return schedules
}

func (repository *scheduleRepositoryImpl) Today() (schedules []entity.Schedule) {
	today := time.Now().Weekday()
	query := "SELECT id, code, name, start_time, end_time, lecturer_name, day FROM schedules WHERE day = ? ORDER BY start_time"

	statement, err := repository.connection.Prepare(query)
	exception.PanicIfNeeded(err)

	rows, err := statement.Query(today)
	exception.PanicIfNeeded(err)

	defer statement.Close()
	repository.serializeRows(rows, &schedules)

	return schedules
}

func (repository *scheduleRepositoryImpl) FindById(id int) (schedule entity.Schedule, err error) {
	query := "SELECT id, code, name, start_time, end_time, lecturer_name, day FROM schedules WHERE id = ? LIMIT 1"

	statement, err := repository.connection.Prepare(query)
	exception.PanicIfNeeded(err)
	defer statement.Close()

	row := statement.QueryRow(id)
	err = repository.serializeRow(row, &schedule)

	return schedule, err
}

func (repository *scheduleRepositoryImpl) Edit(schedule entity.Schedule) error {
	query := "UPDATE schedules SET code = ?, name = ?, start_time = ?, end_time = ?, lecturer_name = ?, day = ? WHERE id = ?"

	statement, err := repository.connection.Prepare(query)
	exception.PanicIfNeeded(err)
	defer statement.Close()

	result, err := statement.Exec(schedule.Code, schedule.Name, schedule.StartTime, schedule.EndTime, schedule.LecturerName, schedule.Day, schedule.Id)
	exception.PanicIfNeeded(err)

	affectedRows, err := result.RowsAffected()
	exception.PanicIfNeeded(err)

	if affectedRows == 0 {
		return errors.New("No row affected.")
	}

	return nil
}

func (repository *scheduleRepositoryImpl) Delete(id int) {
	query := "DELETE FROM schedules WHERE id = ?"

	statement, err := repository.connection.Prepare(query)
	exception.PanicIfNeeded(err)

	_, err = statement.Exec(id)
	exception.PanicIfNeeded(err)

	statement.Close()
}

func (repository *scheduleRepositoryImpl) DeleteAll() {
	//
}

func (repository *scheduleRepositoryImpl) serializeRows(rows *sql.Rows, schedules *[]entity.Schedule) {
	for rows.Next() {
		var schedule entity.Schedule

		err := rows.Scan(&schedule.Id, &schedule.Code, &schedule.Name, &schedule.StartTime, &schedule.EndTime, &schedule.LecturerName, &schedule.Day)
		exception.PanicIfNeeded(err)

		*schedules = append(*schedules, schedule)
	}
}

func (repository *scheduleRepositoryImpl) serializeRow(row *sql.Row, schedule *entity.Schedule) error {
	err := row.Scan(&schedule.Id, &schedule.Code, &schedule.Name, &schedule.StartTime, &schedule.EndTime, &schedule.LecturerName, &schedule.Day)

	if err == sql.ErrNoRows {
		return nil
	}

	return err
}
