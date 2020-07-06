package common

import (
	"github.com/sethvargo/go-password/password"
	"os"
)

const DefaultHost = "https://shareterm.tech"

func GetHost() string {
	CustomHost := os.Getenv("SHARETERM_HOST")
	if CustomHost != "" {
		return CustomHost
	}
	return DefaultHost
}

func GenerateAESKey() string {
	key, generateErr := password.Generate(32, 8, 0, true, false)
	if generateErr != nil {
		panic(generateErr.Error())
	}
	return key
}