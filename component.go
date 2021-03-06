package main

import (
	ext "github.com/maciejmrowiec/pgmonitor/newrelic"
	platform "github.com/yvasiyarov/newrelic_platform_go"
	"regexp"
)

func InitSingleSamplesComponent(hostname string, verbose bool) platform.IComponent {

	component := platform.NewPluginComponent(hostname, "com.github.maciejmrowiec.cfclientmonitor", verbose)

	component.AddMetrica(NewCommandMetric("disksize/cfengine", MakeSizeCommand("/var/cfengine"), "B", SizeCommandToFloat))

	component.AddMetrica(NewCommandMetric("cpu/loadaverage/1m", "cat /proc/loadavg | awk '{print$1}'", "", BytesToFloat))
	component.AddMetrica(NewCommandMetric("cpu/loadaverage/5m", "cat /proc/loadavg | awk '{print$2}'", "", BytesToFloat))
	component.AddMetrica(NewCommandMetric("cpu/loadaverage/15m", "cat /proc/loadavg | awk '{print$3}'", "", BytesToFloat))

	component.AddMetrica(NewCommandMetric("cpu/tasks/active", "cat /proc/loadavg | awk '{print$4}' | awk -F '/' '{print $1}'", "units", BytesToFloat))
	component.AddMetrica(NewCommandMetric("cpu/tasks/total", "cat /proc/loadavg | awk '{print$4}' | awk -F '/' '{print $2}'", "units", BytesToFloat))

	return component

}

func InitDynamicSamplesComponent(hostname string, verbose bool) platform.IComponent {

	rss := NewRssProcessor()
	NewDynamicCollector(rss).Run()

	component := ext.NewDynamicPluginComponent(hostname, "com.github.maciejmrowiec.cfclientmonitor", verbose)
	component.AddDynamicMetrica(NewRssPerCommand(rss, "memory/rss/average", regexp.MustCompile("/var/cfengine/bin/.*")))
	component.AddDynamicMetrica(NewMaxRssPerCommand(rss, "memory/rss/max", regexp.MustCompile("/var/cfengine/bin/.*")))

	return component
}
