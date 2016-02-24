package bamboocli

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type project struct {
	Name string `xml:"name,attr"`
	Key  string `xml:"key,attr"`
	Link string `xml:"key,href"`
}

type Allprojects struct {
	projects []project
}

func getAllProjects(cred credentials) {
	log.Debug("Retrieving all projects")
	token := getAuthToken(cred)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", cred.baseurl+cred.apiuri, nil)
	tokencookie := &http.Cookie{Name: "atl.xsrf.token", Value: token, HttpOnly: false}
	req.AddCookie(tokencookie)
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Error retrieving projects")
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Error("Error reading response body")
		}
		log.Info(string(body))

		var pr Allprojects

		xml.Unmarshal(body, &pr)

		log.Info(pr.projects)
	}

	//log.Info(resp.Body)
}
