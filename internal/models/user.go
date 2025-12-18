package models

import "time"

type UserBasic struct {
	ID   int       `json:"id"`
	Name string    `json:"string"`
	Dob  time.Time `json:"dob"`
}

type UserRead struct {
	ID   int       `json:"id"`
	Name string    `json:"string"`
	Dob  time.Time `json:"dob"`
	Age  int       `json:"age"`
}
