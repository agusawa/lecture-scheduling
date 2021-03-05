package service

type ScheduleService interface {
	Add() error

	ShowAllSchedulesWithId()

	ShowAllSchedules()

	ShowTodaySchedule()
}
