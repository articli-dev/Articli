package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	EnvDebug = "DEBUG"
)

func InitLogger() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	if os.Getenv(EnvDebug) == "true" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
