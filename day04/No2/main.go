package main

import (
	"log"
	"os"
)

func main() {
	// log.Println("打印日志信息")
	// log.Panicln("打印fatal信息")
	// log.Fatalln("error")

	logFilePath := "/Users/aomsir/MyStudyProject/Go/study/day04/No2/resource/golog.log"
	file, _ := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer file.Close()
	logger := log.New(file, "[INFO]\t", log.Ltime)
	logger.Println("打印日志信息")

}
