package util

import (
	api_model "github.com/scott-k-huang/homechores/api-model"
	"github.com/scott-k-huang/homechores/model"
)

func ChoreCategoriesToModels(choreCategories []api_model.ChoreCategory) []model.ChoreCategory {
	results := make([]model.ChoreCategory, len(choreCategories))
	for i, cat := range choreCategories {
		results[i] = ChoreCategoryToModel(cat)
	}
	return results
}

func ChoreCategoriesToAPIModel(choreCategories []model.ChoreCategory) []api_model.ChoreCategory {
	results := make([]api_model.ChoreCategory, len(choreCategories))
	for i, cat := range choreCategories {
		results[i] = ChoreCategoryToAPIModel(cat)
	}
	return results
}

func ChoreCategoryToModel(choreCategory api_model.ChoreCategory) model.ChoreCategory {
	var result model.ChoreCategory
	result.ID = choreCategory.ID
	result.Name = choreCategory.Name
	result.UpdatedAt = choreCategory.UpdatedAt
	result.CreatedAt = choreCategory.CreatedAt
	return result
}

func ChoreCategoryToAPIModel(choreCategory model.ChoreCategory) api_model.ChoreCategory {
	var result api_model.ChoreCategory
	result.Name = choreCategory.Name
	result.ID = choreCategory.ID
	result.CreatedAt = choreCategory.CreatedAt
	result.UpdatedAt = choreCategory.UpdatedAt
	return result
}

func ChoresToAPIModels(chores []model.Chore) []api_model.Chore {
	results := make([]api_model.Chore, len(chores))
	for i, chore := range chores {
		results[i] = ChoreToAPIModel(chore)
	}
	return results
}

func ChoreToAPIModel(chore model.Chore) api_model.Chore {
	var result api_model.Chore
	result.Name = chore.Name
	result.ID = chore.ID
	result.CreatedAt = chore.CreatedAt
	result.ChoreCategory = ChoreCategoryToAPIModel(chore.ChoreCategory)
	return result
}

func ChoreToModel(chore api_model.Chore) model.Chore {
	var result model.Chore
	result.Name = chore.Name
	result.ID = chore.ID
	result.CreatedAt = chore.CreatedAt
	result.ChoreCategory = ChoreCategoryToModel(chore.ChoreCategory)
	return result
}

func TransformUsersToApiModel(users []model.User) []api_model.User {
	results := make([]api_model.User, len(users))
	for i, user := range users {
		results[i] = TransformUserToApiModel(user)
	}
	return results
}

func TransformUserToApiModel(user model.User) api_model.User {
	var result api_model.User
	result.ID = user.ID
	result.FirstName = user.FirstName
	result.LastName = user.LastName
	result.CreatedAt = user.CreatedAt
	result.UpdatedAt = user.UpdatedAt
	result.Email = user.Email
	return result
}
