package model

import (
	"gorm.io/gorm"
	"time"
)

type Household struct {
	gorm.Model
	Name       string
	PostalCode string
	PhotoPath  string
}

type HouseholdAdhocTask struct {
	gorm.Model
	Name                              string
	Points                            uint8
	SpecialInstructions               string
	Household                         Household
	HouseholdId                       int `gorm:"column:household_id"`
	HouseholdMember                   HouseholdMember
	HouseholdMemberId                 int `gorm:"column:household_member_id"`
	Complete                          bool
	Rating                            float32
	DueDate                           time.Time
	ChoreCategory                     ChoreCategory
	ChoreCategoryId                   int `gorm:"column:chore_category_id"`
	FractionalPointsEnabled           bool
	HouseholdMemberPointsCollection   HouseholdMemberPointsCollection
	HouseholdMemberPointsCollectionId int `gorm:"column:points_collection_id"`
}

type HouseholdChore struct {
	gorm.Model
	Name                string
	SpecialInstructions string
	Points              uint8
	ChoreCategory       ChoreCategory
	ChoreCategoryId     int `gorm:"column:chore_category_id"`
	Chore               Chore
	ChoreId             int `gorm:"column:chore_id"`
	Household           Household
	HouseholdId         int `gorm:"column:household_id"`
}

type HouseholdChoreAssignment struct {
	gorm.Model
	HouseholdChore    HouseholdChore
	HouseholdChoreId  int `gorm:"column:household_chore_id"`
	HouseholdMember   HouseholdMember
	HouseholdMemberId int `gorm:"column:household_member_id"`
	LastCompleted     time.Time
}

type HouseholdChoreAssignmentComplete struct {
	gorm.Model
	HouseholdChoreAssignment          HouseholdChoreAssignment
	HouseholdChoreAssignmentId        int `gorm:"column:household_chore_assignment_id"`
	Rating                            float32
	Notes                             string
	VerifiedBy                        HouseholdMember
	VerifiedById                      int `gorm:"column:verified_by"`
	FractionalPointsEnabled           bool
	HouseholdMemberPointsCollection   HouseholdMemberPointsCollection
	HouseholdMemberPointsCollectionId int `gorm:"column:points_collection_id"`
}

type HouseholdChoreFrequency struct {
	gorm.Model
	HouseholdChore   HouseholdChore
	HouseholdChoreId int `gorm:"column:household_chore_id"`
	DayOfWeek        string
}

type HouseholdChorePause struct {
	gorm.Model
	Name        string
	Household   Household
	HouseholdId int `gorm:"column:household_id"`
}

type HouseholdMember struct {
	gorm.Model
	Nickname                string
	PhotoPath               string
	Household               Household
	HouseholdId             int `gorm:"column:household_id"`
	FractionalPointsEnabled bool
	User                    User
	UserId                  int `gorm:"column:user_id"`
}

type HouseholdMemberPointsCollection struct {
	gorm.Model
	Collector           HouseholdMember
	CollectorId         int `gorm:"column:collector_id"`
	Approver            HouseholdMember
	ApproverId          int `gorm:"column:approver_id"`
	ApproverNotes       string
	CollectorReflection string
	PointsCollected     float64
}
