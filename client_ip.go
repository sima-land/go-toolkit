package toolkit

import (
	"net/http"
	"strings"
)

// ClientIP returns client's IP.
// If IP can not be detected function returns 127.0.0.1
func ClientIP(req *http.Request) string {
	ip := req.Header.Get("X-Real-IP")
	if ip == "" {
		ip = req.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = req.RemoteAddr
		}
	}
	if comma := strings.LastIndex(ip, ","); comma != -1 {
		ip = strings.Trim(ip[comma+1:], " ")
	}
	if colon := strings.LastIndex(ip, ":"); colon != -1 {
		ip = ip[:colon]
	}
	if ip == "" {
		ip = "127.0.0.1"
	}
	return ip
}
