package models

type Book struct {
	Id          int     `json:"id" `
	Title       string  `json:"title"`
	PublishDate string  `json:"publishDate"`
	Author      *Author `json:"author"`
}

//db BOok id,title,pbdate,author(*author name)
//db Author id, name, book (*book title)

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Book *Book  `json:"book"`
}

func (a Author) GetBook() {
	/*a.Name
	a.Book.PublishDate
	a.Book.Title*/
	//stmt := "select title,publishdate from books where ? <= a.Name"
}
func (b Book) GetAuthor() {
	//stmt := "select ? <= b.Author.Name from books where title=? <= b.Title"
}
