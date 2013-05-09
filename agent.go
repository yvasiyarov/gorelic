package gorelic

import (
	//	"encoding/json"
	"os"
	"fmt"
	//	"io/ioutil"
	//	"net/http"
	//	"net/url"
	"reflect"
)

//TODO: change to "go"
const AGENT_LANGUAGE = "python"
const AGENT_VERSION  = "0.1.0.0"
const AGENT_IDENTIFIER = "Go Agent Test"

type Agent struct {
	AppName      []string          `json:"app_name"`
	Language     string            `json:"language"`
	Settings     *AgentSettings    `json:"settings"`
	Pid          int               `json:"pid"`
	Environment  *AgentEnvironment `json:"environment"`
	Host         string            `json:"host"`
	Identifier   string            `json:"identifier"`
	AgentVersion string            `json:"agent_version"`
}

func NewAgent() *Agent {
    host, _ := os.Hostname()
	a := &Agent{
		AppName:      []string{"Python Agent Test"}, //TODO: take it from config file
		Language:     AGENT_LANGUAGE,
		Identifier:   AGENT_IDENTIFIER,
		AgentVersion: AGENT_VERSION,
		Environment:  NewAgentEnvironment(),
		Host:         host,
		Settings:     NewAgentSettings(),
	}
	return a
}

func (agent *Agent) InitAgent() error {
	packet := NewPacketGetRedirectHost()
	if err := packet.Send(); err != nil {
		log.Printf("Can not send get redirect host packet: %#v \n", err)
        return err
	} else {
		if collectorUrl, err := packet.GetResponse(); err != nil {
			log.Printf("Can not get redirect host: %#v \n", err)
            return err
		} else if collectorUrlStr, ok := collectorUrl.(string); !ok {
			return fmt.Errorf("Redirect host is not string: %#v \n", collectorUrl)
		} else {
			agent.Settings.Host = collectorUrlStr
			jsonParams := []*Agent{agent}
 
            //TODO: check Settings to use https
			packet := NewPacketConnect("http://" + collectorUrlStr, jsonParams)

			if err := packet.Send(); err != nil {
				log.Printf("Can not send connect packet: %#v \n", err)
                return err
			} else if resp, err := packet.GetResponseConfig(); err != nil {
				log.Printf("Can not get server config %#v \n", err)
                return err
			} else {
				agent.Settings.ApplyConfigFromServer(resp)
                //TODO: log messages received from remote server
			}
		}
	}
}

func (agent *Agent) SendMetricData(data *MetricData) error {
    startTime := 0
    endTime   := 0
    jsonParams := []JsonParams{agent.Settings.AgentRunId, startTime, endTime, data}        
    
    packet := NewPacketMetricData("http://" + agent.Settings.Host, jsonParams, agent.Settings.AgentRunId)
    if err := packet.Send(); err != nil {
        log.Printf("Can not send metric data: %#v \n", err)
        return err
    }
    //TODO: parse response data
}

type MetricData struct {
}

