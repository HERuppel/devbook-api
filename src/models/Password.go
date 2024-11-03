package models

type Password struct {
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}
