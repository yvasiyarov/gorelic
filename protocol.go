package gorelic

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
    "strconv"
)

var LicenseKey string
var CollectorUrl string

const AGENT_VERSION = "0.1"
const START_COLLECTOR_URL = "http://collector.newrelic.com"
const COLLECTOR_METHOD = "/agent_listener/invoke_raw_method"

type JsonParams interface{}
type Packet struct {
	params       url.Values
	header       http.Header
	url          string
	rawResponse  []byte
	responseCode int
	jsonParams   JsonParams
}

type ResponseJson map[string]interface{}
type ResponseConfig map[string]AgentSettings

func NewPacket() *Packet {
	p := &Packet{
		params: url.Values{},
		header: http.Header{},
	}
	p.params.Add("license_key", "7bceac019c7dcafae1ef95be3e3a3ff8866de246")
	//p.params.Add("license_key", LicenseKey)
	p.params.Add("protocol_version", "12")
	p.params.Add("marshal_format", "json")

	p.header["User-Agent"] = []string{fmt.Sprintf("NewRelic-GoAgent/%s", AGENT_VERSION)}
	p.header["Accept-Encoding"] = []string{"identity, deflate, compress, gzip"}
	p.header["Accept"] = []string{"*/*"}
	p.header["Content-Encoding"] = []string{"identity"}

	return p
}

func NewPacketGetRedirectHost() *Packet {

	packet := NewPacket()
	packet.params.Add("method", "get_redirect_host")
	packet.url = START_COLLECTOR_URL
	packet.jsonParams = map[string]string{}

	return packet
}

func NewPacketConnect(url string, jsonParams JsonParams) *Packet {

	packet := NewPacket()
	packet.params.Add("method", "connect")
	packet.url = url
	packet.jsonParams = jsonParams
	return packet
}

func NewPacketMetricData(url string, jsonParams JsonParams, runId int) *Packet {

	packet := NewPacket()
	packet.params.Add("method", "metric_data")
	packet.params.Add("run_id", strconv.Itoa(runId))
	packet.url = url
	packet.jsonParams = jsonParams
	return packet
}

func (packet *Packet) Send() error {
	bodyContent, err := json.Marshal(packet.jsonParams)
	if err != nil {
		return err
	}
	//if len(bodyContent) > 64*1024 {
	//    packet.header["Content-Encoding"] = []string{"deflate"}
	//}

	body := ioutil.NopCloser(strings.NewReader(string(bodyContent)))
	fmt.Println(string(bodyContent))

	urlParams := packet.params.Encode()
	if req, err := http.NewRequest("POST", packet.url+COLLECTOR_METHOD+"?"+urlParams, body); err != nil {
		return err
	} else {
		req.Header = packet.header
		req.ContentLength = int64(len(bodyContent))
		if resp, err := http.DefaultClient.Do(req); err != nil {
			return err
		} else {
			defer resp.Body.Close()
			if responseBody, err := ioutil.ReadAll(resp.Body); err != nil {
				return err
			} else {
				packet.responseCode = resp.StatusCode
				packet.rawResponse = responseBody
			}
		}
	}
	return nil
}

func (packet *Packet) GetResponse() (interface{}, error) {
	if packet.responseCode != 200 {
		return nil, errors.New(string(packet.rawResponse))
	}
	var result ResponseJson
	err := json.Unmarshal(packet.rawResponse, &result)
	if err != nil {
		return nil, err
	}
	if returnValue, ok := result["return_value"]; ok {
		return returnValue, nil
	} else {
		err = errors.New("Can not get return value")
	}
	return nil, err
}
func (packet *Packet) GetResponseConfig() (*AgentSettings, error) {
	if packet.responseCode != 200 {
		return nil, errors.New(string(packet.rawResponse))
	}
	var result ResponseConfig

	err := json.Unmarshal(packet.rawResponse, &result)
	if err != nil {
		return nil, err
	}
	if returnValue, ok := result["return_value"]; ok {
		return &returnValue, nil
	} else {
		err = errors.New("Can not get return value")
	}
	return nil, err
}

//This method make one extra memory allocation, it should not be used
func (packet *Packet) GetResponseTyped(resultValue interface{}) (interface{}, error) {
	if packet.responseCode != 200 {
		return nil, errors.New(string(packet.rawResponse))
	}

	result := make(ResponseJson, 1)
	result["return_value"] = resultValue

	err := json.Unmarshal(packet.rawResponse, &result)
	if err != nil {
		return nil, err
	}
	if returnValue, ok := result["return_value"]; ok {
		return returnValue, nil
	} else {
		err = errors.New("Can not get return value")
	}
	return nil, err
}
