package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/HrushiBorhade/golang/melkey-fm-course/project/internal/app"
)

func main() {

	var port int
	flag.IntVar(&port, "port", 8080, "Go Backend Server Port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/health", HealthCheck)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
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

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is up and running ✅\n")
}
