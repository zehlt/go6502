package go6502

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("hi!, %v. Welcome!", name)
	return message
}
