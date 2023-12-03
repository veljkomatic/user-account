package handler

import (
	"github.com/pkg/errors"
	"html/template"
	"net/http"

	"github.com/veljkomatic/user-account/pkg/user/domain"
)

type RequestValidator interface {
	ValidateRequest() error
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	User *domain.User `json:"user"`
}

func (CreateUserResponse) StatusCode() int {
	return http.StatusCreated
}

var _ RequestValidator = (*CreateUserRequest)(nil)

func (r *CreateUserRequest) ValidateRequest() error {
	// validate name
	if err := validateName(r.Name); err != nil {
		return err
	}

	// safeguard against XSS
	r.Name = template.HTMLEscapeString(r.Name)
	return nil
}

type GetUserRequest struct {
	UserID domain.UserID `uri:"userID" binding:"required,uuid"`
}

type GetUserResponse struct {
	User *domain.User `json:"user"`
}

func (GetUserResponse) StatusCode() int {
	return http.StatusOK
}

func validateName(name string) error {
	if len(name) < 1 || len(name) > 50 {
		return errors.New("name must be between 1 and 50 characters")
	}
	return nil
}
