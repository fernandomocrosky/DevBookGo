package models

type Password struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}
