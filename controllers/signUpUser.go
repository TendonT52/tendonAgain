package controllers

import ()

type SignUpUser struct {
	Name     string `json:"name" bson:"name"`
	Surname  string `json:"surname" bson:"surname"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

func (signUpuser *SignUpUser) GetEmail() string {
	return signUpuser.Email
}
