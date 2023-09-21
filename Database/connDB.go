package conndb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "hello"
	dbname   = "mydb"
)

func ConnectPSQL() *sql.DB {
	// psqlinfo := fmt.Sprint("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	psqlinfo := fmt.Sprintf("host=%s port=%v dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, user, password)

	// psqlinfo := `host=127.0.0.1 port=5433 dbname=mydb user=postgres password=hello sslmode=disable`

	db, err := sql.Open("postgres", psqlinfo)

	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}
