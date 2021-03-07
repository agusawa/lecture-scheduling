package entity

import (
	"lecture-scheduling/exception"
	"lecture-scheduling/lib/day"
)

type Schedule struct {
	Id           int
	Code         string
	Name         string
	StartTime    string
	EndTime      string
	LecturerName string
	Day          int8
}

func (entity *Schedule) GetDay() day.Day {
	dayObject, err := day.DayOf(entity.Day)
	exception.PanicIfNeeded(err)

	return dayObject
}
