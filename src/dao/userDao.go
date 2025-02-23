package dao

import (
	"ebd/src/domain"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (r *UserDao) FindByID(id int64) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserDao) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserDao) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserDao) Save(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserDao) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *UserDao) Delete(id int64) error {
	return r.db.Delete(&domain.User{}, id).Error
}
