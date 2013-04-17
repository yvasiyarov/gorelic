package gorelic

import (
    "fmt"
    "encoding/json"
)

//redirect_host = send_request(None, url, 'get_redirect_host',
//                license_key)

var LicenseKey string
var CollectorUrl string
const AGENT_VERSION = "0.1"
const START_COLLECTOR_URL = "http://collector.newrelic.com"
const COLLECTOR_METHOD = "/agent_listener/invoke_raw_method"

type struct Packet {
    params: map[string]string 
    headers: map[string]string 
    url: string
}

func NewPacket() *Packet {
    p := &Packet{
        params: make(map[string]string, 100)
        headers: make(map[string]string, 10)
    } 
    p.params["license_key"] = LicenseKey
    p.params["protocol_version"] = "12"
    p.params["marshal_format"] = "json"

    p.headers["User-Agent"] = fmt.Sprintf("NewRelic-GoAgent/%s", AGENT_VERSION)
    
    return p
}

func (packet *Packet) GetRedirectHostPacket(license_key string) {
    packet.params["method"] = "get_redirect_host"
    packet.url = START_COLLECTOR_URL
}


func (packet *Packet) Send() error {
    //	Header = map[string][]string{
    //		"Accept-Encoding": {"gzip, deflate"},
    //		"Accept-Language": {"en-us"},
    //		"Connection": {"keep-alive"},
    //	}
    req := &http.NewRequest("POST", packet.url + COLLECTOR_METHOD, nil)
    //TODO:  add headers
    // pass body content
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
        return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
}
