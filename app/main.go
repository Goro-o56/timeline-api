package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
	"timeline-api/app/config"
	"timeline-api/app/registry"
)

func main() {

	//Initialize logger
	logrus.SetFormatter(&logrus.JSONFormatter{})

	//Set Configure
	c := config.App{
		DB: config.DB{
			Host:     config.GetenvOrDefault("DATABASE_HOST", "127.0.0.1"),
			Port:     config.GetenvOrDefault("DATABASE_PORT", "3306"),
			Name:     config.GetenvOrDefault("DATABASE_NAME", "rest_api"),
			User:     config.GetenvOrDefault("DATABASE_USER", "rest_api"),
			Password: config.GetenvOrDefault("DATABASE_PASSWORD", ""),
			Timezone: time.Local,
		},
	}
	//Initialize application
	r := registry.NewRouter(c)

	host := config.GetenvOrDefault("APP_HOST", "0.0.0.0")
	port := config.GetenvOrDefault("APP_PORT", "8080")
	if err := r.Run(fmt.Sprintf("%s:%s", host, port)); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

}
