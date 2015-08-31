package main

import (
	"errors"
	"regexp"
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

func BytesToFloat(data []byte) (float64, error) {
	return strconv.ParseFloat(strings.TrimSpace(string(data)), 64)
}

type RssPerCommand struct {
	processor *RssProcessor
	path      string
	filter    *regexp.Regexp
	samples   map[string]ISampleStats
}

func NewRssPerCommand(processor *RssProcessor, path string, commandFilter *regexp.Regexp) *RssPerCommand {
	return &RssPerCommand{
		processor: processor,
		path:      path,
		filter:    commandFilter,
	}
}

func (r *RssPerCommand) GetName(id string) string {
	return r.path + "/" + id
}

func (r *RssPerCommand) GetUnits() string {
	return "KB"
}

func (r *RssPerCommand) GetValue(id string) (float64, error) {
	return r.samples[id].GetAverage(), nil
}

func (r *RssPerCommand) GetIdList() []string {
	r.samples = r.processor.GetAndPurgeRss()
	var list []string

	for _, v := range r.processor.GetUniqKeys(r.samples) {
		if r.filter != nil && !r.filter.MatchString(v) {
			continue
		}

		list = append(list, v)
	}

	return list
}
