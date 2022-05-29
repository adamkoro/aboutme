package errorhandler

// Error handling functions

import "log"

func IsFatalError(err error) bool {
	if err != nil {
		log.Fatal(err)
		return true
	}
	return false
}

func IsError(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}
