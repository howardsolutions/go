package app

import (
	"database/sql"
	"fmt"
	"howardsolutions/go/backend-project/internal/api"
	"howardsolutions/go/backend-project/internal/store"
	"howardsolutions/go/backend-project/migrations"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

// Here we want to modifying the application and using it, that's why we need to return pointer to Application (struct)
func NewApplication() (*Application, error) {
	pgDB, err := store.Open()

	if err != nil {
		return nil, err
	}

	// Setup file structure for DB migration
	err = store.MigrateFS(pgDB, migrations.FS, ".")

	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Store will go here

	// Handlers
	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")
}
