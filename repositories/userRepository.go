package repositories

import (
	"goChatApp/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) Create(user *domain.User) error {
	err := u.db.Create(user).Error
	return err
}

func (u userRepository) Update(user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetById(id int64) (*domain.User, error) {
	var user domain.User
	err := u.db.Raw("SELECT * FROM go_chat_app.users WHERE id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := u.db.Raw("SELECT * FROM go_chat_app.users WHERE email = ?", email).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepository) CheckUserExists(email string) (bool, error) {
	var count int64
	err := u.db.Table("users").Where("email = ?", email).Count(&count).Error
	if err != nil {
		return true, err
	}
	return count > 0, nil
}

func (u userRepository) List() ([]*domain.User, error) {
	var users []*domain.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u userRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *gorm.DB) domain.UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}
