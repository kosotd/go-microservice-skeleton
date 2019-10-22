package utils

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func LogInfo(msg string) {
	log.Printf("INFO: %s", msg)
}

func LogError(msg string) {
	log.Printf("ERROR: %s", msg)
}

func LogAndSetStatus(w http.ResponseWriter, code int, err error) {
	LogError(err.Error())
	http.Error(w, http.StatusText(code), code)
}

func LogAndSetStatusIfError(w http.ResponseWriter, code int, err error) {
	if err != nil {
		LogAndSetStatus(w, code, err)
	}
}

func LogAndSetStatusIfRecover(w http.ResponseWriter, code int) {
	if r := recover(); r != nil {
		LogAndSetStatus(w, code, errors.New(fmt.Sprint(r)))
	}
}
