package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HrushiBorhade/golang/melkey-fm-course/project/internal/api"
	"github.com/HrushiBorhade/golang/melkey-fm-course/project/internal/store"
	"github.com/HrushiBorhade/golang/melkey-fm-course/project/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		fmt.Printf("Error in app.go")
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	workoutHandler := api.NewWorkoutHandler()
	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}
	return app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is up and running ✅\n")
}
