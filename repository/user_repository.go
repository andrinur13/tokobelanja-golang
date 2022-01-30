package repository

import (
	"tokobelanja-golang/model/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(input entity.User) (entity.User, error)
	CheckSameEmail(email string) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
	GetByID(ID int) (entity.User, error)
	Update(ID int, user entity.User) (entity.User, error)
	Delete(ID int) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) CheckSameEmail(email string) (entity.User, error) {
	userSame := entity.User{}

	err := r.db.Where("email = ?", email).Find(&userSame).Error

	if err != nil {
		return entity.User{}, err
	}

	return userSame, nil
}

func (r *userRepository) GetByEmail(email string) (entity.User, error) {
	userResult := entity.User{}

	err := r.db.Where("email = ?", email).Find(&userResult).Error

	if err != nil {
		return entity.User{}, err
	}

	return userResult, nil
}

func (r *userRepository) GetByID(ID int) (entity.User, error) {
	userResult := entity.User{}

	err := r.db.Where("id = ?", ID).Find(&userResult).Error

	if err != nil {
		return entity.User{}, err
	}

	return userResult, nil
}

func (r *userRepository) Update(ID int, user entity.User) (entity.User, error) {
	err := r.db.Where("id = ?", ID).Updates(&user).Error

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) Delete(ID int) (bool, error) {
	userDeleted := entity.User{
		ID: ID,
	}

	err := r.db.Where("id = ?", ID).Delete(userDeleted).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
