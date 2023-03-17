package route

import (
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()
	router.GET("/", homePage)
	router.GET("/chorecategories", returnAllChoreCategories)
	router.GET("/chorecategories/:id", returnChoreCategory)
	router.GET("/chores", returnAllChores)
	router.GET("/chores/:id", returnChore)
	router.POST("/chores", createChore)
	router.GET("/users", returnAllUsers)
	router.GET("/users/:id", returnUser)
	router.POST("/users", createUser)
	router.Run()

}
