package toolkit

import (
	"net"
	"net/http"
	"strings"
)

const defaultHost = "127.0.0.1"

// ClientIP returns client's IP.
// If IP can not be detected function returns 127.0.0.1
func ClientIP(req *http.Request) string {
	ips := req.Header.Get("X-Real-IP")
	if ips == "" {
		ips = req.Header.Get("X-Forwarded-For")
		if ips == "" {
			ips = req.RemoteAddr
		}
	}

	ipList := strings.Split(ips, ", ")

	ip := net.ParseIP(ipList[len(ipList)-1])
	if ip == nil {
		host, _, err := net.SplitHostPort(ipList[len(ipList)-1])
		if err != nil {
			return defaultHost
		}

		return host
	}

	return ip.String()
}
