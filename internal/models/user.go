package models

type UserBasic struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Dob  Date   `json:"dob"`
}

type UserRead struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Dob  Date   `json:"dob"`
	Age  int    `json:"age"`
}

type UserCreateRequest struct {
	Name string `json:"name" validate:"required"`
	Dob  *Date  `json:"dob" validate:"required"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
	Dob  *Date  `json:"dob"`
}
