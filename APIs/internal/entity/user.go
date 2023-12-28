package entity

import (
	"github.com/danubiobwm/goexpert/APIs/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(Name string, Email string, Password string) (*User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		Name:     Name,
		Email:    Email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(Password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(Password))

	return err == nil
}
