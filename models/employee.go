package models

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Position string `json:"position"`
}

type Employees []Employee
