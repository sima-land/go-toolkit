package toolkit

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientID(t *testing.T) {
	req, _ := http.NewRequest("GET", `http://www.test.com/`, nil)
	res := httptest.NewRecorder()
	cid := ClientID(req, res, "cid")
	assert.NotEmpty(t, cid)
	assert.NotEmpty(t, res.Header().Get("Set-Cookie"))
}

func TestEmptyClientID(t *testing.T) {
	req, _ := http.NewRequest("GET", `http://www.test.com/`, nil)
	res := httptest.NewRecorder()
	req.AddCookie(&http.Cookie{
		Value: "",
		Name:  "cid",
	})
	cid := ClientID(req, res, "cid")
	assert.NotEmpty(t, cid)
	assert.NotEmpty(t, res.Header().Get("Set-Cookie"))
}

func TestClientID(t *testing.T) {
	req, _ := http.NewRequest("GET", `http://www.test.com/`, nil)
	res := httptest.NewRecorder()
	req.AddCookie(&http.Cookie{
		Value: "111",
		Name:  "cid",
	})
	cid := ClientID(req, res, "cid")
	assert.NotEmpty(t, cid)
	assert.Equal(t, "111", cid)
}
