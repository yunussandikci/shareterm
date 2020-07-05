package common

import (
	"encoding/hex"
	"github.com/sethvargo/go-password/password"
	"os"
)

const DefaultHost = "http://127.0.0.1:9000"

func EncodeHex(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

func DecodeHex(text string) []byte {
	decodeText, decodeErr := hex.DecodeString(text)
	if decodeErr != nil {
		panic(decodeErr)
	}
	return decodeText
}

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