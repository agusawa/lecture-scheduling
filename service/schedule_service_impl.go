package service

import (
	"errors"
	"fmt"
	"lecture-scheduling/entity"
	"lecture-scheduling/lib"
	"lecture-scheduling/repository"
	"time"

	"github.com/fatih/color"
)

func NewScheduleService(scheduleRepository *repository.ScheduleRepository) ScheduleService {
	return &scheduleServiceImpl{
		ScheduleRepository: *scheduleRepository,
	}
}

type scheduleServiceImpl struct {
	ScheduleRepository repository.ScheduleRepository
}

func (service *scheduleServiceImpl) Add() error {
	var schedule entity.Schedule

	fmt.Print("Class code: ")
	fmt.Scan(&schedule.Code)

	fmt.Print("\nClass name: ")
	fmt.Scan(&schedule.Name)

	fmt.Print("\nStart time: ")
	fmt.Scan(&schedule.StartTime)

	fmt.Print("\nEnd time: ")
	fmt.Scan(&schedule.EndTime)

	fmt.Print("\nLecturer name: ")
	fmt.Scan(&schedule.LecturerName)

	for _, day := range lib.Days {
		fmt.Printf("%d. %s\n", day.GetId(), day.GetName())
	}

	fmt.Print("Select day by number: ")
	fmt.Scan(&schedule.Day)

	if schedule.Day < 0 || schedule.Day > 6 {
		return errors.New("Invalid day")
	}

	service.ScheduleRepository.Add(&schedule)

	return nil
}

func (service *scheduleServiceImpl) ShowAllSchedules() {
	schedules := service.ScheduleRepository.FindAll()

	if len(schedules) == 0 {
		fmt.Println("No schedule found.")
		return
	}

	lastDay := ""

	for index, schedule := range schedules {
		dayName := schedule.GetDayName()

		if lastDay != dayName {
			lastDay = dayName
		}

		if index == 0 {
			fmt.Printf("%s\n", lastDay)
		} else {
			fmt.Printf("\n%s\n", lastDay)
		}

		service.print(schedule)
	}
}

func (service *scheduleServiceImpl) ShowTodaySchedule() {
	schedules := service.ScheduleRepository.Today()

	if len(schedules) == 0 {
		fmt.Println("No schedule today.")
	}

	for index, schedule := range schedules {
		if index == 0 {
			fmt.Printf("%s\n", schedule.GetDayName())
		}

		service.print(schedule)
	}
}

func (service *scheduleServiceImpl) print(schedule entity.Schedule) {
	format := "[%s - %s] %s %s (%s)\n"
	args := []interface{}{schedule.StartTime, schedule.EndTime, schedule.Code, schedule.Name, schedule.LecturerName}

	if time.Now().Weekday() == time.Weekday(schedule.Day) {
		color.Yellow(format, args...)
	} else {
		fmt.Printf(format, args...)
	}
}
