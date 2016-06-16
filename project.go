package bamboocli

import (
	"encoding/json"
	"fmt"
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
	Plans []plan `json:"plan"`
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

func js(what string, data interface{}) {
	final_struct := make(map[string]interface{})
	final_struct[what] = data
	js, _ := json.MarshalIndent(final_struct, "", "  ")
	fmt.Println(string(js))
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
	//js("projects", pr.Projects.Projects)
	return pr.Projects.Projects
}

func getAllPlansInProject(cred credentials, projectName string) plans {
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

	var allplans plans
	var projFound bool
	for _, proj := range pr.Projects.Projects {
		if proj.Name == projectName || proj.Key == projectName {
			allplans = proj.Plans
			projFound = true
			break
		}
	}
	if !projFound {
		log.Info("Matching plan with name or key for found for: ", projectName)
	}
	return allplans
}
