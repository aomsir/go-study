package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Hello World")
	fmt.Fprint(os.Stdout, "内容1")
	fmt.Fprintln(os.Stdout, "内容2")
	fmt.Fprintf(os.Stdout, "内容3")
}
