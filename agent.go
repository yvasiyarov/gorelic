package main

import (
	"flag"
	"github.com/yvasiyarov/newrelic_platform_go"
	"log"
	"math/rand"
	"time"
)

var newrelicName = flag.String("newrelic-name", "Go daemon", "Component name in New Relic")
var newrelicLicense = flag.String("newrelic-license", "", "Newrelic license")

var verbose = flag.Bool("verbose", false, "Verbose mode")

const (
	NEWRELIC_POLL_INTERVAL = 60 //Send data to newrelic every 60 seconds

	AGENT_GUID                  = "com.github.yvasiyarov.GoRelic"
	AGENT_VERSION               = "0.0.1"
	GC_POLL_INTERVAL_IN_SECONDS = 10
)

func allocateAndSum(arraySize int) int {
	arr := make([]int, arraySize, arraySize)
	for i, _ := range arr {
		arr[i] = rand.Int()
	}
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	result := 0
	for _, v := range arr {
		result += v
	}
	log.Printf("Array size is: %d, sum is: %d\n", arraySize, result)
	return result
}

func doSomeJob(numRoutines int) {
	for i := 0; i < numRoutines; i++ {
		go allocateAndSum(rand.Intn(1024) * 1024)
	}
}

func main() {

	flag.Parse()
	if *newrelicLicense == "" {
		log.Fatalf("Please, pass a valid newrelic license key.\n Use --help to get more information about available options\n")
	}

	plugin := newrelic_platform_go.NewNewrelicPlugin(AGENT_VERSION, *newrelicLicense, NEWRELIC_POLL_INTERVAL)
	component := newrelic_platform_go.NewPluginComponent(*newrelicName, AGENT_GUID)
	plugin.AddComponent(component)

	component.AddMetrica(&NOGoroutinesMetrica{})
	component.AddMetrica(&NOCgoCallsMetrica{})
	addGCMericsToComponent(component)

	plugin.Verbose = *verbose
	plugin.Run()
	/*
		doSomeJob(100)
		log.Println("All routines started\n")
		select {}
	*/
}
