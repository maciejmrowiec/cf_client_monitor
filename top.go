package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func init() {
	RowRegex = regexp.MustCompile("[\\s]+")
}

var RowRegex *regexp.Regexp

type CommandRss struct {
	rss    float64
	rssMax float64
	name   string
}

func (p *CommandRss) GetName() string {
	return p.name
}

func (p *CommandRss) Aggregate(item IItem) {
	process := item.(*CommandRss)
	p.rss += process.rss
	p.rssMax += process.rssMax
}

type RssProcessor struct {
	rss               map[string]ISampleStats
	rssMax            map[string]ISampleStats
	measurements_lock sync.Mutex
}

func NewRssProcessor() *RssProcessor {
	return &RssProcessor{
		rss:    make(map[string]ISampleStats, 200),
		rssMax: make(map[string]ISampleStats, 200),
	}
}

func (i *RssProcessor) GetAndPurgeRss() map[string]ISampleStats {
	i.measurements_lock.Lock()
	samples := i.rss
	i.rss = make(map[string]ISampleStats, 200)
	i.measurements_lock.Unlock()

	return samples
}

func (i *RssProcessor) GetAndPurgeRssMax() map[string]ISampleStats {
	i.measurements_lock.Lock()
	samples := i.rssMax
	i.rssMax = make(map[string]ISampleStats, 200)
	i.measurements_lock.Unlock()

	return samples
}

func (i *RssProcessor) GetUniqKeys(data map[string]ISampleStats) []string {
	var sample_list []string

	for key := range data {
		sample_list = append(sample_list, key)
	}

	return sample_list
}

func (p *RssProcessor) GetCmdName() string {
	return "bash"
}

func (p *RssProcessor) GetCmdArgs() []string {
	return []string{"-c", "ps aux | awk '{print $6,$11}'"}
}

func (i *RssProcessor) AggregateSample(sample ISample) {
	i.measurements_lock.Lock()

	for key, val := range sample.GetMap() {
		pio := val.(*CommandRss)

		if entry, has := i.rss[key]; has {
			entry.Append(pio.rss)
		} else {
			i.rss[key] = NewStatSample(pio.rss)
		}
		if entry, has := i.rssMax[key]; has {
			entry.Append(pio.rssMax)
		} else {
			i.rssMax[key] = NewStatSample(pio.rssMax)
		}
	}
	i.measurements_lock.Unlock()
}

func (p *RssProcessor) NewSample() ISample {
	return NewSample()
}

func (p *RssProcessor) IsHeader(row string) bool {
	return strings.Contains(row, "RSS")
}

func (p *RssProcessor) ParseRow(row string) (IItem, error) {
	tokens := RowRegex.Split(row, 2)
	if len(tokens) != 2 {
		return nil, errors.New("Failed to parse row data for sample")
	}

	var err error
	commandRss := new(CommandRss)

	if commandRss.rss, err = strconv.ParseFloat(tokens[0], 64); err != nil {
		return nil, err
	}

	commandRss.rssMax = commandRss.rss
	commandRss.name = tokens[1]

	return commandRss, nil
}
