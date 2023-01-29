package main

import (
	"fmt"
	"github.com/scott-k-huang/homechores/service"
)

func main() {
	fmt.Println("hello world")
	user, err := service.CreateUser("Scott", "Huang", "scott.huang@protonmail.com")
	if err != nil {
		fmt.Println("bombed creating user: ", err)
	} else {
		fmt.Println("newly creeated user successful; ", user.ID)
	}
	//route.HandleRequests()
}
