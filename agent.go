package gorelic

import (
	"fmt"
	"github.com/yvasiyarov/newrelic_platform_go"
	"log"
)

const (
	// Send data to newrelic every 60 seconds
	NEWRELIC_POLL_INTERVAL = 60
	// Get garbage collector run statistic every 10 seconds
	// During GC stat pooling - mheap will be locked, so be carefull changing this value
	GC_POLL_INTERVAL_IN_SECONDS = 10
	// Get memory allocator statistic every 60 seconds
	// During this process stoptheword() is called, so be carefull changing this value
	MEMORY_ALLOCATOR_POLL_INTERVAL_IN_SECONDS = 60

	AGENT_GUID    = "com.github.yvasiyarov.GoRelic"
	AGENT_VERSION = "0.0.1"
	AGENT_NAME    = "Go daemon"
)

type Agent struct {
	NewrelicName                string
	NewrelicLicense             string
	NewrelicPollInterval        int
	Verbose                     bool
	CollectGcStat               bool
	CollectMemoryStat           bool
	GCPollInterval              int
	MemoryAllocatorPollInterval int
	AgentGUID                   string
	AgentVersion                string
	plugin                      *newrelic_platform_go.NewrelicPlugin
}

func NewAgent() *Agent {
	agent := &Agent{
		NewrelicName:                AGENT_NAME,
		NewrelicPollInterval:        NEWRELIC_POLL_INTERVAL,
		Verbose:                     false,
		CollectGcStat:               true,
		CollectMemoryStat:           true,
		GCPollInterval:              GC_POLL_INTERVAL_IN_SECONDS,
		MemoryAllocatorPollInterval: MEMORY_ALLOCATOR_POLL_INTERVAL_IN_SECONDS,
		AgentGUID:                   AGENT_GUID,
		AgentVersion:                AGENT_VERSION,
	}
	return agent
}

func (agent *Agent) Run() {
	if agent.NewrelicLicense == "" {
		log.Fatalf("Please, pass a valid newrelic license key.")
	}

	agent.plugin = newrelic_platform_go.NewNewrelicPlugin(agent.AgentVersion, agent.NewrelicLicense, agent.NewrelicPollInterval)
	component := newrelic_platform_go.NewPluginComponent(agent.NewrelicName, agent.AgentGUID)
	agent.plugin.AddComponent(component)

	component.AddMetrica(&NOGoroutinesMetrica{})
	component.AddMetrica(&NOCgoCallsMetrica{})
	if agent.CollectGcStat {
		addGCMericsToComponent(component, agent.GCPollInterval)
		agent.Debug(fmt.Sprintf("Init GC metrics collection. Poll interval %d seconds.", agent.GCPollInterval))
	}
	if agent.CollectMemoryStat {
		addMemoryMericsToComponent(component, agent.MemoryAllocatorPollInterval)
		agent.Debug(fmt.Sprintf("Init memory allocator metrics collection. Poll interval %d seconds.", agent.MemoryAllocatorPollInterval))
	}

	agent.plugin.Verbose = agent.Verbose
	go agent.plugin.Run()
}

func (agent *Agent) Debug(msg string) {
	if agent.Verbose {
		log.Println(msg)
	}
}
