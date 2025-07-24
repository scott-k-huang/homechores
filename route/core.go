package route

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	api_model "github.com/scott-k-huang/homechores/api-model"
	"github.com/scott-k-huang/homechores/dao"
	"github.com/scott-k-huang/homechores/service"
	"github.com/scott-k-huang/homechores/util"
)

func createUser(c *gin.Context) {
	var userRequest api_model.CreateUserRequest
	err := c.BindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		userModel, err := dao.CreateUser(userRequest.FirstName, userRequest.LastName, userRequest.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, util.TransformUserToApiModel(*userModel))
		}
	}
}

func updateUser(c *gin.Context) {
	var userRequest api_model.User
	err := c.BindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		userModel, err := dao.FindUser(userRequest.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			userModel.LastName = userModel.LastName
			userModel.FirstName = userModel.FirstName
			userModel.Email = userModel.Email
			userModel, err := dao.UpdateUser(*userModel)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
			} else {
				c.JSON(http.StatusOK, util.TransformUserToApiModel(*userModel))
			}
		}
	}
}

func returnAllUsers(c *gin.Context) {
	users, err := dao.FindUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, util.TransformUsersToApiModel(*users))
	}
}

func returnUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("error converting to int: ", err)
		c.JSON(http.StatusInternalServerError, err)
	} else {
		user, err := dao.FindUser(uint(id))
		if err != nil {
			log.Println("error retrieving entity: ", err)
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, util.TransformUserToApiModel(*user))
		}
	}
}

func updateChore(c *gin.Context) {
	var choreRequest api_model.Chore
	err := c.BindJSON(choreRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		chore, err := dao.FindChore(choreRequest.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			chore.Name = choreRequest.Name
			chore.ChoreCategory = util.ChoreCategoryToModel(choreRequest.ChoreCategory)
			chore, err := dao.UpdateChore(*chore)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
			} else {
				c.JSON(http.StatusOK, util.ChoreToAPIModel(*chore))
			}
		}
	}
}

func createChore(c *gin.Context) {
	var choreRequest api_model.CreateChoreRequest
	err := c.BindJSON(&choreRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		chore, err := service.CreateChore(choreRequest.Name, choreRequest.ChoreCategoryID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			println("error creating chore")
			println(err)
		} else {
			c.JSON(http.StatusOK, util.ChoreToAPIModel(*chore))
		}
	}

}
func returnAllChores(c *gin.Context) {
	chores, err := dao.FindChores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, util.ChoresToAPIModels(*chores))
	}
}

func returnChore(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("error converting to int: ", err)
		c.JSON(http.StatusInternalServerError, err)
	} else {
		chore, err := dao.FindChore(uint(id))
		if err != nil {
			log.Println("error retrieving entity: ", err)
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, util.ChoreToAPIModel(*chore))
		}
	}
}

func createChoreCategory(c *gin.Context) {
	var choreCategoryRequest api_model.CreateChoreCategoryRequest
	err := c.BindJSON(&choreCategoryRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		choreCategoryModel, err := dao.CreateChoreCategory(choreCategoryRequest.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, util.ChoreCategoryToAPIModel(*choreCategoryModel))
		}
	}
}

func returnAllChoreCategories(c *gin.Context) {
	choreCategories, err := dao.FindChoreCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, util.ChoreCategoriesToAPIModel(*choreCategories))
	}
}

func returnChoreCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("error converting to int :", err)
		c.JSON(http.StatusInternalServerError, err)
	} else {
		choreCategory, err := dao.FindChoreCategory(uint(id))
		if err != nil {
			log.Println("error retrieving entity: ", err)
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, util.ChoreCategoryToAPIModel(*choreCategory))
		}
	}

}
