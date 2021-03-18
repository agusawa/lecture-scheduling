package service

type ScheduleService interface {
	Add() error

	Edit(id int) error

	Delete(id int) error

	ShowAllSchedulesWithId()

	ShowAllSchedules()

	ShowTodaySchedule()
}
