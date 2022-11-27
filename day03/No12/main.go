package main

import "os"

func main() {
	filePath := "/Users/aomsir/MyStudyProject/Go/study/day03/No12/src/demo.txt"
	file, error := os.OpenFile(filePath, os.O_APPEND, 0660) // 追更,读写权限
	defer file.Close()

	if error != nil {
		file, _ = os.Create(filePath)
	}
	file.Write([]byte("package srfffffc"))
	file.WriteString("Hello\n\thhhhh")
}
