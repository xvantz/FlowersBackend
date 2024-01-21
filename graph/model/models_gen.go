// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type NewPost struct {
	Name        string `json:"name"`
	Author      string `json:"Author"`
	Description string `json:"Description"`
	Section     string `json:"section"`
}

type NewWork struct {
	Name        string `json:"name"`
	Section     string `json:"section"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Post struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"Author"`
	Description string `json:"Description"`
	Section     string `json:"section"`
	CreatedAt   int    `json:"createdAt"`
	UpdatedAt   int    `json:"updatedAt"`
}

type Query struct {
}

type Work struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Section     string `json:"section"`
	Description string `json:"description"`
	Image       string `json:"image"`
	CreatedAt   int    `json:"createdAt"`
	UpdatedAt   int    `json:"updatedAt"`
}