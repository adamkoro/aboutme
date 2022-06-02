package env

import (
	"fmt"
	"testing"
)

func TestCheckDatabaseTimeout(t *testing.T) {
	testData := 2
	defaultData := CheckDatabaseTimeout(testData)
	if defaultData != testData {
		message := fmt.Sprintf("Testdata: %d and function return value: %d not match", testData, defaultData)
		t.Error(message)
	}
	customData := 10
	defaultData = CheckDatabaseTimeout(customData)
	if defaultData != customData {
		message := fmt.Sprintf("Customdata: %d and function return value: %d not match", customData, defaultData)
		t.Error(message)
	}
}

func TestCheckDatabaseName(t *testing.T) {
	testData := "aboutme"
	defaultData := CheckDatabaseName(testData)
	if defaultData != testData {
		fmt.Println(defaultData)
		message := fmt.Sprintf("Testdata: %s and function return value: %s not match", testData, defaultData)
		t.Error(message)
	}
}

func TestCheckAddress(t *testing.T) {
	testData := "localhost"
	defaultData := CheckDatabaseAddress(testData)
	if defaultData != testData {
		fmt.Println(defaultData)
		message := fmt.Sprintf("Testdata: %s and function return value: %s not match", testData, defaultData)
		t.Error(message)
	}
}

func TestCheckDatabasePort(t *testing.T) {
	testData := 27017
	defaultData := CheckDatabasePort(testData)
	if defaultData != testData {
		message := fmt.Sprintf("Testdata: %d and function return value: %d not match", testData, defaultData)
		t.Error(message)
	}
	customData := 27777
	defaultData = CheckDatabaseTimeout(customData)
	if defaultData != customData {
		message := fmt.Sprintf("Customdata: %d and function return value: %d not match", customData, defaultData)
		t.Error(message)
	}
}

func TestCheckDatabaseUsername(t *testing.T) {
	testData := "dev"
	defaultData := CheckDatabaseUsername(testData)
	if defaultData != testData {
		fmt.Println(defaultData)
		message := fmt.Sprintf("Testdata: %s and function return value: %s not match", testData, defaultData)
		t.Error(message)
	}
}

func TestCheckDatabasePassword(t *testing.T) {
	testData := "dev"
	defaultData := CheckDatabasePassword(testData)
	if defaultData != testData {
		fmt.Println(defaultData)
		message := fmt.Sprintf("Testdata: %s and function return value: %s not match", testData, defaultData)
		t.Error(message)
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
