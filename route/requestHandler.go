package route

import (
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()
	router.GET("/", homePage) //all good
	router.POST("/chorecategories", createChoreCategory)
	router.GET("/chorecategories", returnAllChoreCategories)
	router.GET("/chorecategories/:id", returnChoreCategory)
	router.GET("/chores", returnAllChores)
	router.GET("/chores/:id", returnChore)
	router.POST("/chores", createChore)
	router.GET("/users", returnAllUsers) //all good
	router.GET("/users/:id", returnUser) //all good
	router.POST("/users", createUser)    //all good
	router.Run()

}
