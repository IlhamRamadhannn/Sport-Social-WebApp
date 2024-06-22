package user

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]UserInput, error)
	FindByID(ID int) (UserInput, error)
	Create(user UserInput) (UserInput, error)
	Update(user UserInput) (UserInput, error)
	Delete(user UserInput) (UserInput, error)
}

type repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]UserInput, error){
	var users []UserInput

	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) FindByID(ID int) (UserInput, error){
	var user UserInput

	err := r.db.Find(&user).Error

	return user, err
}

func (r *repository) Create(user UserInput) (UserInput, error){
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Update(user UserInput) (UserInput, error){
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) Delete(user UserInput) (UserInput, error){
	err := r.db.Delete(&user).Error

	return user, err
}
