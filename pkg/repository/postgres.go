package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
)

const urlPg = "postgres://postgres:root@localhost:5432/books_authors"

var ctx = context.Background()

type Pgx struct {
	conn *pgx.Conn
}

func (p *Pgx) PostgresInit() (conn *pgx.Conn,err error) {
	p.conn = conn
	err = os.Setenv("DATABASE_URL", urlPg)
	if err != nil {
		log.Printf("Could not connect:\n%v", err.Error())
	}
	conn, err = pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		if err != nil {
			log.Fatalf(err.Error())
		}
		os.Exit(1)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
panic(err.Error())
		}
	}(conn, ctx)
	return conn,err
}