package main

import (
	"fmt"
	"github.com/ErrNotNill/booksAndAuthors/pkg/repository"
)

func main() {
	//test connection
	pg := new(repository.Pgx)
	pg.PostgresInit()
	fmt.Println("postgres connected")
}
