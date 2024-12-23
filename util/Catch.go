package util

import "log"

func Catch() {
	errRecovery := recover()
	if errRecovery != nil {
		log.Println("Panic occured : ", errRecovery)
	}
}
