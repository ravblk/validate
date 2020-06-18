package main

import (
	"errors"
	"regexp"
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"
)

var (
	errType = errors.New("wrong type")
	user    = User{
		ID:       1,
		Name:     "",
		Age:      30,
		Password: "123456789",
		Email:    "user@mail.ru",
	}

	regexpEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type User struct {
	ID       int    `json:"id" validate:"required" valid:"required"`
	Name     string `json:"name" validate:"required,min=2,max=100" valid:"required,length(2|100)"`
	Age      int    `json:"age" validate:"required" valid:"required"`
	Password string `json:"password" validate:"required,min=8,max=100" valid:"required,length(8|100)"`
	Email    string `json:"email" validate:"required,email" valid:"required,email"`
}

func (u User) Validate() (errs []error) {
	if user.ID == 0 {
		errs = append(errs, errors.New("wrong id"))
	}

	if len(u.Name) < 2 || len(u.Name) > 100 {
		errs = append(errs, errors.New("wrong name"))
	}

	if user.Age == 0 {
		errs = append(errs, errors.New("wrong age"))
	}

	if len(user.Password) < 2 && len(user.Password) > 100 {
		errs = append(errs, errors.New("wrong pass"))
	}

	if !regexpEmail.MatchString(user.Email) {
		errs = append(errs, errors.New("wrong email"))
	}

	return
}

func BenchmarkAsaskevichGovalidator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		govalidator.ValidateStruct(user)
	}
}
func BenchmarkGoPlaygroundValidator(b *testing.B) {
	validate := validator.New()
	for i := 0; i < b.N; i++ {
		validate.Struct(user)
	}
}

func BenchmarkSimpleValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user.Validate()
	}
}
