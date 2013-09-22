package main

import (
	"flag"
	"github.com/yvasiyarov/newrelic_platform_go"
	"log"
)

var newrelicName = flag.String("newrelic-name", "Go daemon", "Component name in New Relic")
var newrelicLicense = flag.String("newrelic-license", "", "Newrelic license")

var verbose = flag.Bool("verbose", false, "Verbose mode")

const (
	NEWRELIC_POLL_INTERVAL = 60 //Send data to newrelic every 60 seconds

	AGENT_GUID    = "com.github.yvasiyarov.GoRelic"
	AGENT_VERSION = "0.0.1"
)

func main() {
	flag.Parse()
	if *newrelicLicense == "" {
		log.Fatalf("Please, pass a valid newrelic license key.\n Use --help to get more information about available options\n")
	}

	plugin := newrelic_platform_go.NewNewrelicPlugin(AGENT_VERSION, *newrelicLicense, NEWRELIC_POLL_INTERVAL)
	component := newrelic_platform_go.NewPluginComponent(*newrelicName, AGENT_GUID)
	plugin.AddComponent(component)

	plugin.Verbose = *verbose
	plugin.Run()
}
