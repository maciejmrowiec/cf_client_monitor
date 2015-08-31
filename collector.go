package main

import (
	"bufio"
	"log"
	"os/exec"
	"time"
)

type IItem interface {
	GetName() string
	Aggregate(item IItem)
}

type ISample interface {
	Append(item IItem)
	Empty() bool
	GetMap() map[string]IItem
}

type ISampleProcessor interface {
	GetCmdName() string
	GetCmdArgs() []string
	AggregateSample(sample ISample)
	NewSample() ISample
	IsHeader(row string) bool
	ParseRow(row string) (IItem, error)
}

type ISampleStats interface {
	Append(value float64)
	GetAverage() float64
}

type DynamicCollector struct {
	processor ISampleProcessor
}

func NewDynamicCollector(processor ISampleProcessor) *DynamicCollector {
	return &DynamicCollector{
		processor: processor,
	}
}

func (d *DynamicCollector) Run() {
	pipe := make(chan string, 1000)
	go d.executeCmd(d.processor.GetCmdName(), d.processor.GetCmdArgs(), pipe)
	go d.processOutput(pipe)
}

func (d *DynamicCollector) processOutput(ch <-chan string) {
	sample := d.processor.NewSample()

	for row := range ch {

		if d.processor.IsHeader(row) {
			if sample != nil && !sample.Empty() {
				d.processor.AggregateSample(sample)
			}

			sample = d.processor.NewSample()
			continue
		}

		p, err := d.processor.ParseRow(row)
		if err != nil {
			continue
		}

		sample.Append(p)
	}
}

func (d *DynamicCollector) executeCmd(name string, args []string, ch chan<- string) {

	execute_cmd := func() {
		cmd := exec.Command(name, args...)
		stdout, err := cmd.StdoutPipe()

		if err != nil {
			log.Fatal(err)
		}

		if err = cmd.Start(); err != nil {
			log.Fatal(err)
		}
		defer cmd.Wait()

		in := bufio.NewScanner(stdout)

		for in.Scan() {
			ch <- in.Text()
		}
	}

	for true {
		execute_cmd()
		time.Sleep(time.Second)
	}

}

type StatSample struct {
	total float64
	count float64
}

func NewStatSample(value float64) *StatSample {
	return &StatSample{
		total: value,
		count: 1,
	}
}

func (s *StatSample) Append(value float64) {
	s.count += 1
	s.total += value
}

func (s *StatSample) GetAverage() float64 {
	return s.total / s.count
}

type Sample struct {
	data map[string]IItem
}

func NewSample() *Sample {
	return &Sample{
		data: make(map[string]IItem, 10),
	}
}

func (d *Sample) Append(item IItem) {
	name := item.GetName()

	if val, has := d.data[name]; has {
		val.Aggregate(item)
	} else {
		d.data[name] = item
	}
}

func (d *Sample) Empty() bool {
	if len(d.data) > 0 {
		return false
	}
	return true
}

func (d *Sample) GetMap() map[string]IItem {
	return d.data
}
