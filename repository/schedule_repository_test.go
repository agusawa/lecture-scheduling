package repository

import (
	"lecture-scheduling/entity"
	"lecture-scheduling/exception"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScheduleRepository_Add(t *testing.T) {
	schedule := entity.Schedule{
		Id:           1,
		Code:         "ABC",
		Name:         "Matematics",
		StartTime:    "10:00",
		EndTime:      "12:00",
		LecturerName: "John Doe",
		Day:          1,
	}
	scheduleRepository.Add(&schedule)

	t.Run("ScheduleIdChanged", func(t *testing.T) {
		assert.NotNil(t, schedule.Id)
	})

	t.Run("ScheduleSaved", func(t *testing.T) {
		var actualSchedule entity.Schedule
		query := "SELECT id, code, name, start_time, end_time, lecturer_name, day FROM schedules WHERE id = ? LIMIT 1"

		row := db.QueryRow(query, schedule.Id)

		err := row.Scan(&actualSchedule.Id, &actualSchedule.Code, &actualSchedule.Name, &actualSchedule.StartTime, &actualSchedule.EndTime, &actualSchedule.LecturerName, &actualSchedule.Day)
		exception.PanicIfNeeded(err)

		assert.Equal(t, schedule.Id, actualSchedule.Id)
		assert.Equal(t, schedule.Code, actualSchedule.Code)
		assert.Equal(t, schedule.Name, actualSchedule.Name)
		assert.Equal(t, schedule.StartTime, actualSchedule.StartTime)
		assert.Equal(t, schedule.EndTime, actualSchedule.EndTime)
		assert.Equal(t, schedule.LecturerName, actualSchedule.LecturerName)
		assert.Equal(t, schedule.Day, actualSchedule.Day)
	})
}
