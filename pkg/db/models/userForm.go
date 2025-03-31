/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package models

type UserForm struct {
	FormBase
	Username  string `json:"username" validate:"required,min=2,max=50"`
	Password  string `json:"password" validate:"required,min=2,max=80"`
	Firstname string `json:"firstname" validate:"required,min=2,max=80"`
	Lastname  string `json:"lastname" validate:"required,min=2,max=80"`
	Email     string `json:"email" validate:"required,min=2,max=80"`
}

func (form *UserForm) MapToModel() *User {
	return &User{
		Username:  form.Username,
		Password:  form.Password,
		Firstname: form.Firstname,
		Lastname:  form.Lastname,
		Email:     form.Email,
	}
}
