package models


type Book struct {
	Id          int     `json:"id" `
	Title       string  `json:"title"`
	PublishDate string  `json:"publishDate"`
	Author      *Author `json:"author"`
}


type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Book *Book  `json:"book"`
}