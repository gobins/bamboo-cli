package bamboocli

import (
	"bufio"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type credentials struct {
	username string
	password string
	baseurl  string
}

func getAuthToken(cred credentials) *http.Response {
	log.Debug("basicAuth to bamboo server")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", cred.baseurl, nil)
	req.SetBasicAuth(cred.username, cred.password)
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Authentication Error")
	}
	log.Debug(resp)
	// log.Info(resp.Header)
	// for k, v := range resp.Header {
	// 	log.Println("key:", k, "value:", v)
	// }
	log.Info(resp.Header.Get("Set-Cookie value"))
	reader := bufio.NewReader(strings.NewReader(resp.Header.Get("Set-Cookie")))

	logReq, err := http.ReadRequest(reader)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(logReq.Header)
	// cookies := resp.Cookies()
	// for _, cookie := range cookies {
	// 	log.Info(cookie)
	// }

	return resp
}
