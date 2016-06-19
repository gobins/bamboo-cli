package main

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

func httpPost(cred credentials) *http.Response {
	log.Debug("httpPost function")

	httpclient := &http.Client{}
	req, _ := http.NewRequest("POST", cred.baseurl+cred.apiuri, nil)
	req.SetBasicAuth(cred.username, cred.password)
	req.Header.Set("Accep", "application/json")
	resp, err := httpclient.Do(req)
	if err != nil {
		log.Error("Error calling url:", cred.baseurl+cred.apiuri)
	}
	log.Debug(resp)
	return resp
}
