package types

type LoginData struct {
	UsernameOrMail string `json:"username_or_mail"`
	Password       string `json:"password"`
}

type RegistrationData struct {
	Username  string  `json:"username"`
	Mail      string  `json:"mail"`
	Password  string  `json:"password"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
}
