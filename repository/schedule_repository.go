package repository

import "lecture-scheduling/entity"

type ScheduleRepository interface {
	Add(schedule *entity.Schedule)

	Today() (schedule []entity.Schedule)

	FindAll() (schedules []entity.Schedule)

	DeleteAll()
}
