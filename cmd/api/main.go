package main

import (
	"context"
	"database/sql"
	"flag"
	_ "github.com/lib/pq"
	"github.com/rrebeiz/quickmovies/internal/data"
	"log"
	"os"
	"time"
)

type config struct {
	env  string
	port int
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	models   data.Models
}

func main() {
	var cfg config
	flag.StringVar(&cfg.env, "environment", "develop", "default app environment")
	flag.IntVar(&cfg.port, "port", 4000, "the default app port")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DSN"), "the database DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "db max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "db max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "db max idle time")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO", log.Ltime|log.Ldate|log.Llongfile)
	errorLog := log.New(os.Stdout, "ERROR", log.Ltime|log.Ltime|log.Lshortfile)

	db, err := openDB(cfg)
	if err != nil {
		log.Fatalf("failed to start the db connection %s", err)
	}
	defer db.Close()

	app := application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		models:   data.NewModels(db),
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Fatal("failed to start the server")
	}

}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
