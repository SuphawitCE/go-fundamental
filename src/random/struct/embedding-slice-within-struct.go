package main

import "fmt"

type Author struct {
	Firstname, Lastname, bio string
}

type Book struct {
	ISBN, Title, Edition, Publisher string
	Authors                         []Author
	Price                           float64
}

func main() {
	authors := []Author{
		{
			Firstname: "Bank",
			Lastname:  "Developer",
			bio:       "Winning Nobel prize",
		},
		{Firstname: "Lorem", Lastname: "Ipsum", bio: "Best Developer"},
		{Firstname: "John", Lastname: "Doe", bio: "Placeholder name"},
	}

	fmt.Println("authors", authors)

	book1 := Book{
		ISBN:      "1234",
		Title:     "Introduction to Go Programming",
		Edition:   "1st edition",
		Publisher: "2020",
		Authors:   authors,
		Price:     120.55,
	}

	fmt.Println("Book1", book1)
}
