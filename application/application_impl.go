package application

import (
	"fmt"
	"lecture-scheduling/service"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

func New(scheduleService *service.ScheduleService) Application {
	return &applicationImpl{
		ScheduleService: *scheduleService,
	}
}

type applicationImpl struct {
	ScheduleService service.ScheduleService
}

func (app *applicationImpl) Run() {
	app.clearScreen()

	manageFlag, _ := strconv.ParseBool(os.Getenv("MANAGE"))
	todayFlag, _ := strconv.ParseBool(os.Getenv("TODAY"))

	if manageFlag {
		app.manage()
	} else if todayFlag {
		app.ScheduleService.ShowTodaySchedule()
		fmt.Println("\nUse --manage to manage the schedule.")
	} else {
		app.ScheduleService.ShowAllSchedules()
		fmt.Println("\nUse --manage to manage the schedule.")
	}
}

func (app *applicationImpl) clearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}

func (app *applicationImpl) manage() {
	app.showMenu()
}

func (app *applicationImpl) showMenu() {
	var input string

	fmt.Println("====== MENU ======")
	fmt.Println("1. Show schedules")
	fmt.Println("2. Add schedule")
	fmt.Println("3. Edit schedule")
	fmt.Println("4. Delete schedule")
	fmt.Println("\n9. Reset schedule")
	fmt.Println("0. Exit")
	fmt.Println("==================")
	fmt.Print("Select your choice: ")
	fmt.Scan(&input)

	switch input {
	case "1":
		app.clearScreen()

		app.ScheduleService.ShowAllSchedulesWithId()
		fmt.Print("\nPress enter to back to menu [ENTER]")
		fmt.Scanln(&input)

		app.clearScreen()
		app.showMenu()

	case "2":
		if err := app.ScheduleService.Add(); err != nil {
			app.clearScreen()
			color.Red(err.Error())
		} else {
			app.clearScreen()
			color.Green("Schedule saved successfully.")
		}
	// case "3":
	// 	//

	case "4":
		app.clearScreen()
		fmt.Print("Enter the id of the schedule you want to delete: ")
		fmt.Scanln(&input)

		app.clearScreen()
		if value, err := strconv.Atoi(input); err != nil {
			color.Red("Invalid schedule id.")
		} else {
			if err := app.ScheduleService.Delete(value); err != nil {
				color.Red(err.Error())
			} else {
				color.Green("Schedule successfully deleted.")
			}
		}
		fmt.Println("")
		app.showMenu()

	// case "9":
	// 	//
	case "0":
		fmt.Println("\nGood bye. Have a great day!")

	default:
		app.clearScreen()
		color.Red("Invalid choice!\n\n")
		app.showMenu()
	}
}
