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

	smtp, _ := connection.Prepare("UPDATE people SET name = ?, address = ? WHERE id = ?")
	defer func() {
		if smtp != nil {
			smtp.Close()
		}
	}()

	response, _ := smtp.Exec("Jewix", "Hushi", 1)
	count, _ := response.RowsAffected()
	if count > 0 {
		fmt.Println("修改成功")
	} else {
		fmt.Println("修改失败")
	}
}
