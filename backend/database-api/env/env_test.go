package env

import (
	"fmt"
	"testing"
)

func TestCheckDatabaseTimeout(t *testing.T) {
	timeOut := CheckDatabaseTimeout(0)
	if timeOut != 2 {
		fmt.Println(timeOut)
		t.Error("0 value is not equals default value, which is 2")
	}
	customTimeOut := 10
	timeOut = CheckDatabaseTimeout(customTimeOut)
	if customTimeOut != timeOut {
		fmt.Println(timeOut)
		message := fmt.Sprintf("Test timeout value (%d) is not equals itself", customTimeOut)
		t.Error(message)
	}
}

func TestCheckDatabaseName(t *testing.T) {
	db := CheckDatabaseName("aboutme")
	if db != "aboutme" {
		fmt.Println(db)
		t.Error("Default database name value (aboutme) is not setted")
	}
}

func TestConvertTimeoutStringToInt(t *testing.T) {
	stringTimeOut := "2"
	timeOut, _ := ConvertTimeoutStringToInt(stringTimeOut)
	if timeOut != 2 {
		message := fmt.Sprintf("Convert timeout string: %s to int: %d not the same value or type", stringTimeOut, timeOut)
		t.Error(message)
	}
}

func TestConvertPortStringToInt(t *testing.T) {
	stringPort := "27017"
	port, _ := ConvertPortStringToInt(stringPort)
	if port != 27017 {
		message := fmt.Sprintf("Convert port string: %s to int: %d not the same value or type", stringPort, port)
		t.Error(message)
	}
}

func TestCheckEnvs(t *testing.T) {
	CheckEnvs()
	message := fmt.Sprintf(" Default timeout: %d\n Default database name: %s\n Default database address: %s\n Default database port: %d\n Default database user: %s \n Default database user password: %s", DatabaseTimeout, DatabaseName, DatabaseAddress, DatabasePort, DatabaseUsername, DatabasePassword)
	fmt.Println(message)
}
