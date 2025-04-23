package main

import (
	"flag"
	"fmt"
	"howardsolutions/go/backend-project/internal/app"
	"howardsolutions/go/backend-project/internal/routes"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	var port int
	flag.IntVar(&port, "port", 8080, "Go backend server port")
	flag.Parse()

	application, err := app.NewApplication()
	if err != nil {
		logger.Fatalf("Failed to initialize application: %v", err)
	}

	// When everything is done => close the db connection => Defer it till the end
	defer application.DB.Close()

	application.Logger.Println("Application initialized successfully!")

	r := routes.SetupRoutes(application)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	application.Logger.Printf("Server starting on port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		application.Logger.Fatalf("Server failed to start: %v", err)
	}
}
