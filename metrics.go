package main

import (
	"errors"
	"strconv"
	"strings"
)

func MakeSizeCommand(path string) string {
	return "du -c " + path + " | grep total | awk '{print $1}'"
}

func SizeCommandToFloat(data []byte) (float64, error) {

	fields := strings.Fields(string(data))

	if len(fields) == 0 {
		return 0, errors.New("Invalid command output: " + string(data))
	}

	val, err := strconv.ParseFloat(fields[len(fields)-1], 64)
	if err != nil {
		return 0, err
	}

	return val, nil
}
