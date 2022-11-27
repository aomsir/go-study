package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	connection, _ := sql.Open("mysql", "root:123456@tcp(localhost:3307)/gostudy")
	defer func() {
		if connection != nil {
			connection.Close()
		}
	}()

	stmp, _ := connection.Prepare("SELECT * FROM people")
	defer func() {
		if stmp != nil {
			stmp.Close()
		}
	}()

	rows, _ := stmp.Query()
	for rows.Next() {
		var id int
		var name, address string
		rows.Scan(&id, &name, &address) // 将行内值赋值给变量
		fmt.Println(id, name, address)
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
}
