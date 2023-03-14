package main

import (
	"fmt"
	"ginchat/service"
)

func main() {
	users, err := service.FindBy("id", 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}
