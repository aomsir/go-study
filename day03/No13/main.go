package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := os.Open("/Users/aomsir/MyStudyProject/Go/study/day03/No13/resource/go.txt")
	buffer, _ := ioutil.ReadAll(file)
	fmt.Println(string(buffer))

	buffer, _ = ioutil.ReadFile("/Users/aomsir/MyStudyProject/Go/study/day03/No13/resource/go.txt")
	fmt.Println(string(buffer))
}
