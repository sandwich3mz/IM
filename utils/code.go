package utils

import "math/rand"

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRegisterCode(length int8) string {
	code := make([]byte, length)
	for i := range code {
		code[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(code)
}
