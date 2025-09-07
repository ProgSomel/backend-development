package main

import "fmt"

type customError struct {
	code int
	message string
}

// Implementing Error() method of error interface 
func (e customError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.message)
}

func doSomething() error {
	return customError{ 
		code: 500,
		message: "Something went wrong!",
	}
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Operation Completed Successfully")
}