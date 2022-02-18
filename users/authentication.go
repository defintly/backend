package users

import (
	"github.com/andskur/argon2-hashing"
	"github.com/defintly/backend/database"
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/types"
	"github.com/google/uuid"
)

func hashPassword(password string) (string, error) {
	hash, err := argon2.GenerateFromPassword([]byte(password), argon2.DefaultParams)

	return string(hash), err
}

func compareHashAndPassword(hash string, password string) error {
	return argon2.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GetUserByAuthenticationKey(key string) (*types.User, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.UserType,
		"SELECT users.id, users.username, users.mail, users.first_name, users.last_name "+
			"FROM user_sessions, users "+
			"WHERE user_sessions.session_key = $1 AND users.id = users_sessions.user_id",
		key)

	if err != nil {
		return nil, err
	}

	users := slice.([]*types.User)
	if len(users) == 0 {
		return nil, UserNotFound
	}

	return users[0], err
}

func Login(nameOrMail string, password string, userAgent string) (*types.AuthenticationInformation, error) {
	ok, nameOrMail := general.IsNameValid(nameOrMail, false)
	if !ok {
		return nil, InvalidMailAddressOrUsername
	}

	userInformation, err := getUserLoginInformation(nameOrMail)
	if err != nil {
		return nil, err
	}

	if err = compareHashAndPassword(userInformation.PasswordHash, password); err != nil {
		return nil, IncorrectPassword
	}

	sessionKey, err := getNewSessionKey()
	if err != nil {
		return nil, err
	}

	err = database.PrepareAsync(database.DefaultTimeout,
		"INSERT INTO user_sessions VALUES ($1, $2, $3, NOW())", userInformation.Id, sessionKey, userAgent)

	if err != nil {
		return nil, err
	}

	return &types.AuthenticationInformation{
		Id:         userInformation.Id,
		Username:   userInformation.Username,
		SessionKey: sessionKey,
	}, nil
}

func Register(name string, mailAddress string, password string, firstName *string,
	lastName *string, userAgent string) (*types.AuthenticationInformation, error) {
	ok, mailAddress := general.IsEmailValid(mailAddress)
	if !ok {
		return nil, InvalidMailAddressOrUsername
	}

	ok, name = general.IsNameValid(name, true)
	if !ok {
		return nil, InvalidMailAddressOrUsername
	}

	slice, err := database.QueryAsync(database.DefaultTimeout, types.IdInformationType,
		"SELECT id FROM users WHERE mail = $1 OR username = $2", mailAddress, name)
	if err != nil {
		return nil, err
	}

	if len(slice.([]*types.IdInformation)) != 0 {
		return nil, UserAlreadyExists
	}

	passwordHash, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	err = database.PrepareAsync(database.DefaultTimeout,
		"INSERT INTO users(username, mail, first_name, last_name, password) VALUES($1, $2, $3, $4, $5)",
		name, mailAddress, firstName, lastName, passwordHash)
	if err != nil {
		return nil, err
	}

	return Login(mailAddress, password, userAgent)
}

func getNewSessionKey() (string, error) {
	var key string
	var err error

	for {
		key = uuid.New().String()

		slice, err := database.QueryAsync(database.DefaultTimeout, types.IdInformationType,
			"SELECT id FROM user_sessions WHERE session_key = $1", key)
		if err != nil {
			break
		}

		if len(slice.([]*types.IdInformation)) == 0 {
			break
		}
	}

	return key, err
}

func getUserLoginInformation(nameOrMail string) (*types.UserLoginInformation, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.UserLoginInformationType,
		"SELECT id, username, password FROM users WHERE mail = $1 OR username = $2", nameOrMail, nameOrMail)
	if err != nil {
		return nil, err
	}

	information := slice.([]*types.UserLoginInformation)
	if len(information) == 0 {
		return nil, UserNotFound
	}

	return information[0], nil
}
