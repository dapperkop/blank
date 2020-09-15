package database

import (
	"database/sql"
	"strconv"

	"github.com/dapperkop/blank/logger"
	"github.com/dapperkop/blank/types"
	_ "github.com/go-sql-driver/mysql" // import mysql driver for sql package
)

var (
	// DB var ...
	DB *sql.DB

	driver   string
	hostname string
	username string
	password string
	dbname   string
	dbport   int
)

func connect() {
	var (
		dataSource string
		err        error
	)

	dataSource += username + ":" + password
	dataSource += "@tcp(" + hostname + ":" + strconv.FormatInt(int64(dbport), 10) + ")"
	dataSource += "/" + dbname + "?parseTime=true"

	DB, err = sql.Open(driver, dataSource)

	if err != nil {
		logger.Logger.Fatalln(err)
	}
}

// Setup func ...
func Setup(config types.Database) {
	driver = config.Driver
	hostname = config.Hostname
	username = config.Username
	password = config.Password
	dbname = config.DBName
	dbport = config.DBPort

	connect()
}
