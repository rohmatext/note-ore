package config

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger(viper *viper.Viper) *logrus.Logger {
	log := logrus.New()
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err = os.Mkdir("logs", 0755)

		if err != nil {
			// Failed to create logs directory. Log only to console
			log.Warnf("Cannot create logs directory: %v", err)
			return nil
		}
	}

	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		// Failed to open log file. Log only to console
		log.Warnf("Cannot open log file: %v.", file)
		return nil
	}

	log.SetLevel(logrus.Level(viper.GetInt32("LOG_LEVEL")))
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(io.MultiWriter(os.Stdout, file))

	return log
}
