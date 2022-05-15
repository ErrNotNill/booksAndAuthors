package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

const urlExample = "postgres://postgres:root@localhost:5432/books_authors"

//pass root, port 5432

var ctx = context.Background()

type Pgx struct {
	conn *pgx.Conn
}

func (p *Pgx) PostgresInit() (conn *pgx.Conn) {
	p.conn = conn
	os.Setenv("DATABASE_URL", urlExample)
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {

		}
	}(conn, context.Background())
	return conn
}

type Find interface {
	FIndBook()
	FindAllBooks()
	FindAuthor()
	FindAllAuthors()
}

func FindBook() {
	/* todo
	searching in title
	find book = title, date, author
	*/
}
func FindAllBooks() {

}
func FindAuthor() {

}
func FindAllAuthors() {

}

func (p *Pgx) Query() {

	var name string
	var weight int64
	p.conn.Query(ctx, "select * from books")
	err := p.conn.QueryRow(context.Background(), "select * from books").Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	//p.conn.Exec()
	fmt.Println(name, weight)
}
