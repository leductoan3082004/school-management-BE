package usermodel

import (
	"SchoolManagement-BE/appCommon"
	"errors"
	"net/http"
)

const (
	RoleUser = iota
	RoleTeacher
	RoleAdmin
)

const EntityName = "user"

type AuthenticationUser struct {
	Username string `json:"-" bson:"username"`
	Password string `json:"-" bson:"password"`
	Salt     string `json:"-" bson:"salt"`
}

type SpecUser struct {
	Role    int    `json:"role" bson:"role"`
	Name    string `json:"name" bson:"name"`
	Phone   string `json:"phone" bson:"phone"`
	Address string `json:"address" bson:"address"`
}
type User struct {
	appCommon.MgDBModel `json:",inline" bson:",inline"`
	AuthenticationUser  `json:",inline" bson:",inline"`
	SpecUser            `json:",inline" bson:",inline"`
}

func (User) TableName() string {
	return "user"
}

type UserCreate struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     int    `json:"role"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserChangePassword struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type UserList struct {
	Role *int `form:"role"`
}

var (
	ErrUsernameExisted = appCommon.NewCustomError(
		http.StatusBadRequest,
		errors.New("username has already existed"),
		"username has already existed",
		"ErrUsernameExisted",
	)
	ErrUsernameOrPasswordInvalid = appCommon.NewCustomError(
		http.StatusBadRequest,
		errors.New("username or password is invalid"),
		"username or password is invalid",
		"ErrUsernameOrPasswordInvalid",
	)
)
