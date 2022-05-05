package models

type Books struct {
	id     int
	title  string
	author *Authors
}

type Authors struct {
	id   int
	name string
	book *Books
}
