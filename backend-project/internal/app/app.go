package app

import (
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

// Here we want to modifying the application and using it, that's why we need to return pointer to Application (struct)
func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		Logger: logger,
	}

	return app, nil
}
