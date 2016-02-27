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
	Link Link   `xml:"link"`
}

//Link attribute which holds the url to the object
type Link struct {
	Href string `xml:"href,attr"`
}

//AllProjects parent element for all projects
type AllProjects struct {
	Projects Projects `xml:"projects"`
	Link     Link     `xml:"link"`
}

//Projects containing multiple project definition
type Projects struct {
	Projects []project `xml:"project"`
}

// type Allprojects struct {
// 	//XMLName  xml.Name  `xml:projects"`
// 	projects []project `xml:project"`
// }

func getAllProjects(cred credentials) []project {
	log.Debug("Retrieving all projects")
	token := getAuthToken(cred)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", cred.baseurl+cred.apiuri, nil)
	tokencookie := &http.Cookie{Name: "atl.xsrf.token", Value: token, HttpOnly: false}
	req.AddCookie(tokencookie)
	resp, err := client.Do(req)
	var pr AllProjects
	if err != nil {
		log.Error("Error retrieving projects")
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error("Error reading response body")
		}
		log.Debug(string(body))
		xml.Unmarshal(body, &pr)
	}
	return pr.Projects.Projects
}
