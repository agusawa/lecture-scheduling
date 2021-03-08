package service

type ScheduleService interface {
	Add() error

	Delete(id int) error

	ShowAllSchedulesWithId()

	ShowAllSchedules()

	ShowTodaySchedule()
}
