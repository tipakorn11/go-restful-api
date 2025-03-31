/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package models

var user *User = &User{}

func UserModel() *User {
	return user
}

type User struct {
	ModelBase
	Username  string `gorm:"size:255"`
	Password  string `gorm:"size:255"`
	Firstname string `gorm:"size:255`
	Lastname  string `gorm:"size:255`
	Email     string `gorm:"size:255`
}

func (model *User) MapToForm() *UserForm {
	form := &UserForm{
		Username:  model.Username,
		Password:  model.Password,
		Firstname: model.Firstname,
		Lastname:  model.Lastname,
		Email:     model.Email,
	}
	form.ID = model.ID
	form.CreatedAt = model.CreatedAt
	form.UpdatedAt = model.UpdatedAt
	return form
}

func (model *User) FindAll() (models []*User, err error) {
	result := db.Model(model).Find(&models)
	return models, result.Error
}

func (model *User) FindMany(ids []string) (models []*User, err error) {
	result := db.Model(model).Find(&models, ids)
	return models, result.Error
}

func (model *User) Find(id string) (m *User, err error) {
	result := db.Model(model).Where("ID=?", id).First(&m)
	return m, result.Error
}

func (model *User) Save() error {
	return db.Model(model).Create(&model).Error
}

func (model *User) Update() error {
	return db.Model(model).Updates(&model).Error
}

func (model *User) Delete(id string) error {
	return db.Model(model).Where("ID=?", id).Delete(&model).Error
}
