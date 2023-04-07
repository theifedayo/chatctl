package utils

import (
	"net"
)

func GetIPAddress() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, i := range interfaces {
		if (i.Flags&net.FlagUp) != 0 && (i.Flags&net.FlagLoopback) == 0 {
			addrs, err := i.Addrs()
			if err != nil {
				return ""
			}

			for _, addr := range addrs {
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					return ipNet.IP.String()
				}
			}
		}
	}

	return "unable to determine IP address"
}
