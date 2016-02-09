package bamboocli

import (
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type credentials struct {
	username string
	password string
	baseurl  string
}

func getAuthToken(cred credentials) string {
	log.Debug("basicAuth to bamboo server")
	var token []string
	client := &http.Client{}
	req, _ := http.NewRequest("GET", cred.baseurl, nil)
	req.SetBasicAuth(cred.username, cred.password)
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Authentication Error")
	}
	log.Debug(resp)
	cookie := resp.Header.Get("Set-Cookie")
	if len(cookie) != 0 {
		splitcookie := strings.Split(cookie, ";")

		for _, item := range splitcookie {
			if strings.Contains(item, "atl.xsrf.token") {
				token = strings.Split(item, "=")
			}
		}
	}

	return token[1]
}
