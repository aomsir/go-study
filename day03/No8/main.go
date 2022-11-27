package main

import (
	"fmt"
	"os/user"
)

func main() {
	user, error := user.Current()
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(user.Uid)
	fmt.Println(user.Name)
	fmt.Println(user.Gid)
	fmt.Println(user.HomeDir)
	fmt.Println(user.Username)
}
