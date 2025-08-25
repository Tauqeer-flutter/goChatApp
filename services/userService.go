package services

import (
	"github.com/golang-jwt/jwt/v5"
	"goChatApp/domain"
	"goChatApp/middlewares"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type UserService struct {
	userRepository domain.UserRepositoryInterface
}

func (u UserService) SignUp(user *domain.User) error {
	exists, err := u.userRepository.CheckUserExists(user.Email)
	if err != nil || exists {
		return domain.ErrUserAlreadyExists
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	err = u.userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserService) Login(email string, password string) (*domain.User, error) {
	user, err := u.userRepository.GetByEmail(email)
	if err != nil || user == nil {
		return nil, domain.ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, domain.ErrInvalidCredentials
	}
	return user, nil
}

func (u UserService) GenerateJWT(user *domain.User) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	claims := middlewares.JWTClaims{
		UserId:  user.Id,
		Email:   user.Email,
		Expiry:  time.Now().Add(time.Hour).Unix(),
		Created: time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func NewUserService(userRepository domain.UserRepositoryInterface) domain.UserServiceInterface {
	return UserService{
		userRepository: userRepository,
	}
}
