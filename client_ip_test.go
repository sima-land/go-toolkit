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
