package user

import (
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	Login(user LoginRequest) (result User, err error)
	SignUp(user User) (err error)
	Update(user User) (err error)
	Delete(user User) (err error)
	GetList() (users []User, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(database *gorm.DB) Repository {
	return &userRepository{
		db: database,
	}
}

func (r *userRepository) Login(user LoginRequest) (result User, err error) {
	err = r.db.Where("username = ?", user.Username).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return result, err
	}

	return result, nil
}

func (r *userRepository) SignUp(user User) (err error) {
	err = r.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Update(user User) (err error) {
	err = r.db.Model(&User{}).Where("username = ?", user.Username).Update("full_name", user.FullName).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(user User) (err error) {
	err = r.db.Where("username = ?", user.Username).Delete(&User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetList() (users []User, err error) {
	err = r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
