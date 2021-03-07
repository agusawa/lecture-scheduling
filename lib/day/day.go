package day

import (
	"errors"
	"time"
)

type Day struct {
	Id   int8
	Name string
}

type Days []Day

func GetDays() Days {
	return Days{
		newDay(time.Sunday, time.Sunday.String()),
		newDay(time.Monday, time.Monday.String()),
		newDay(time.Tuesday, time.Tuesday.String()),
		newDay(time.Wednesday, time.Wednesday.String()),
		newDay(time.Thursday, time.Thursday.String()),
		newDay(time.Friday, time.Friday.String()),
		newDay(time.Saturday, time.Saturday.String()),
	}
}

func DayOf(id int8) (day Day, err error) {
	for _, dayLoop := range GetDays() {
		if dayLoop.Id == id {
			day = dayLoop
		}
	}

	if day.Id == 0 {
		return day, errors.New("Invalid day id.")
	}

	return day, nil
}

func newDay(id time.Weekday, name string) Day {
	return Day{
		Id:   int8(id),
		Name: name,
	}
}
