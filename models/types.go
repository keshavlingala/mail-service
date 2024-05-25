package models

type PortfolioBody struct {
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Content  string      `json:"content"`
	UserData interface{} `json:"userData"`
}
