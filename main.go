package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("el nombre está vacío")
	}
	greeting := fmt.Sprintf("Hola %v, bienvenido!", name)
	return greeting, nil
}
