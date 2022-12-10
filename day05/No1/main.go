package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 驱动已经在标准库了,所有需要空导入
)

func main() {
	// 1、打开连接
	connection, error := sql.Open("mysql", "root:123456@tcp(localhost:3307)/gostudy")
	connection.Ping()
	defer func() {
		if connection != nil {
			fmt.Println("2")
			connection.Close()
		}
	}()
	if error != nil {
		fmt.Println("连接创建失败")
		return
	}

	// 2、预处理SQL
	// ?代表占位符
	stmp, error := connection.Prepare("INSERT INTO people values (default,?,?)")
	defer func() {
		if stmp != nil {
			fmt.Println("1")
			stmp.Close()
		}
	}()
	if error != nil {
		fmt.Println("预处理失败")
		return
	}

	response, error := stmp.Exec("Aomsir", "Wuhan")
	if error != nil {
		fmt.Println("SQL执行失败")
		return
	}

	// 3、获取结果
	count, error := response.RowsAffected()
	if error != nil {
		fmt.Println("结果获取失败")
		return
	}
	if count > 0 {
		fmt.Println("新增成功")
	} else {
		fmt.Println("新增失败")
	}

	// 获取新增主健的值
	id, _ := response.LastInsertId()
	fmt.Println(id)

}
