package repository

import "lecture-scheduling/entity"

type ScheduleRepository interface {
	Add(schedule *entity.Schedule)

	Today() (schedule []entity.Schedule)

	FindById(id int) (schedule entity.Schedule, err error)

	FindAll() (schedules []entity.Schedule)

	Edit(schedule entity.Schedule) error

	Delete(id int)

	DeleteAll()
}
