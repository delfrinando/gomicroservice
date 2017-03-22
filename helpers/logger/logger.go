package helpers

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func Error(err error, msg string) {
	if err != nil {
		param := fmt.Sprintf("%s: %s", msg, err)
		log.Error(param)
	}
}

func Info(msg string) {
	log.Info(msg)
}
