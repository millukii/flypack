package user

import (
	"fmt"

	"github.com/sethvargo/go-password/password"
)

type PasswordGenerator struct{}

func (p PasswordGenerator) GeneratePassword() (string, error) {
	res, err := password.Generate(64, 10, 0, false, true)
	if err != nil {
		fmt.Println("Error generating password ", err)
		return "", err
	}
	return res, nil
}