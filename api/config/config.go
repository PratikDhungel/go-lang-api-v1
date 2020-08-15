package config

import (
	"database/sql"
	"fmt"
)

// DB - DB connection variable
var DB *sql.DB

// ConnectToDB - connection to database
func ConnectToDB() error{

	dbDriver := "mysql"
	dbUser := "root"
	//dbPass := ""
	dbName := "QualityDashboard"

	var err error

	//DB, err = sql.Open(dbDriver, "root@tcp(127.0.0.1:3306)/QualityDashboard")

	DB, err = sql.Open(dbDriver, dbUser + ":" + "@/" + dbName)

	fmt.Println("Successfully connected to the database")

	return err
}
