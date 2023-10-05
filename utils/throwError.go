package utils

import "log"

func throwExit(err error, message string) {
	if err != nil {
		log.Fatal(err, message)
	}
}
