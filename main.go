package main

import "github.com/scott-k-huang/homechores/route"

func main() {
	//fmt.Println("hello world")
	//user, err := service.CreateUser("Scott", "Huang", "scott.huang@protonmail.com")
	//if err != nil {
	//	fmt.Println("bombed creating user: ", err)
	//} else {
	//	fmt.Println("newly creeated user successful; ", user.ID)
	//}

	//choreCategory, err := service.CreateChoreCategory("Take Out")
	//if err != nil {
	//	fmt.Println("bombed creating ChoreCategory: ", err)
	//} else {
	//	fmt.Println("newly created ChoreCategory successful: ", choreCategory.ID)
	//}

	//choreCategory, err := service.FindChoreCategory(1)
	//if err != nil {
	//	fmt.Println("bombed finding Chore")
	//} else {
	//	fmt.Println("found ChoreCategory: ", choreCategory.Name)
	//}

	//chore, err := service.CreateChore("Clean Toilet", *choreCategory)
	//if err != nil {
	//	fmt.Println("bombed creating Chore: ", err)
	//} else {
	//	fmt.Println("newly created Chore successful: ", chore.ID)
	//}

	route.HandleRequests()
}
