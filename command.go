package main

import (
	"os/exec"
)

type CommandMetric struct {
	units        string
	shellCommand string
	path         string
	convertFunc  func([]byte) (float64, error)
}

func NewCommandMetric(newrelicPath, command, units string, convertFunc func([]byte) (float64, error)) *CommandMetric {
	return &CommandMetric{
		units:        units,
		shellCommand: command,
		path:         newrelicPath,
		convertFunc:  convertFunc,
	}
}

func (t *CommandMetric) GetUnits() string {
	return t.units
}

func (t *CommandMetric) GetName() string {
	return t.path
}

func (t *CommandMetric) GetValue() (float64, error) {

	out, err := exec.Command("/bin/bash", "-c", t.shellCommand).Output()
	if err != nil {
		return 0, err
	}

	return t.convertFunc(out)
}
