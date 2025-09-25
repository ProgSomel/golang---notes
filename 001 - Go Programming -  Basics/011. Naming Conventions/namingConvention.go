package main

type EmployeGoogle struct{
	FirstName string
	LastName string
	Age int
}

func main(){
	//! PascalCase---> for naming types
	//? structs, interface, enums
	//Ex: CalculateArea, UserInfo, NewHTTPRequest

	//! snake_case
	//Ex: user_id, first_name, http_request

	//! UPPERCASE ---> for naming constants(variables are mutable but constants are immutable)

	//! mixedCase
	//Ex: javaScript, htmlDocument, isValid

	const MAXRETRIES = 5;
	var employeeID  = 1001;

}