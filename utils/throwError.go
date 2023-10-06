package utils

import (
	"fmt"
	"os"
)

const (
	BasicError   = "Erro:"
	PathNotFound = "File not found:"
	RemoveFile   = "Error removing file:"
)

func ThrowErrorIfHas(err error, message string, code int) {
	if err != nil {
		fmt.Println()
		fmt.Println()
		fmt.Println("     ⚠️ ", message, err.Error())
		fmt.Println()
		fmt.Println()
		os.Exit(code)
	}
}
