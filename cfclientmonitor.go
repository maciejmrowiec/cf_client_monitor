package main

import (
	platform "github.com/yvasiyarov/newrelic_platform_go"
	"log"
	"os"
)

func main() {
	config := HandleUserOptions()

	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	plugin := platform.NewNewrelicPlugin("0.0.1", config.newRelicKey, config.interval*60)

	plugin.AddComponent(InitSingleSamplesComponent(hostName, config.verbose))

	plugin.Verbose = config.verbose
	plugin.Run()
}
