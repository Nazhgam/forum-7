package service

import (
	"errors"
	"fmt"
	"forum/entity"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(user *entity.User) error {
	s.log.Println("start signup")
	b, err := s.repo.CheckUsernameForExist(user.Email)
	if err != nil {
		return fmt.Errorf("error while insert new user: %v, error: %s", user, err)
	}
	s.log.Printf("count of user by email: %d", b)

	if b == 0 {
		return s.InsertNewUser(user)
	}
	return errors.New("user already exist")
}

func (s *service) InsertNewUser(user *entity.User) error {
	s.log.Println("Start insertNewUser method")
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return fmt.Errorf("invalid email: %w", err)
	}
	s.log.Println("email validation correct")
	if !IsValidPassword(user.Password) {
		return errors.New("invalid password")
	}
	s.log.Println("password valid")
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		s.log.Println("can not create hashed password")
		return err
	}
	user.Password = string(hashed)
	s.log.Println("start to insert new user to bd")

	return s.repo.CreateUser(user)
}

func IsValidPassword(password string) bool {
	// Password length should be between 8 and 30 characters
	if len(password) < 8 || len(password) > 30 {
		return false
	}

	// Password should contain at least one uppercase letter, one lowercase letter, one digit, and one special character
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWZYX0123456789!@#$%^&*()_+abcdefjhigklmnopqrstuvwxzy") {
		return false
	}
	return true
}

func (s *service) LogIn(user *entity.User) (entity.User, error) {
	dbUser, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		return entity.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return entity.User{}, err
	}

	if dbUser.Email != user.Email {
		return entity.User{}, errors.New("email otirik")
	}
	return *dbUser, nil
}
