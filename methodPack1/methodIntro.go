package main

import "fmt"

/*
A method is just a fx with a special receiver type b2n the func keyword and
the method name. The receiver can either be a struct type or non-struct type.

Syntax:
-------
func (t Type) methodName(parameter list) {}
methodName is created, with
receiver type "Type"
"t" is called as the receiver & it can be accessed within the method
*/

// Creating a method on a struct type and calling it
type Employee struct {
	firstName string
	lastName  string
	salary    float64
	currency  string
}

// I created method displaySalary on Employee struct type.
// displaySalary method has Employee as the receiver type.
// I am using the receiver employee & printing the data of the Employee
func (employee Employee) displaySalary() {
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

	// U can (1):
	employee1.displaySalary()
	employee2.displaySalary()
	fmt.Println("*********************")
	// U can (2) also:
	fmt.Println(employee1) // {Sadio Mane 450000 rwf}
	fmt.Println(employee2) // {Cristiano Ronaldo 500000 rwf}
}
