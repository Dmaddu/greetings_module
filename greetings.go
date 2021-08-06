package greetings

import "fmt"

func Hello(name string) (string, error) {

	if name == "" {
		return "", fmt.Errorf("%v", "Invalid name, name should not be empty")
	}
	message := fmt.Sprintf("Hello %s", name)
	return message, nil
}
