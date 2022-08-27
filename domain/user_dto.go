package domain

import "time"

type UserLoginRequest struct {
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}

type UserRegisterRequest struct {

}

type UserResponse struct {
	UUID			string			`json:"uuid"`
	Name			string			`json:"name"`
	Role			string			`json:"role"`
	Email			*string			`json:"email,omitempty"`
	Password		*string			`json:"password,omitempty"`
	Phone			*string			`json:"phone,omitempty"`
	WA				*string			`json:"wa,omitempty"`
	ProfileUrl		*string			`json:"profile_url,omitempty"`
	Token			*string			`json:"access_token,omitempty"`
	RefreshToken	*string			`json:"refresh_token,omitempty"`
	CreatedAt		*time.Time	 	`json:"created_at,omitempty"`
	UpdatedAt		*time.Time		`json:"updated_at,omitempty"`
	DeletedAt		*time.Time		`json:"deleted_at,omitempty"`
}
