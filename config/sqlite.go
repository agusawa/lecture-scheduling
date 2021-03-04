package config

import (
	"database/sql"
	"lecture-scheduling/exception"
	"os"
)

func NewSqliteDatabase(config Config) *sql.DB {
	filename := config.Get("SQL_FILENAME")

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		createNewDatabaseFile(filename)
	}

	connection, err := sql.Open("sqlite3", filename)
	exception.PanicIfNeeded(err)

	createTable(connection)

	return connection
}

func createNewDatabaseFile(filename string) {
	file, err := os.Create(filename)
	exception.PanicIfNeeded(err)
	file.Close()
}

func createTable(database *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS schedules (
		id 				INTEGER 		NOT NULL PRIMARY KEY AUTOINCREMENT,
		code 			VARCHAR(10) 	NULL,
		name 			VARCHAR(100) 	NOT NULL,
		start_time 		TIME 			NOT NULL,
		end_time 		TIME 			NOT NULL,
		lecturer_name	VARCHAR(100) 	NOT NULL,
		day				TINYINT(1)		NOT NULL
	)`

	_, err := database.Exec(query)
	exception.PanicIfNeeded(err)
}
