package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"lendahand.samuelskinner.xyz/internal/repository"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	repo *repository.Queries
	cfg  config
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// temporary because I am having trouble with the flag.
	if cfg.db.dsn == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cfg.db.dsn = os.Getenv("LENDAHAND_DB_DSN")
	}

	conn, err := openDB(cfg.db.dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := &application{
		repo: repository.New(conn),
		cfg:  cfg,
	}

	err = app.server()

	if err != nil {
		os.Exit(1)
	}

}

func openDB(dsn string) (*pgx.Conn, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
	defer conn.Close(ctx)

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
