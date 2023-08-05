package main

import "fmt"

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

// TODO: Initial Employee struct value or like class constructor
func (this *Employee) New(firstname string, lastname string, totalLeaves int, leavesTaken int) {
	this.FirstName = firstname
	this.LastName = lastname
	this.TotalLeaves = totalLeaves
	this.LeavesTaken = leavesTaken
}

func main() {
	// TODO: Struct
	employee := new(Employee)
	employee.New("Bank", "Developer", 15, 0)
	fmt.Println("Employee:", *employee)
}
