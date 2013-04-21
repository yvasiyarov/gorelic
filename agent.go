package gorelic

import (
//	"encoding/json"
//	"errors"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"net/url"
//	"strings"
)

type AgentSettings struct {
}

type environmentAttribute []interface{}
type AgentEnvironment []environmentAttribute

func NewAgentEnvironment() *AgentEnvironment {
    //TODO:  ["Plugin List", []]

    env := &AgentEnvironment{
        environmentAttribute{"Agent Version", "1.10.2.38"},
        environmentAttribute{"Arch", "x86_64"},
        environmentAttribute{"OS", "Linux"},
        environmentAttribute{"OS version", "3.2.0-24-generic"},
        environmentAttribute{"CPU Count", "1"},
        environmentAttribute{"System Memory", "2003.6328125"},
        environmentAttribute{"Python Program Name", "/usr/local/bin/newrelic-admin"},
        environmentAttribute{"Python Executable", "/usr/bin/python"},
        environmentAttribute{"Python Home", ""},
        environmentAttribute{"Python Path", ""},
        environmentAttribute{"Python Prefix", "/usr"},
        environmentAttribute{"Python Exec Prefix", "/usr"},
        environmentAttribute{"Python Version", "2.7.3 (default, Apr 20 2012, 22:39:59) \n[GCC 4.6.3]"},
        environmentAttribute{"Python Platform", "linux2"},
        environmentAttribute{"Python Max Unicode", "1114111"},
        environmentAttribute{"Compiled Extensions", ""},
    }
    return env
}


type Agent struct {
    AppName []string `json:"app_name"`
    Language string `json:"language"`
    Settings *AgentSettings `json:"settings"`
    Pid      int `json:"pid"`
    Environment *AgentEnvironment `json:"environment"`
    Host  string `json:"host"`
    Identifier  string `json:"identifier"`
    AgentVersion string `json:"agent_version"`
}

func NewAgent() *Agent {
    a := &Agent{
        AppName: []string{"Python Agent Test"},
        Language: "python",
        Identifier: "Python Agent Test",
        AgentVersion: "1.10.2.38",
        Environment: NewAgentEnvironment(),
    }
    return a
}

type Settings struct {
}


