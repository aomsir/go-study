package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connection, _ := sql.Open("mysql", "root:123456@tcp(localhost:3307)/gostudy")
	connection.Ping() // 在这里就连接
	defer func() {
		if connection != nil {
			connection.Close()
			fmt.Println("connection关闭成功")
		}
	}()

	stmp, _ := connection.Prepare("DELETE FROM people WHERE id = ?")
	defer func() {
		if stmp != nil {
			stmp.Close()
			fmt.Println("stmp关闭成功")
		}
	}()

	response, _ := stmp.Exec(3)
	count, _ := response.RowsAffected()
	if count > 0 {
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除失败")
	}

}
