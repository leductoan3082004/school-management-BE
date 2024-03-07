package usermodel

import (
	"SchoolManagement-BE/appCommon"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type UserRole int

const (
	RoleStudent UserRole = iota
	RoleTeacher
	RoleAdmin
)

var allUserRoles = [3]string{"Student", "Teacher", "Admin"}

func (role *UserRole) String() string {
	return allUserRoles[*role]
}

func parseStr2UserRole(s string) (UserRole, error) {
	for i := range allUserRoles {
		if allUserRoles[i] == s {
			return UserRole(i), nil
		}
	}

	return UserRole(0), appCommon.ErrInternal(errors.New("invalid user role string"))
}

func (role *UserRole) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return appCommon.ErrInternal(errors.New(fmt.Sprintf("fail to scan data from sql: %s", value)))
	}

	v, err := parseStr2UserRole(string(bytes))

	if err != nil {
		return appCommon.ErrInternal(errors.New(fmt.Sprintf("fail to scan data from sql: %s", value)))
	}

	*role = v

	return nil
}

func (role *UserRole) Value() (driver.Value, error) {
	if role == nil {
		return nil, appCommon.ErrInternal(errors.New("invalid user role"))
	}

	return role.String(), nil
}

func (role *UserRole) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", role.String())), nil
}

func (role *UserRole) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	roleValue, err := parseStr2UserRole(str)

	if err != nil {
		return err
	}

	*role = roleValue

	return nil
}
