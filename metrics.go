package main

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

type PathSize struct {
	path     string
	filePath string
}

func NewPathSize(newrelicPath, filepath string) *PathSize {
	return &PathSize{
		path:     newrelicPath,
		filePath: filepath,
	}
}

func (t *PathSize) GetUnits() string {
	return "B"
}

func (t *PathSize) GetName() string {
	return t.path
}

func (t *PathSize) GetValue() (float64, error) {

	out, err := exec.Command("du", "-c", t.filePath).Output()
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(out), "\n")

	// total line is last in the stream

	lastLine := lines[len(lines)-2]
	field := strings.Fields(lastLine)

	if len(field) != 2 || field[1] != "total" {
		return 0, errors.New("Invalid command output: " + lastLine)
	}

	val, err := strconv.ParseFloat(field[0], 64)
	if err != nil {
		return 0, err
	}

	return val, nil
}
