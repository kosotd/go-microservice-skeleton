package utils

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func FailIfNotInitialized(i int, msg string) {
	if i == 0 {
		log.Fatalf(fmt.Sprintf("%s", msg))
	}
}

func RecoverAndSetStatus(w http.ResponseWriter, code int) {
	if r := recover(); r != nil {
		LogAndSetStatusIfError(w, code, errors.New(fmt.Sprint(r)))
	}
}
