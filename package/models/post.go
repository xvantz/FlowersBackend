package models

type Post struct {
	BaseModel
	Name        string
	Author      string
	Description string
	Section   	string
}