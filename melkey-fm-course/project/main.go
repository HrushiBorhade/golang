package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/HrushiBorhade/golang/melkey-fm-course/project/internal/app"
	"github.com/HrushiBorhade/golang/melkey-fm-course/project/internal/routes"
)

func main() {

	var port int
	flag.IntVar(&port, "port", 8080, "Go Backend Server Port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	defer app.DB.Close()

	r := routes.SetUpRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	app.Logger.Printf("App is Running on port : %d ✅\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
