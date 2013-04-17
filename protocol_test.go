package gorelic

import (
    "testing"
    "fmt"
)


func TestGetRedirectHost(t *testing.T) {
    packet := NewPacketGetRedirectHost()
    if err := packet.Send(); err != nil {
        t.Errorf("Can not send get redirect host packet: %#v \n", err)
    } else {
        fmt.Printf("Request sent: %s \n", string(packet.responseBody))
    } 
}

