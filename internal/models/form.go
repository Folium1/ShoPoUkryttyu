package models

type UserRegister struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email,unique"`
	Password string `json:"password" bson:"password"`
}
