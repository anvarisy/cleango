package domains

import "time"

type Users []User

type User struct {
	ID           string    `json:"id"`
	FullName     string    `json:"full_name"`
	PlaceOfBirth string    `json:"pob"`
	DateOfBirth  time.Time `json:"dob"`
	Avatar       string    `json:"avatar"`
}
