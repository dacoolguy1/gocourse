package basics

import "fmt"

type EmployeeGoogle struct {
	firstName string
	lastName  string
	age       int
}

type EmployeeApple struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	//PascalCase
	// eg CalculateArea, UserInfo, NewHttpRequest
	// Struct, enum, interfaces

	// snake_case
	// eg user_id, first_name, http_request
	// variables, file names

	//UPPERCASE
	// use case is constants

	// mixedCase
	// javaScripts, htmlDocument, isValid
	// usecase variables, file names

	const MAXRETRIES = 5
	var employeeID = 1000
	fmt.Println(employeeID)
}
