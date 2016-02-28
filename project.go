package bamboocli

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

type project struct {
	Name string `json:"name"`
	Key  string `json:"key"`
	Link Link   `json:"link"`
}

//Link attribute which holds the url to the object
type Link struct {
	Href string `json:"href"`
}

//Projects containing multiple project definition
type Projects struct {
	Projects []project `json:"project"`
}

//AllProjects parent element for all projects
type AllProjects struct {
	Projects Projects `json:"projects"`
}

func getAllProjects(cred credentials) []project {
	log.Debug("Retrieving all projects")
	if cred.token == "" {
		cred.token = getAuthToken(cred)
	}
	cred.apiuri = "/rest/api/latest/project.json"

	resp := httpclient(cred, "GET")
	var pr AllProjects
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error reading response body")
	}

	json.Unmarshal(body, &pr)
	log.Debug(pr.Projects.Projects)
	return pr.Projects.Projects
}

// func getAllPlans(cred credentials, plan) {
//   log.Debug("Retrieving all plans in the project")
// 	if cred.token == "" {
// 		cred.token = getAuthToken(cred)
// 	}
// 	cred.apiuri = "/rest/api/latest/project"
//
//
// }
