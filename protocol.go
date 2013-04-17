package gorelic

import (
//	"encoding/json"
	"fmt"
    "net/http"
    "net/url"
    "strings"
    "io/ioutil"
)

//redirect_host = send_request(None, url, 'get_redirect_host',
//                license_key)

var LicenseKey string
var CollectorUrl string

const AGENT_VERSION = "0.1"
const START_COLLECTOR_URL = "http://collector.newrelic.com"
const COLLECTOR_METHOD = "/agent_listener/invoke_raw_method"

type Packet struct {
	params  url.Values
	header  http.Header
	url     string
    responseBody []byte
}

func NewPacket() *Packet {
	p := &Packet{
		params:  url.Values{},
		header:  http.Header{},
	}
	p.params.Add("license_key", LicenseKey)
	p.params.Add("protocol_version", "12")
	p.params.Add("marshal_format", "json")

	p.header["User-Agent"] = []string{"NewRelic-PythonAgent/1.10.2.38 (Python 2.7.3 linux2)"}
	//p.header["User-Agent"] = []string{fmt.Sprintf("NewRelic-GoAgent/%s", AGENT_VERSION)}
	p.header["Accept-Encoding"] = []string{"identity, deflate, compress, gzip"}
	p.header["Accept"] = []string{"*/*"}
	p.header["Content-Encoding"] = []string{"identity"}

	return p
}

func NewPacketGetRedirectHost() *Packet {

	packet := NewPacket()
	packet.params.Add("method", "get_redirect_host")
	packet.url = START_COLLECTOR_URL

    return packet
}

func (packet *Packet) Send() error {

    bodyContent := packet.params.Encode()
    //if len(bodyContent) > 64*1024 {
	//    packet.header["Content-Encoding"] = []string{"deflate"}
    //}

	body := ioutil.NopCloser(strings.NewReader(""))
	//body := ioutil.NopCloser(strings.NewReader(bodyContent))
    //fmt.Println(packet.url + COLLECTOR_METHOD + "?" + bodyContent)

    if req, err := http.NewRequest("POST", packet.url + COLLECTOR_METHOD + "?" + bodyContent, body); err != nil {
        return err
    } else {
	    //packet.header["Content-Encoding"] = []string{"identity"}
        req.Header = packet.header
        req.ContentLength = 0
        fmt.Printf("req:  %s", req.URL.String()) 
        if resp, err := http.DefaultClient.Do(req); err != nil {
            return err
        } else {
	        defer resp.Body.Close()
	        if b, err := ioutil.ReadAll(resp.Body); err != nil {
                return err
            } else {
                packet.responseBody = b
            }
        }
    }
    return nil
}
