package types

import "reflect"

var (
	UserType                 = reflect.TypeOf(&User{})
	UsernameType             = reflect.TypeOf(&UsernameInformation{})
	MailInformationType      = reflect.TypeOf(&MailInformation{})
	RoleType                 = reflect.TypeOf(&Role{})
	RolePermissionType       = reflect.TypeOf(&RolePermission{})
	UserLoginInformationType = reflect.TypeOf(&UserLoginInformation{})
)

type User struct {
	Id          int     `json:"id" db:"id"`
	Username    string  `json:"username" db:"username"`
	MailAddress *string `json:"mail,omitempty" db:"mail"`
	FirstName   *string `json:"first_name,omitempty" db:"first_name"`
	LastName    *string `json:"last_name,omitempty" db:"last_name"`
}

type AuthenticationInformation struct {
	Id         int    `json:"id" db:"id"`
	Username   string `json:"username" db:"username"`
	SessionKey string `json:"session_key" db:"session_key"`
}

type UserLoginInformation struct {
	Id           int    `db:"id"`
	Username     string `db:"username"`
	PasswordHash string `db:"password"`
}

type UsernameInformation struct {
	Username string `db:"username"`
}

type MailInformation struct {
	Mail string `db:"mail"`
}

type Role struct {
	Id          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
}

type RolePermission struct {
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
}

type RoleIncludingPermissions struct {
	*Role
	Permissions []*RolePermission `json:"permissions" db:"permissions"`
}
