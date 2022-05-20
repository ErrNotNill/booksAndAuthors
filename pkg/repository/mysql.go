package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const urlMysql = "user:pass@/dbname"

type Sql struct {
	conn *sql.DB
}

func (s *Sql)ConnectMysql() error{
db, err := sql.Open("mysql", urlMysql)

if err != nil {
panic(err)
}
defer func(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf(err.Error())
	}
}(db)

result, err := db.Exec("insert into Books (id, title) values (?, ?)",
1, "test")
if err != nil{
log.Fatalf(err.Error())
}
fmt.Println(sql.ErrTxDone)
fmt.Println(result.LastInsertId())
fmt.Println(result.RowsAffected())
return err
}