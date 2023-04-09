package error

import "fmt"

// Consume prints the error and returns nothing
func Consume(err error, print string) {
	if err != nil {
		fmt.Println(print, err)
	}
}

// Process prints the error and returns a boolean
func Process(err error, print string) bool {
	if err != nil {
		fmt.Println(print, err)
		return false
	}
	return true
}
