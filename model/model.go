package model

// User ...
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Response string `json:"response"`
}
