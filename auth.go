package bambooapi

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type credentials struct {
	username string
	password string
	baseurl  string
}

func basicAuth(cred credentials) {
	log.Debug("basicAuth to bamboo server")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", cred.baseurl, nil)
	req.SetBasicAuth(cred.username, cred.password)
	resp, err := client.Do(req)
	if err != nil {

	}
}
