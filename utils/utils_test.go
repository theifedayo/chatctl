package utils

import (
	"net"
	"testing"
)

//This test ensures that the GetIPAddress() function returns a valid IP address
//and does not fail to determine the IP address

func TestGetIPAddress(t *testing.T) {
	ip := GetIPAddress()

	if ip == "unable to determine IP address" {
		t.Errorf("unable to determine IP address")
	}

	if net.ParseIP(ip) == nil {
		t.Errorf("invalid IP address: %s", ip)
	}
}
