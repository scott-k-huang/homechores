package dao

import (
	"log"
	"time"

	"github.com/scott-k-huang/homechores/config"
	"github.com/scott-k-huang/homechores/model"
)

func CreateHousehold(name string, postalCode string, photoPath string) (*model.Household, error) {
	db := config.DB
	var household model.Household
	household.Name = name
	household.PostalCode = postalCode
	household.PhotoPath = photoPath
	tx := db.Create(&household)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		log.Println("Newly created Household ID: ", household.ID)
		return &household, nil
	}
}

func UpdateHousehold(household model.Household) (*model.Household, error) {
	db := config.DB
	tx := db.Save(household)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		return &household, nil
	}
}

func CreateHouseholdMember(nickname string, photoPath string, fractionalPointsEnabled bool, user model.User, household model.Household) (*model.HouseholdMember, error) {
	db := config.DB
	var householdMember model.HouseholdMember
	householdMember.User = user
	householdMember.Household = household
	householdMember.Nickname = nickname
	householdMember.PhotoPath = photoPath
	householdMember.FractionalPointsEnabled = fractionalPointsEnabled
	tx := db.Create(&householdMember)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		log.Println("Newly created HouseholdMember ID: ", householdMember.ID)
		return &householdMember, nil
	}
}

func UpdateHoseholdMember(householdMember model.HouseholdMember) (*model.HouseholdMember, error) {
	db := config.DB
	tx := db.Save(householdMember)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		return &householdMember, nil
	}
}

func CreateHouseholdMemberPointsCollection(collector model.HouseholdMember, approver model.HouseholdMember, approverNotes string, collectorReflection string, pointsCollected float64) (*model.HouseholdMemberPointsCollection, error) {
	db := config.DB
	var householdMemberPointsCollection model.HouseholdMemberPointsCollection
	householdMemberPointsCollection.Collector = collector
	householdMemberPointsCollection.Approver = approver
	householdMemberPointsCollection.ApproverNotes = approverNotes
	householdMemberPointsCollection.CollectorReflection = collectorReflection
	householdMemberPointsCollection.PointsCollected = pointsCollected
	tx := db.Create(&householdMemberPointsCollection)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		log.Println("Newly created HouseholdMemberPointsCollection ID: ", householdMemberPointsCollection.ID)
		return &householdMemberPointsCollection, nil
	}
}

func UpdateHouseholdMemberPointsCollection(householdMemberPointsCollection model.HouseholdMemberPointsCollection) (*model.HouseholdMemberPointsCollection, error) {
	db := config.DB
	tx := db.Save(householdMemberPointsCollection)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		return &householdMemberPointsCollection, nil
	}
}

func CreateHouseholdAdhocTask(name string, points uint8, specialInstructions string, household model.Household, householdMember model.HouseholdMember, dueDate time.Time, fractionalPointsEnabled bool) (*model.HouseholdAdhocTask, error) {
	db := config.DB
	var householdAdhocTask model.HouseholdAdhocTask
	householdAdhocTask.Name = name
	householdAdhocTask.Points = points
	householdAdhocTask.SpecialInstructions = specialInstructions
	householdAdhocTask.Household = household
	householdAdhocTask.HouseholdMember = householdMember
	householdAdhocTask.DueDate = dueDate
	householdAdhocTask.FractionalPointsEnabled = fractionalPointsEnabled
	tx := db.Create(&householdAdhocTask)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		log.Println("Newly created HouseholdAdhocTask ID: ", householdAdhocTask.ID)
		return &householdAdhocTask, nil
	}
}

func UpdateHouseholdAdhocTask(householdAdhocTask model.HouseholdAdhocTask) (*model.HouseholdAdhocTask, error) {
	db := config.DB
	tx := db.Save(householdAdhocTask)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		return &householdAdhocTask, nil
	}
}

func CreateHouseholdChore(name string, points uint8, specialInstructions string, choreCategory model.ChoreCategory, chore model.Chore, household model.Household) (*model.HouseholdChore, error) {
	db := config.DB
	var householdChore model.HouseholdChore
	householdChore.Name = name
	householdChore.Points = points
	householdChore.SpecialInstructions = specialInstructions
	householdChore.ChoreCategory = choreCategory
	householdChore.Chore = chore
	householdChore.Household = household
	tx := db.Create(&householdChore)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		log.Println("Newly created HouseholdChore ID: ", householdChore.ID)
		return &householdChore, nil
	}
}

func UpdateHouseholdChore(householdChore model.HouseholdChore) (*model.HouseholdChore, error) {
	db := config.DB
	tx := db.Save(householdChore)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		return &householdChore, nil
	}
}

func CreateHouseholdChoreAssignment(householdChore model.HouseholdChore, householdMember model.HouseholdMember) (*model.HouseholdChoreAssignment, error) {
	db := config.DB
	var householdChoreAssignment model.HouseholdChoreAssignment
	householdChoreAssignment.HouseholdChore = householdChore
	householdChoreAssignment.HouseholdMember = householdMember
	tx := db.Create(&householdChoreAssignment)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		log.Println("Newly created HouseholdChoreAssignment ID: ", householdChoreAssignment.ID)
		return &householdChoreAssignment, nil
	}
}

func UpdateHouseholdChoreAssignment(householdChoreAssignment model.HouseholdChoreAssignment) (*model.HouseholdChoreAssignment, error) {
	db := config.DB
	tx := db.Save(householdChoreAssignment)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		return &householdChoreAssignment, nil
	}
}
