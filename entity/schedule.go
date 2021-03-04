package entity

import (
	"lecture-scheduling/lib"
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

func (entity *Schedule) GetDayName() string {
	return lib.GetDay(entity.Day).Name
}
