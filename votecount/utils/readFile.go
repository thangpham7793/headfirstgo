package utils

import (
	"errors"
	"general"
	"os"
)

//ReadTextFile returns a ponter to file read
func ReadTextFile() *os.File {
	if len(os.Args) == 1 {
		general.HandleErr(errors.New("no file path found"))
	}

	f, err := os.Open(os.Args[1])
	general.HandleErr(err)
	return f
}
