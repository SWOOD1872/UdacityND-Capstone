package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	address := flag.String("address", "0.0.0.0:8080", "address used to run the server, typically in the form ':<port>'")
	flag.Parse()

	// New ServeMux (router)
	r := mux.NewRouter()

	// File server to serve static files
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	// Root handler
	r.HandleFunc("/", indexHandler).Methods("GET")

	// Create a new custom server
	server := &http.Server{
		Addr:    *address,
		Handler: r,
	}

	// Start the server in it's own goroutine, so we can implement graceful shutdown
	go func() {
		_ = server.ListenAndServe() // Ignoring the error here as it's a very basic web server
	}()

	// Channel to receive and handle interrupt and kill signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	// Block until something is received on the channel i.e. an interrupt or kill signal
	<-c

	// Create a contect which gives the server 10 seconds to close all active and idle connections
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Unable to gracefully shutdown the server")
	}
}

// default/root handler which serves the index page and associated styling
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	http.ServeFile(w, r, "./static/index.html")
}
