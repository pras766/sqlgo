package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	dat := []byte("Welcome to first page of go lang")
	w.WriteHeader(200)
	w.Write(dat)
}

func CreateTableHandler(w http.ResponseWriter, r *http.Request) {
	_, err := Db.Exec(`CREATE TABLE IF NOT EXISTS users(
		name varchar(20) not null,
		age smallint, 
		gender char(1)
	)`)
	checkError(err)
	successMsg := []byte("Table created successfully")

	w.Write(successMsg)
}

func GetTableData(w http.ResponseWriter, r *http.Request) {
	rows, err := Db.Query(`SELECT * FROM users`)

	if err != nil {
		panic(err)
	}

	var rowArr []User

	defer rows.Close()

	for rows.Next() {
		var u User

		err = rows.Scan(&u.Name, &u.Age, &u.Gender)

		rowArr = append(rowArr, u)
	}

	dat, err := json.Marshal(rowArr)

	w.Write(dat)
}

func InsertIntoTable(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}
	var u User

	err = json.Unmarshal(body, &u)

	if err != nil {
		panic(err)
	}

	fmt.Println(u.Age, u.Gender, u.Name)

	query := `INSERT INTO "users" ("name", "age", "gender") VALUES ($1, $2, $3)`

	_, err = Db.Exec(query, u.Name, u.Age, u.Gender)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Values are not inserted"))
		panic(err)
	}
	w.Write([]byte("insertion is success"))

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
