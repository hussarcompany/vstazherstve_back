package client

type User struct {
	Username     string `json:"username,omitempty" bson:"username,omitempty"`
	Userpassword string `json:"userpassword,omitempty" bson:"userpassword,omitempty"`
}
