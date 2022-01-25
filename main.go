package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SolBaa/Greenlight/cmd/handlers" // New import
	"github.com/SolBaa/Greenlight/internal/data"

	// New import
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// type Config struct {
// 	Port int
// 	Env  string
// 	db   struct {
// 		dsn string
// 	}
// }

// type Application struct {
// 	Config Config
// 	Logger *log.Logger
// }

func main() {
	var cfg handlers.Config
	var appl handlers.Application
	flag.IntVar(&cfg.Port, "port", 8000, "API server port")

	flag.StringVar(&cfg.Env, "env", "development", "Environment(development|staging|production)")
	flag.StringVar(&cfg.DB.DSN, "db-dsn", "postgres://greenlight:pa55word@db/greenlight?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	// Defer a call to db.Close() so that the connection pool is closed before the
	// main() function exits.
	defer db.Close()
	// Also log a message to say that the connection pool has been successfully
	// established.
	logger.Printf("database connection pool established")

	migrationDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Fatal(err, nil)
	}
	migrator, err := migrate.NewWithDatabaseInstance("./migrations", "postgres", migrationDriver)
	if err != nil {
		logger.Fatal(err, nil)
	}
	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.Fatal(err, nil)
	}
	logger.Printf("database migrations applied")

	mux := http.NewServeMux()

	app := &handlers.Application{
		Config: cfg,
		Logger: logger,
		Models: data.NewModels(db),
	}
	mux.HandleFunc("/v1/healthcheck", app.HealthcheckHandler)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      appl.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server.
	logger.Printf("starting %s server on %s", cfg.Env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

// The openDB() function returns a sql.DB connection pool.
func openDB(cfg handlers.Config) (*sql.DB, error) {
	// Use sql.Open() to create an empty connection pool, using the DSN from the config
	// struct.
	db, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}
	// Create a context with a 5-second timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Use PingContext() to establish a new connection to the database, passing in the
	// context we created above as a parameter. If the connection couldn't be
	// established successfully within the 5 second deadline, then this will return an
	// error.
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	// Return the sql.DB connection pool.
	return db, nil
}
