package lib

import "time"

var Days = []Day{
	newDay(time.Sunday, time.Sunday.String()),
	newDay(time.Monday, time.Monday.String()),
	newDay(time.Tuesday, time.Tuesday.String()),
	newDay(time.Wednesday, time.Wednesday.String()),
	newDay(time.Thursday, time.Thursday.String()),
	newDay(time.Friday, time.Friday.String()),
	newDay(time.Saturday, time.Saturday.String()),
}

func GetDay(key int8) *day {
	var selectedDay *day

	for _, dayLoop := range Days {
		if dayLoop.(*day).Id == key {
			selectedDay = dayLoop.(*day)
		}
	}

	return selectedDay
}

func newDay(id time.Weekday, name string) Day {
	return &day{
		Id:   int8(id),
		Name: name,
	}
}

type day struct {
	Id   int8
	Name string
}

func (d *day) GetId() int8 {
	return d.Id
}

func (d *day) GetName() string {
	return d.Name
}
