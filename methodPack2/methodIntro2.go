package main

import "fmt"

// Creating a method on a struct type and calling it
type User struct {
	username string
	gender   string
	age      int8
}

// This method has User as the receiver type
func (user User) displayUserData() {
	if user.gender == "F" {
		fmt.Println("She is ", user.username, " & ", user.age, "years old")
	} else {
		fmt.Println("He is ", user.username, " & ", user.age, "years old")
	}
}

// This method has User as the receiver type as well
func (user User) getAgeRangeByGender() {
	if user.age >= 1 && user.age <= 10 {
		fmt.Println("A user is still young!")
	} else if user.age >= 10 && user.age <= 19 {
		fmt.Println("A user is just teenager!")
	} else if user.age >= 20 && user.age <= 50 {
		fmt.Println("A user is adult person!")
	} else {
		fmt.Println("Old or Invalid")
	}
}

func main() {

	// Testing
	user1 := User{
		username: "NastyC",
		gender:   "M",
		age:      17,
	}

	user2 := User{
		username: "Rebecca",
		gender:   "F",
		age:      21,
	}

	user3 := User{
		username: "Flora",
		gender:   "F",
		age:      9,
	}

	// Print data for user1 
	user1.displayUserData()     // He is  NastyC  &  17 years old
	user1.getAgeRangeByGender() // A user is just teenager!

	// Print data for user1
	user2.displayUserData()     // She is  Rebecca  &  21 years old
	user2.getAgeRangeByGender() // A user is adult person!

	// Print data for user3
	user3.displayUserData()     // She is  Flora  &  9 years old
	user3.getAgeRangeByGender() // A user is still young!
}
