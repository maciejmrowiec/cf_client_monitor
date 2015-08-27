package main

import (
	// ext "github.com/maciejmrowiec/pgmonitor/newrelic"
	platform "github.com/yvasiyarov/newrelic_platform_go"
)

func InitSingleSamplesComponent(hostname string, verbose bool) platform.IComponent {

	component := platform.NewPluginComponent(hostname, "com.github.maciejmrowiec.cfclientmonitor", verbose)

	component.AddMetrica(NewPathSize("disksize/cfengine", "/var/cfengine/"))

	return component

}
