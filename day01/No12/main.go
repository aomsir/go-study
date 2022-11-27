package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	s := t.Format("2006-01-02 15:04:05")
	fmt.Println(s)

	t2, err := time.Parse("2006-01-02 15:04:05", s)
	fmt.Println(t2, err)

}
