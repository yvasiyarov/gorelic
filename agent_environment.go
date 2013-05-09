package gorelic

import (
    "runtime"
    "strconv"
)

type environmentAttribute []interface{}
type AgentEnvironment []environmentAttribute

func NewAgentEnvironment() *AgentEnvironment {
    //TODO: detect real environment settings
   
    numCpu := strconv.Itoa(runtime.NumCPU())
    
	env := &AgentEnvironment{
		environmentAttribute{"Agent Version", AGENT_VERSION},
		environmentAttribute{"Arch", "x86_64"},
		environmentAttribute{"OS", "Linux"},
		environmentAttribute{"OS version", "3.2.0-24-generic"},
		environmentAttribute{"CPU Count", numCpu},
		environmentAttribute{"System Memory", "2003.6328125"},
		environmentAttribute{"Go Version", runtime.Version()},
	}
	return env
}

