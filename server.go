package main

import (
	"context"
	"embed"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "embed"

	"github.com/gorilla/mux"
)

//go:embed static/*
var static embed.FS

// CustomFS is a custom filesystem to prevent directory listings
type CustomFS struct {
	fs http.FileSystem
}

// Open opens a file
func (cfs CustomFS) Open(p string) (http.File, error) {
	f, err := cfs.fs.Open(p)
	if err != nil {
		return nil, err
	}

	// Checking to see if the given path is a directory, if true an error is returned
	s, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if s.IsDir() {
		return nil, os.ErrNotExist
	}

	return f, nil
}

// address holds the address in which we'll run the server e.g. 0.0.0.0:8080
var address string

func init() {
	flag.StringVar(&address, "address", "0.0.0.0:8080",
		"address used to run the server including port number, typically in the form ':<port>'")
}

func main() {
	// Parse any flags defined in the init() function
	flag.Parse()

	// New ServeMux (router)
	r := mux.NewRouter()

	// File server with a custom filesystem to serve static assets
	fs := http.FileServer(CustomFS{http.FS(static)})
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	// Root handler
	r.HandleFunc("/", indexHandler).Methods("GET")

	// Create a new custom server
	server := &http.Server{
		Addr:    address,
		Handler: r,
	}

	// Start the server in it's own goroutine, so we can implement graceful shutdown
	go func() {
		// Ignoring the error here as it's a very basic web server
		// In a real production environment, the error should be handled properly
		_ = server.ListenAndServe()
	}()

	// Channel to receive and handle interrupt and kill signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	// Block until something is received on the channel i.e. an interrupt or kill signal
	<-c

	// Creating a context which gives the server 10 seconds to close all active and idle connections before server shutdown
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
