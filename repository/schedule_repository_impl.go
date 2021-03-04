package repository

import (
	"database/sql"
	"fmt"
	"lecture-scheduling/entity"
	"lecture-scheduling/exception"
	"time"
)

func NewScheduleRepository(database *sql.DB) ScheduleRepository {
	return &scheduleRepositoryImpl{
		database: database,
	}
}

type scheduleRepositoryImpl struct {
	database *sql.DB
}

func (repository *scheduleRepositoryImpl) Add(schedule *entity.Schedule) {
	query := "INSERT INTO schedules (code, name, start_time, end_time, lecturer_name, day) VALUES (?, ?, ?, ?, ?, ?)"

	statement, err := repository.database.Prepare(query)
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

	rows, err := repository.database.Query(query)
	exception.PanicIfNeeded(err)
	defer rows.Close()
	repository.serializeRows(rows, &schedules)

	return schedules
}

func (repository *scheduleRepositoryImpl) Today() (schedules []entity.Schedule) {
	today := time.Now().Weekday()
	query := "SELECT id, code, name, start_time, end_time, lecturer_name, day FROM schedules WHERE day = ? ORDER BY start_time"

	statement, err := repository.database.Prepare(query)
	exception.PanicIfNeeded(err)

	rows, err := statement.Query(today)
	exception.PanicIfNeeded(err)

	defer statement.Close()
	repository.serializeRows(rows, &schedules)

	return schedules
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
