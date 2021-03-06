package users

import (
	"github.com/defintly/backend/database"
	"github.com/defintly/backend/types"
)

func ChangeUsername(user *types.User, newUsername string) error {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.UsernameType,
		"SELECT username FROM users WHERE username = $1", newUsername)
	if err != nil {
		return err
	}

	users := slice.([]*types.UsernameInformation)
	if len(users) != 0 {
		return UserAlreadyExists
	}

	err = database.PrepareAsync(database.DefaultTimeout, "UPDATE users SET username = $1 WHERE id = $2",
		newUsername, user.Id)
	if err != nil {
		return err
	}

	user.Username = newUsername
	return nil
}

func ChangeFirstName(user *types.User, newFirstName string) error {
	err := database.PrepareAsync(database.DefaultTimeout, "UPDATE users SET first_name = $1 WHERE id = $2",
		newFirstName, user.Id)
	if err != nil {
		return err
	}

	user.FirstName = &newFirstName
	return nil
}

func ChangeLastName(user *types.User, newLastName string) error {
	err := database.PrepareAsync(database.DefaultTimeout, "UPDATE users SET last_name = $1 WHERE id = $2",
		newLastName, user.Id)
	if err != nil {
		return err
	}

	user.LastName = &newLastName
	return nil
}

func ChangeMailAddress(user *types.User, newMailAddress string) error {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.MailInformationType,
		"SELECT mail FROM users WHERE mail = $1", newMailAddress)
	if err != nil {
		return err
	}

	users := slice.([]*types.MailInformation)
	if len(users) != 0 {
		return MailAlreadyInUse
	}

	err = database.PrepareAsync(database.DefaultTimeout, "UPDATE users SET mail = $1 WHERE id = $2",
		newMailAddress, user.Id)
	if err != nil {
		return err
	}

	user.MailAddress = &newMailAddress
	return nil
}

func ChangePassword(user *types.User, oldPassword string, newPassword string) error {
	userInformation, err := getUserLoginInformation(user.Username)
	if err != nil {
		return err
	}

	if err = compareHashAndPassword(userInformation.PasswordHash, oldPassword); err != nil {
		return IncorrectPassword
	}

	newPasswordHash, err := hashPassword(newPassword)
	if err != nil {
		return err
	}

	return database.PrepareAsync(database.DefaultTimeout, "UPDATE users SET password = $1 WHERE id = $2",
		newPasswordHash, user.Id)
}

func GetUserById(showMail bool, id int) (*types.User, error) {
	query := "SELECT id, username, first_name, last_name"
	if showMail {
		query += ", mail"
	}
	query += " FROM users WHERE id = $1"

	slice, err := database.QueryAsync(database.DefaultTimeout, types.UserType, query, id)

	if err != nil {
		return nil, err
	}

	users := slice.([]*types.User)
	if len(users) == 0 {
		return nil, UserNotFound
	}

	return users[0], err
}
