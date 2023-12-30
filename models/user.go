package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username       string `json:"username"`
	Email          string `json:"email"`
	PasswordDigest string `json:"password_digest"`
	LastLogin      string `json:"last_login"`
}

type UserCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserForm struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepositoryGorm struct {
	db *gorm.DB
}

func NewUserRepositoryGorm(db *gorm.DB) *UserRepositoryGorm {
	return &UserRepositoryGorm{db}
}

func (r *UserRepositoryGorm) CreateUser(user *User, password string) error {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	user.PasswordDigest = hashedPassword

	result := r.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepositoryGorm) SignIn(username, password string) (*User, error) {
	var user User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	if err != nil {
		return nil, err
	}

	user.LastLogin = time.Now().Format(time.RFC3339)
	r.db.Save(&user)

	return &user, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
