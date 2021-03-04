package service

type ScheduleService interface {
	Add() error

	ShowAllSchedules()

	ShowTodaySchedule()
}
