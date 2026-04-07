package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ashwathsingh/learning-api-golang/internal/product"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// global struct for api which has methods
type application struct {
	config config
	// logger
	// db driver
}

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID) // important for rate limiting
	r.Use(middleware.RealIP) 	// important for analytics, rate-limiting, and tracing
	r.Use(middleware.Logger)	// optional
	r.Use(middleware.Recoverer) // for crashes

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// how to define a method
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})

	productHandler := product.NewHandler(nil)

	r.Get("/products", productHandler.ListProducts)
	//	http.ListenAndServe(":3333", r)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second * 30,
		IdleTimeout: time.Minute,
	}

	log.Printf("server has started at addr %s", app.config.addr)

	return srv.ListenAndServe()
}

type config struct {
	addr string // address of the server (the port)
	db   dbConfig
}

type dbConfig struct {
	dsn string // domain string to connect to db
}
