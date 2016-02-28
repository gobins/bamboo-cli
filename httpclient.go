package bamboocli

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func httpclient(cred credentials, method string) *http.Response {
	log.Debug("httpclient function")
	if cred.token == "" {
		cred.token = getAuthToken(cred)
	}
	httpclient := &http.Client{}
	req, _ := http.NewRequest(method, cred.baseurl+cred.apiuri, nil)

	//Adding auth cookie to call bamboo REST API
	tokencookie := &http.Cookie{Name: "atl.xsrf.token", Value: cred.token, HttpOnly: false}
	req.AddCookie(tokencookie)
	resp, err := httpclient.Do(req)
	if err != nil {
		log.Error("Error calling url:", cred.baseurl+cred.apiuri)
	}
	log.Debug(resp)
	return resp
}
