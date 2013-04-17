package gorelic

import (
	"encoding/json"
	"fmt"
)

//redirect_host = send_request(None, url, 'get_redirect_host',
//                license_key)

var LicenseKey string
var CollectorUrl string

const AGENT_VERSION = "0.1"
const START_COLLECTOR_URL = "http://collector.newrelic.com"
const COLLECTOR_METHOD = "/agent_listener/invoke_raw_method"

type Packet struct {
	params  map[string]string
	headers map[string]string
	url     string
}

func NewPacket() *Packet {
	p := &Packet{
		params:  map[string]string{},
		headers: map[string]string{},
	}
	p.params["license_key"] = LicenseKey
	p.params["protocol_version"] = "12"
	p.params["marshal_format"] = "json"

	p.headers["User-Agent"] = fmt.Sprintf("NewRelic-GoAgent/%s", AGENT_VERSION)
	p.headers["Accept-Encoding"] = "identity, deflate, compress, gzip"
	p.headers["Accept"] = "*/*"
	p.headers["Content-Encoding"] = "identity"

	return p
}

func NewPacketGetRedirectHost() *Packet {
	packet := NewPacket()
	packet.params["method"] = "get_redirect_host"
	packet.url = START_COLLECTOR_URL
}

func (packet *Packet) Send() error {
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Connection": {"keep-alive"},
	//	}
	req := &http.NewRequest("POST", packet.url+COLLECTOR_METHOD, nil)
	//TODO:  add headers
	// pass body content
	// use deflate as Content-Encoding: if len(data) > 64*1024
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
}
