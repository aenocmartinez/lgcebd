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

type UserData struct {
	ID           int64
	Username     string
	Email        string
	Password     string
	SessionToken string
	Name         string
}

func mapToUser(userData *UserData) *domain.User {
	user := domain.NewUser(nil)
	user.SetID(userData.ID)
	user.SetUsername(userData.Username)
	user.SetEmail(userData.Email)
	user.SetPassword(userData.Password)
	user.SetSessionToken(userData.SessionToken)
	user.SetName(userData.Name)
	return user
}

func (r *UserDao) FindByID(id int64) (*domain.User, error) {
	var userData UserData
	result := r.db.Table("users").Where("id = ?", id).First(&userData)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return domain.NewUser(nil), nil
		}
		return nil, result.Error
	}
	return mapToUser(&userData), nil
}

func (r *UserDao) FindByEmail(email string) (*domain.User, error) {
	var userData UserData
	result := r.db.Table("users").Where("email = ?", email).First(&userData)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return domain.NewUser(nil), nil
		}
		return nil, result.Error
	}
	return mapToUser(&userData), nil
}

func (r *UserDao) FindByUsername(username string) (*domain.User, error) {
	var userData UserData
	result := r.db.Table("users").Where("username = ?", username).First(&userData)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return domain.NewUser(nil), nil
		}
		return nil, result.Error
	}
	return mapToUser(&userData), nil
}

func (r *UserDao) Save(user *domain.User) error {
	userData := UserData{
		ID:           user.GetID(),
		Username:     user.GetUsername(),
		Email:        user.GetEmail(),
		Password:     user.GetPassword(),
		SessionToken: user.GetSessionToken(),
		Name:         user.GetName(),
	}
	return r.db.Table("users").Create(&userData).Error
}

func (r *UserDao) Update(user *domain.User) error {
	userData := UserData{
		ID:           user.GetID(),
		Username:     user.GetUsername(),
		Email:        user.GetEmail(),
		Password:     user.GetPassword(),
		SessionToken: user.GetSessionToken(),
		Name:         user.GetName(),
	}
	return r.db.Table("users").Where("id = ?", user.GetID()).Updates(&userData).Error
}

func (r *UserDao) Delete(id int64) error {
	return r.db.Table("users").Where("id = ?", id).Delete(&domain.User{}).Error
}
