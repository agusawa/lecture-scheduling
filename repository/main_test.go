package repository

import (
	"database/sql"
	"lecture-scheduling/config"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db                 *sql.DB
	scheduleRepository ScheduleRepository
)

func TestMain(m *testing.M) {
	configuration := config.New("../.env.test")
	db = config.NewSqliteDatabase(configuration)
	scheduleRepository = NewScheduleRepository(db)

	m.Run()
	db.Close()
	configuration.DeleteDatabase()
}
