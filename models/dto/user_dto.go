package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDTO struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username     string             `json:"username,omitempty" bson:"username,omitempty"`
	DOB          string             `json:"dob,omitempty" bson:"dob,omitempty"`
	Useremail    string             `json:"useremail,omitempty" bson:"useremail,omitempty"`
	Userpassword string             `json:"userpassword,omitempty" bson:"userpassword,omitempty"`
}
