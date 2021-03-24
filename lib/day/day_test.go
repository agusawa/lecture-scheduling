package day

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var daysTest = []struct {
	id   time.Weekday
	name string
}{
	{time.Sunday, time.Sunday.String()},
	{time.Monday, time.Monday.String()},
	{time.Tuesday, time.Tuesday.String()},
	{time.Wednesday, time.Wednesday.String()},
	{time.Thursday, time.Thursday.String()},
	{time.Friday, time.Friday.String()},
	{time.Saturday, time.Saturday.String()},
}

func TestDay_newDay(t *testing.T) {
	for _, dayTest := range daysTest {
		t.Run(fmt.Sprintf("newDay(%d,\"%s\")", dayTest.id, dayTest.name), func(t *testing.T) {
			day := newDay(dayTest.id, dayTest.name)

			assert.Equal(t, int8(dayTest.id), day.Id)
			assert.Equal(t, dayTest.name, day.Name)
		})
	}
}

func TestDay_GetDays(t *testing.T) {
	days := GetDays()

	for index, dayTest := range daysTest {
		t.Run(fmt.Sprintf("check %d %s", dayTest.id, dayTest.name), func(t *testing.T) {
			assert.Equal(t, int8(dayTest.id), days[index].Id)
			assert.Equal(t, dayTest.name, days[index].Name)
		})
	}
}

func TestDay_DayOf(t *testing.T) {
	for _, dayTest := range daysTest {
		t.Run(fmt.Sprintf("DayOf(%d)", dayTest.id), func(t *testing.T) {
			day, err := DayOf(int8(dayTest.id))

			assert.NoError(t, err)
			assert.Equal(t, int8(dayTest.id), day.Id)
			assert.Equal(t, dayTest.name, day.Name)
		})
	}

	t.Run(fmt.Sprintf("DayOf(7)"), func(t *testing.T) {
		day, err := DayOf(7)

		assert.Error(t, err)
		assert.Equal(t, 0, int(day.Id))
		assert.Equal(t, "", day.Name)
	})
}
