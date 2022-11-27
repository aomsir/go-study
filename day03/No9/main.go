package main

import (
	"fmt"
	"os"
)

func main() {

	// 创建文件夹
	error := os.Mkdir("/Users/aomsir/MyStudyProject/Go/study/day03/No9/src", os.ModeDir)   // 要求创建的文件夹不存在,但是父目录必须存在
	error = os.MkdirAll("/Users/aomsir/MyStudyProject/Go/study/day03/No9/src", os.ModeDir) // 没有什么限制,父目录没有就帮忙创建
	if error != nil {
		fmt.Println("文件夹创建失败", error)
		return
	} else {
		fmt.Println("文件夹创建成功")
	}

	// 创建文件
	file, error := os.Create("/Users/aomsir/MyStudyProject/Go/study/day03/No9/src/demo.txt")
	if error != nil {
		fmt.Println("文件创建失败", error)
		return
	} else {
		fmt.Println(file.Name(), "文件创建成功")
	}

	// 重命名文件夹
	// os.Rename("oldpath", "newpath")

	// 获取文件信息
	file, error = os.Open("/Users/aomsir/MyStudyProject/Go/study/day03/No9/src/demo.txt")
	if error != nil {
		fmt.Println("文件获取失败", error)
		return
	}
	fileInfo, error := file.Stat() // 获取文件状态
	if error != nil {
		fmt.Println("文件信息获取失败", error)
		return
	} else {
		fmt.Println(fileInfo.Size())
	}

	// 删除文件夹 - Remove/RemoveAll - 规则和上面Mkdir的一致
	error = os.Remove("/Users/aomsir/MyStudyProject/Go/study/day03/No9/src/demo.txt")
	if error != nil {
		fmt.Println("删除失败", error)
		return
	} else {
		fmt.Println("删除成功")
	}
}
