package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//creating structure
type User struct {
	Name string `json:"name"`
}

func main() {
	//connect to database
	fmt.Println("Go connect to msql")
	db, err := sql.Open("mysql", "root:Sharma123@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected")
	defer db.Close()

	//insert to database
	insert, err := db.Query("INSERT INTO Persons VALUES(3,'bla-bla-bla')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("inserted")

	//read from database
	results, err := db.Query("SELECT name FROM Persons")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var user User
		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.Name)
	}

}
