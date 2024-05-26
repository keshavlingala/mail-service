package models

type PortfolioBody struct {
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Content  string      `json:"content"`
	UserData interface{} `json:"userData"`
}

type SwasthaBody struct {
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
