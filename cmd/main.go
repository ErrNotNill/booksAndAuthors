package main

import (
	"fmt"
	"github.com/ErrNotNill/booksAndAuthors/pkg/repository"
	"github.com/ErrNotNill/booksAndAuthors/service/grpc"
)

func main() {
	//check conn
	go func() {
		pg := new(repository.Pgx)
		init, err := pg.PostgresInit()
		if err != nil {
			return
		}
		init.PgConn()
	}()


mysql := new(repository.Sql)
	err := mysql.ConnectMysql()
	if err != nil {
		return
	}

//grpc descriptor
	fmt.Println(grpc.File_books_authors_proto)
}
