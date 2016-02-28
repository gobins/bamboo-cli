package bamboocli

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

type project struct {
	Name  string `json:"name"`
	Key   string `json:"key"`
	Link  Link   `json:"link"`
	Plans plans  `json:"plans"`
}

type plans struct {
	Plan []plan `json:"plan"`
}

type plan struct {
	Name      string  `json:"name"`
	ShortName string  `json:"shortName"`
	ShortKey  string  `json:"shortKey"`
	Enabled   string  `json:"enabled"`
	Link      Link    `json:"link"`
	Key       string  `json:"key"`
	PlanKey   planKey `json:"planKey"`
}

type planKey struct {
	Key string `json:"key"`
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
	if cred.apiuri == "" {
		cred.apiuri = "/rest/api/latest/project.json"
	}

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

func getAllPlansInProject(cred credentials, planName string) {
	log.Debug("Retrieving all plans in the project")
	if cred.token == "" {
		cred.token = getAuthToken(cred)
	}
	cred.apiuri = "/rest/api/latest/project.json?expand=projects.project.plans"

	resp := httpclient(cred, "GET")

	var pr AllProjects
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Error("Error reading response body")
	}

	json.Unmarshal(body, &pr)
	log.Info(pr.Projects.Projects)
}
