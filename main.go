package main

import (
	"fmt"

	"gomysql.com/db"
)

func main() {
	var con *sql.DB
	fmt.Println("phpflow.com – Go MySQL Tutorial")
	con := db.CreateCon()

	// Execute the query
	rows, err := con.Query("SELECT * FROM t1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Println("rows: ", rows)

}
