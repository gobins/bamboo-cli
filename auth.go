package main

import (
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type credentials struct {
	username string
	password string
	baseurl  string
	apiuri   string
	token    string
}

func getAuthToken(cred credentials) string {
	log.Debug("basicAuth to bamboo server")
	var token string
	client := &http.Client{}
	req, _ := http.NewRequest("GET", cred.baseurl, nil)
	req.SetBasicAuth(cred.username, cred.password)
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Authentication Error")
	}
	cookies := resp.Cookies()
	log.Debug(cookies)
	if len(cookies) != 0 {
		for _, cookie := range cookies {
			if strings.Contains(cookie.Name, "atl.xsrf.token") {
				token = cookie.Value
			}
		}
	}
	log.Debug("Token:", token)
	return token
}
