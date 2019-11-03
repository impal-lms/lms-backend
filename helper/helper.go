package helper

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func GetEnv(variable string, fallback interface{}) interface{} {
	err := godotenv.Load()
	if err != nil {
		logrus.Error(err)
	}

	res := os.Getenv(variable)
	if res != "" {
		return res
	}
	return fallback
}
