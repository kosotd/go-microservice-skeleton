package utils

import (
	"log"
	"net/http"
)

func LogInfo(msg string) {
	log.Printf("INFO: %s", msg)
}

func LogError(msg string) {
	log.Printf("ERROR: %s", msg)
}

func LogAndSetStatusIfError(w http.ResponseWriter, code int, err error) {
	if err != nil {
		LogError(err.Error())
		http.Error(w, http.StatusText(code), code)
	}
}
