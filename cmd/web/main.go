package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/matthewrankin/hypermedia-contact/internal/models"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
}

type application struct {
	logger   *slog.Logger
	contacts *models.ContactModel
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4100, "server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("POSTGRES_URL"), "PostgreSQL DSN")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(
		&cfg.db.maxIdleTime,
		"db-max-idle-time",
		15*time.Minute,
		"PostgreSQL max idle time",
	)
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(cfg)
	if err != nil {
		logger.Error(
			"error starting database",
			slog.Any("dsn", cfg.db.dsn),
			slog.Any("err", err),
		)
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("database connection pool established", slog.Any("dsn", cfg.db.dsn))

	app := &application{
		logger:   logger,
		contacts: &models.ContactModel{DB: db},
	}

	logger.Info("starting server", slog.Any("env", cfg.env), slog.Any("port", cfg.port))

	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetConnMaxIdleTime(cfg.db.maxIdleTime)

	// Ping the database to make sure it's up and running. Since the Docker
	// database container could take longer to spin up, we'll retry a few times
	// before truly failing.
	retries := 30
	for {
		if err = db.Ping(); err != nil && retries <= 0 {
			return nil, err
		} else if err == nil {
			break
		}
		retries--
		time.Sleep(900 * time.Millisecond)
	}
	return db, nil
}
