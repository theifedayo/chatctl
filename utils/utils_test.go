package utils

import (
	"net"
	"testing"
)

func fakeInterfaces() ([]net.Interface, error) {
	return nil, &net.OpError{}
}

func TestGetIPAddress(t *testing.T) {
	//originalInterfaces := net.Interfaces
	// defer func() { net.Interfaces = originalInterfaces }()
	// net.Interfaces = fakeInterfaces

	ip := GetIPAddress()

	// expected := "unable to determine IP address"
	// if ip != expected {
	// 	t.Errorf("unexpected result: got %q, expected %q", ip, expected)
	// }

	if net.ParseIP(ip) == nil {
		t.Errorf("invalid IP address: %s", ip)
	}
}
