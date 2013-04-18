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
