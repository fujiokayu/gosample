package gosample

import "fmt"

func Print(value interface{}) {
	s, ok := value.(string) // Type Assertion
	if ok {
		fmt.Printf("value is string: %s\n", s)
	} else {
		fmt.Printf("value is not string\n")
	}
}
