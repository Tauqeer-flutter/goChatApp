package repositories

import (
	"goChatApp/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (u UserRepository) Create(user *domain.User) error {
	err := u.db.Create(user).Error
	return err
}

func (u UserRepository) Update(user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) GetById(id string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := u.db.Raw("SELECT * FROM go_chat_app.users WHERE email = ?", email).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserRepository) CheckUserExists(email string) (bool, error) {
	var count int64
	err := u.db.Table("users").Where("email = ?", email).Count(&count).Error
	if err != nil {
		return true, err
	}
	return count > 0, nil
}

func (u UserRepository) List() ([]*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
