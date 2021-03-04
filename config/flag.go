package config

import (
	"flag"
	"os"
	"strconv"
)

func InitFlags() {
	today := flag.Bool("today", false, "Today's schedule")
	manage := flag.Bool("manage", false, "Manage the schedule")

	flag.Parse()

	os.Setenv("TODAY", strconv.FormatBool(*today))
	os.Setenv("MANAGE", strconv.FormatBool(*manage))
}
