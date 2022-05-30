package env

import (
	"os"
)

var DatabaseTimeout = os.Getenv("DatabaseTimeout")
var DatabaseName = os.Getenv("DatabaseName")
var DatabaseAddress = os.Getenv("DatabaseAddress")
var DatabasePort = os.Getenv("DatabasePort")
var DatabaseUsername = os.Getenv("DatabaseUsername")
var DatabasePassword = os.Getenv("DatabasePassword")

func checkDatabaseTimeout() {
	if len(DatabaseTimeout) == 0 {
		DatabaseTimeout = "2"
	}
}

func checkDatabaseName() {
	if len(DatabaseName) == 0 {
		DatabaseName = "aboutme"
	}
}

func checkDatabaseAddress() {
	if len(DatabaseAddress) == 0 {
		DatabaseAddress = "localhost"
	}
}

func checkDatabasePort() {
	if len(DatabasePort) == 0 {
		DatabasePort = "27017"
	}
}

func checkDatabaseUsername() {
	if len(DatabaseUsername) == 0 {
		DatabaseUsername = "dev"
	}
}

func checkDatabasePassword() {
	if len(DatabasePassword) == 0 {
		DatabasePassword = "dev"
	}
}

func CheckEnvs() {
	checkDatabaseTimeout()
	checkDatabaseName()
	checkDatabaseAddress()
	checkDatabasePort()
	checkDatabaseUsername()
	checkDatabasePassword()
}
