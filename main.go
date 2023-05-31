package main

import (
	"github.com/YuYAlexey/TestTask/internal/app"
	"github.com/YuYAlexey/TestTask/internal/db"
	"github.com/YuYAlexey/TestTask/internal/transport/http"
)

func main() {
	db, err := db.New()
	if err != nil {
		panic(err)
	}

	app := app.New(db)

	if err := http.Service(app); err != nil {
		panic(err)
	}
}
