package gorelic

import (
	"encoding/json"
	"fmt"
	"testing"
)

const SETTINGS_FROM_SERVER_JSON = `
{"agent_run_id":463114387,
"product_level":40,
"episodes_url":"https:\/\/d1ros97qkrwjf5.cloudfront.net\/42\/eum\/rum.js",
"cross_process_id":"142416#1783720",
"collect_errors":true,
"url_rules":
[
    {"ignore":false,
    "replacement":"\/*.\\1",
    "replace_all":false,
    "each_segment":false,
    "terminate_chain":true,
    "eval_order":1000,
    "match_expression":".*\\.(ace|arj|ini|txt|udl|plist|css|gif|ico|jpe?g|js|png|swf|woff|caf|aiff|m4v|mpe?g|mp3|mp4|mov)$"},
    
    {"ignore":false,
    "replacement":"*",
    "replace_all":false,
    "each_segment":true,
    "terminate_chain":false,
    "eval_order":1001,
    "match_expression":"^[0-9][0-9a-f_,.-]*$"},

   {"ignore":false,
   "replacement":"\\1\/.*\\2",
   "replace_all":false,
   "each_segment":false,
   "terminate_chain":false,
   "eval_order":1002,
   "match_expression":"^(.*)\/[0-9][0-9a-f_,-]*\\.([0-9a-z][0-9a-z]*)$"}
],
"messages":[
    {"message":"Reporting to: https:\/\/rpm.newrelic.com\/accounts\/142416\/applications\/1783720",
    "level":"INFO"}
],
"data_report_period":60,
"collect_traces":true,
"sampling_rate":0,
"browser_key":"04ff564b25",
"encoding_key":"d67afc830dab717fd163bfcb0b8b88423e9a1a3b",
"apdex_t":0.5,
"episodes_file":"d1ros97qkrwjf5.cloudfront.net\/42\/eum\/rum.js",
"trusted_account_ids":[142416]
,"beacon":"beacon-1.newrelic.com",
"application_id":"1783720"}
`

func TestApplyConfigFromServer(t *testing.T) {
	settings := NewAgentSettings()
	var result AgentSettings
	err := json.Unmarshal([]byte(SETTINGS_FROM_SERVER_JSON), &result)
	if err != nil {
		t.Errorf("Can not decode example config: %#v \n", err)
		if unmarshalErr, ok := err.(*json.UnmarshalTypeError); ok {
			t.Errorf("Can not decode %#v to %#v\n", unmarshalErr.Value, unmarshalErr.Type.Kind().String())
		}
	} else {
		settings.ApplyConfigFromServer(&result)
		//fmt.Printf("Result config: %#v \n", settings)  
	}
}

func TestInitAgent(t *testing.T) {
	agent := NewAgent()
	packet := NewPacketGetRedirectHost()
	if err := packet.Send(); err != nil {
		t.Errorf("Can not send get redirect host packet: %#v \n", err)
	} else {
		if collectorUrl, err := packet.GetResponse(); err != nil {
			t.Errorf("Can not get redirect host: %#v \n", err)
		} else if collectorUrlStr, ok := collectorUrl.(string); !ok {
			t.Errorf("Redirect host is not string: %#v \n", collectorUrl)
		} else {
			agent.Settings.Host = collectorUrlStr
			jsonParams := []*Agent{agent}
			packet := NewPacketConnect("http://" + collectorUrlStr, jsonParams)
			if err := packet.Send(); err != nil {
				t.Errorf("Can not send connect packet: %#v \n", err)
			} else if resp, err := packet.GetResponseConfig(); err != nil {
				t.Errorf("Can not get server config %#v \n", err)
			} else {
				fmt.Printf("Got config: %#v \n", resp)
				agent.Settings.ApplyConfigFromServer(resp)
			}
		}
	}
}
