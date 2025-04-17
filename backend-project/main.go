package main

import (
	"flag"
	"fmt"
	"howardsolutions/go/backend-project/internal/app"
	"howardsolutions/go/backend-project/internal/routes"
	"net/http"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Go backend server port")
	flag.Parse()

	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}
	// When everything is done => close the db connection => Defer it till the end
	defer app.DB.Close()

	app.Logger.Println("We are running our app!")

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("We are running on port %d \n", port)

	err = server.ListenAndServe()

	if err != nil {
		app.Logger.Fatal(err)
	}
}
