package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("/Users/aomsir/MyStudyProject/Go/study/day03/No11/main.go") // 读取文件
	fileInfo, _ := file.Stat()                                                     // 获取文件的信息

	buffer := make([]byte, fileInfo.Size()) // 创建buffer缓冲区

	file.Read(buffer) // 将文件里面的内容读取到buffer里面

	fmt.Println("文件中的内容为:", string(buffer))
}
