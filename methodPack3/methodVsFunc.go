package main

import "fmt"

// Rewriting methodIntro.go using only fx(s)&without methods.
// Creating a method on a struct type and calling it
type Employee struct {
	firstName string
	lastName  string
	salary    float64
	currency  string
}

// Using method (way)
func (employee Employee) displaySalary() {
	fmt.Printf("Employee name: %s %s\n", employee.firstName, employee.lastName)
	fmt.Printf("Employee salary: %.2f%s\n", employee.salary, employee.currency)
}

// Using Function (only)
// This method converted to function with Employee as parameter
func displaySalary2(employee Employee) {
	fmt.Printf("Employee name: %s %s\n", employee.firstName, employee.lastName)
	fmt.Printf("Employee salary: %.2f%s\n", employee.salary, employee.currency)
}

func main() {

	// Testing my Go program: Gad Iradufasha
	employee1 := Employee{
		firstName: "Sadio",
		lastName:  "Mane",
		salary:    450000,
		currency:  "rwf",
	}

	employee2 := Employee{
		firstName: "Cristiano",
		lastName:  "Ronaldo",
		salary:    500000,
		currency:  "rwf",
	}

	// 1. Calling displaySalary() method of Employee type
	employee1.displaySalary()
	employee2.displaySalary()
	fmt.Println("*************************")

	// 2. Calling displaySalary2() function
	displaySalary2(employee1)
	displaySalary2(employee2)
}
