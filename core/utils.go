package core

import (
	"log"
)

func ErrorHandler(err error) {
	if err != nil {
		log.Panic(err)
	}
}
