package toolkit

import (
	"net/http"
	"time"

	"github.com/jmcvetta/randutil"
)

// ClientID returns randomly generated client's ID
func ClientID(req *http.Request, rw http.ResponseWriter, cookieName string) string {
	var (
		cookie *http.Cookie
		res    string
	)
	cookie, err := req.Cookie(cookieName)
	if err != nil || (cookie != nil && cookie.Value == "") {
		res, err = randutil.AlphaString(16)
		if err != nil {
			panic(err)
		}
		cookie := &http.Cookie{
			Name:    cookieName,
			Value:   res,
			Path:    "/",
			Domain:  req.URL.Host,
			Expires: time.Now().AddDate(1, 0, 0),
		}
		http.SetCookie(rw, cookie)
	} else {
		res = cookie.Value
	}
	return res
}
