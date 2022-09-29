package models

import (
	"time"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {

	Name		string				`json:"name,omitempty" validate:"required"`
	Dob			string				`json:"dob,omitempty" validate:"required"`
	Address		string				`json:"address,omitempty" validate:"required"`
	Description	string				`json:"description,omitempty" validate:"required"`	
	CreatedAt	time.Time			`json:"createdAt"`

}