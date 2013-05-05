package gorelic

import (
	"fmt"
	"testing"
)

func TestGetRedirectHost(t *testing.T) {
	packet := NewPacketGetRedirectHost()
	if err := packet.Send(); err != nil {
		t.Errorf("Can not send get redirect host packet: %#v \n", err)
	} else {
		resp, err := packet.GetResponse()
		fmt.Printf("Got response: %#v \nerror: %v\n", resp, err)
	}
}

func TestConnect(t *testing.T) {
	agent := NewAgent()

	jsonParams := []*Agent{agent}
	packet := NewPacketConnect("http://collector-6.newrelic.com", jsonParams)
	if err := packet.Send(); err != nil {
		t.Errorf("Can not send connect packet: %#v \n", err)
	} else {
		resp, err := packet.GetResponse()
		fmt.Printf("Got response: %#v \nerror: %v\n", resp, err)
	}
}

const CONNECT_PACKET_RESPONSE_JSON = `
{"return_value":
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
}
`

func TestGetResponseTyped(t *testing.T) {
	agent := NewAgent()

	jsonParams := []*Agent{agent}
	packet := NewPacketConnect("http://collector-6.newrelic.com", jsonParams)
	packet.responseCode = 200
	packet.rawResponse = []byte(CONNECT_PACKET_RESPONSE_JSON)

	settings := NewAgentSettings()
	if _, err := packet.GetResponseTyped(settings); err != nil {
		t.Errorf("Can not get settings typed: %#v \n", err)
		//	} else {
		//		fmt.Printf("Got settings %#v \n\n original %#v \n", result, settings)
	}
}

func TestGetResponseConfig(t *testing.T) {

	packet := NewPacketConnect("http://collector-6.newrelic.com", nil)
	packet.responseCode = 200
	packet.rawResponse = []byte(CONNECT_PACKET_RESPONSE_JSON)

	if result, err := packet.GetResponseConfig(); err != nil {
		t.Errorf("Can not get response settings : %#v, %#v \n", err, result)
		//	} else {
		//		fmt.Printf("Got response settings %#v \n", result)
	}
}
