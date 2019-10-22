package utils

import (
	"log"
)

func FailIfError(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func FailIfNotInitialized(i int, msg string) {
	if i == 0 {
		log.Fatalf("%s", msg)
	}
}
