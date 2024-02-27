package model

import (
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"myAPIProject/internal/apperrors"
)

const (
	superUserAdmin              = "admin"
	superUserRoot               = "root"
	noPermissionToUpdateUserErr = "auth user can't change this target user"
)

type User struct {
	ID        uuid.UUID `json:"id" bson:"_id,omitempty"`
	Nickname  string    `json:"nickname" bson:"nickname"`
	Email     string    `json:"email" bson:"email"`
	FirstName string    `json:"first_name" bson:"first_name"`
	LastName  string    `json:"last_name" bson:"last_name"`
	Password  string    `json:"password" bson:"password"`
	Created
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
	DeletedAt primitive.DateTime `json:"deleted_at" bson:"deleted_at,omitempty"`
}

type Created struct {
	By string             `json:"created_by" bson:"created_by,omitempty"`
	At primitive.DateTime `json:"created_at" bson:"created_at"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return apperrors.UserHashGenerateFromPassword.AppendMessage(err)
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return apperrors.UserComparePasswordsCompareHashAndPassword.AppendMessage(err)
	}

	return nil
}

func (u *User) HasPermissionToUpdateUser(authUser *User) error {
	if authUser.ID == u.ID {
		return nil
	}
	if u.Created.By != "" && authUser.ID.String() == u.Created.By {
		return nil
	}
	if authUser.Nickname == superUserAdmin || authUser.Nickname == superUserRoot {
		return nil
	}

	return fmt.Errorf(noPermissionToUpdateUserErr)
}

func (u *User) TableName() string {
	return "users"
}
