package main

import (
	"github.com/HrushiBorhade/golang/melkey-fm-course/project/internal/app"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("App is Running ✅")
}
