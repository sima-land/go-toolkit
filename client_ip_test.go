package toolkit

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestClientIP(t *testing.T) {
	req, _ := http.NewRequest("GET", `http://www.test.com/`, nil)
	assert.Equal(t, "127.0.0.1", ClientIP(req))

	req.RemoteAddr = "192.168.1.1:8000"
	assert.Equal(t, "192.168.1.1", ClientIP(req))

	req.Header.Add("X-Forwarded-For", "zzz, 10.10.10.1, 192.168.1.2")
	assert.Equal(t, "192.168.1.2", ClientIP(req))

	req.Header.Add("X-Real-IP", "192.168.1.3")
	assert.Equal(t, "192.168.1.3", ClientIP(req))
}

func TestClientIP_IPv6(t *testing.T) {
	req, _ := http.NewRequest("GET", `http://www.test.com/`, nil)
	req.Header.Add("X-Forwarded-For", "66.249.83.52, 2a00:1fa0:27d:85c8:0:1d:97d9:d501")
	assert.Equal(t, "2a00:1fa0:27d:85c8:0:1d:97d9:d501", ClientIP(req))
}

func TestClientIP_IPv6WithPort(t *testing.T) {
	req, _ := http.NewRequest("GET", `http://www.test.com/`, nil)
	req.Header.Add("X-Forwarded-For", "66.249.83.52, [2a00:1fa0:27d:85c8:0:1d:97d9:d501]:8080")
	assert.Equal(t, "2a00:1fa0:27d:85c8:0:1d:97d9:d501", ClientIP(req))
}
