package env

import (
	"errors"
	"os"
	"strconv"

	"github.com/adamkoro/aboutme/backend/database-api/errorHandler"
)

var DatabaseTimeout int
var DatabaseName string
var DatabaseAddress string
var DatabasePort int
var DatabaseUsername string
var DatabasePassword string

// Get environment variables
func getDatabaseTimeout() string {
	return os.Getenv("DatabaseTimeout")
}

func getDatabaseName() string {
	return os.Getenv("DatabaseName")
}

func getDatabaseAddress() string {
	return os.Getenv("DatabaseAddress")
}

func getDatabasePort() string {
	return os.Getenv("DatabasePort")
}

func getDatabaseUsername() string {
	return os.Getenv("DatabaseUsername")
}

func getDatabasePassword() string {
	return os.Getenv("DatabasePassword")
}

// Set default value to variable if empty
func CheckDatabaseTimeout(timeout int) int {
	if timeout != 0 {
		return timeout
	}
	timeout = 2
	return timeout
}

func CheckDatabaseName(dbname string) string {
	if len(dbname) != 0 {
		return dbname
	}
	dbname = "aboutme"
	return dbname
}

func CheckDatabaseAddress(address string) string {
	if len(address) != 0 {
		return address
	}
	address = "localhost"
	return address
}

func CheckDatabasePort(port int) int {
	if port != 0 {
		return port
	}
	port = 27017
	return port
}

func CheckDatabaseUsername(username string) string {
	if len(username) != 0 {
		return username
	}
	username = "dev"
	return username
}

func CheckDatabasePassword(password string) string {
	if len(password) != 0 {
		return password
	}
	password = "dev"
	return password
}

func ConvertTimeoutStringToInt(t string) (int, error) {
	if t != "" {
		timeout, err := strconv.Atoi(t)
		if err != nil {
			errorHandler.ErrorLogger.Println(err)
			return 0, err
		}
		return timeout, nil
	}
	err := errors.New("string is empty")
	return 0, err
}

func ConvertPortStringToInt(p string) (int, error) {
	if p != "" {
		port, err := strconv.Atoi(p)
		if err != nil {
			errorHandler.ErrorLogger.Println(err)
			return 0, err
		}
		return port, nil
	}
	err := errors.New("string is empty")
	return 0, err
}

// Set variables value from environment variables
func getValues() {
	timeout := getDatabaseTimeout()
	DatabaseTimeout, _ = ConvertTimeoutStringToInt(timeout)
	DatabaseName = getDatabaseName()
	DatabaseAddress = getDatabaseAddress()
	port := getDatabasePort()
	DatabasePort, _ = ConvertPortStringToInt(port)
	DatabaseUsername = getDatabaseUsername()
	DatabasePassword = getDatabasePassword()

}

// Override variables
func CheckEnvs() {
	getValues()
	DatabaseTimeout = CheckDatabaseTimeout(DatabaseTimeout)
	DatabaseName = CheckDatabaseName(DatabaseName)
	DatabaseAddress = CheckDatabaseAddress(DatabaseAddress)
	DatabasePort = CheckDatabasePort(DatabasePort)
	DatabaseUsername = CheckDatabaseUsername(DatabaseUsername)
	DatabasePassword = CheckDatabasePassword(DatabasePassword)
}
