package model

// @model CarModel
// @description Represents a car in the database.
type CarModel struct {
	CarID  int         `json:"car_id" db:"car_id"`
	RegNum string      `json:"reg_num" db:"reg_num"`
	Mark   string      `json:"mark" db:"mark"`
	Model  string      `json:"model" db:"model"`
	Year   int         `json:"year" db:"year"`
	Owner  PeopleModel `json:"owner_id" db:"owner_id"`
}

// @model PeopleModel
// @description Represents the car owner information.
type PeopleModel struct {
	OwnerID    int    `json:"owner_id" db:"owner_id"`
	Name       string `json:"name" db:"name"`
	Surname    string `json:"surname" db:"surname"`
	Patronymic string `json:"patronymic" db:"patronymic"`
}

// @model CarFilter
// @description Defines filters for searching cars.
type CarFilter struct {
	RegNum     string `json:"reg_num"`
	Mark       string `json:"mark"`
	Model      string `json:"model"`
	Year       int    `json:"year"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Limit      string `json:"limit"`
	Offset     string `json:"offset"`
}
