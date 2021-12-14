package usecases

import "time"

type UserRequest struct {
	FullName     string    `json:"full_name"`
	PlaceOfBirth string    `json:"pob"`
	DateOfBirth  time.Time `json:"dob"`
	Avatar       string    `json:"avatar"`
}
