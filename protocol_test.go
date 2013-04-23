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
		t.Errorf("Can not send get redirect host packet: %#v \n", err)
	} else {
		resp, err := packet.GetResponse()
		fmt.Printf("Got response: %#v \nerror: %v\n", resp, err)
	}
}
