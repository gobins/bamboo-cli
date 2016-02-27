package bamboocli

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type HTTPClient struct {
}

func (client HTTPClient) httpclient(cred credentials, method string) {
	log.Debug("httpclient function")
	if cred.token == "" {
		cred.token = getAuthToken(cred)
	}
	client := &http.Client{}
	req, _ := http.NewRequest(method, cred.baseurl+cred.apiuri, nil)
	tokencookie := &http.Cookie{Name: "atl.xsrf.token", Value: cred.token, HttpOnly: false}
	req.AddCookie(tokencookie)
	resp, err := client.Do(req)
}
