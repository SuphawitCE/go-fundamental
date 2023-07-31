package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getTodo() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	// Send GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer response.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Parse JSON response
	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print the parsed JSON data
	fmt.Println("TODO:", todo)
	// return todo
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	// Send GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer response.Body.Close()

	// Parse JSON response
	var posts []Post
	err = json.NewDecoder(response.Body).Decode(&posts)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	getTodo()

	// Print the parsed JSON data
	fmt.Println("POST")
	for _, post := range posts {
		fmt.Println(post)
	}
}
