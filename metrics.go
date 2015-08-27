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

	out, err := exec.Command("/bin/bash", "-c", "du -c "+t.filePath+" | grep total | awk '{print $1}'").Output()
	if err != nil {
		return 0, err
	}

	fields := strings.Fields(string(out))

	if len(fields) == 0 {
		return 0, errors.New("Invalid command output: " + string(out))
	}

	val, err := strconv.ParseFloat(fields[len(fields)-1], 64)
	if err != nil {
		return 0, err
	}

	return val, nil
}
