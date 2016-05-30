package toolkit

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"net/http/httptest"
)

func TestClientID(t *testing.T) {
	req, _ := http.NewRequest("GET", `http://www.test.com/`, nil)
	res := httptest.NewRecorder()
	cid := ClientID(req, res, "cid")
	assert.NotEmpty(t, cid)
	assert.NotEmpty(t, res.Header().Get("Set-Cookie"))
}
