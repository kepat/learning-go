package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/go")

	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	// // Prepare statement for inserting data
	// stmtIns, err := db.Prepare("INSERT INTO user (name) VALUES( ? )")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// defer stmtIns.Close()

	// stmtIns.Exec("Kris")

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT name FROM user WHERE name = ?")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmtOut.Close()

	var result string

	// Query
	err = stmtOut.QueryRow("kevin").Scan(&result)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("The name is: ", result)

}
