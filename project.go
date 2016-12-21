package main

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

func getAllProjects(cred credentials) []project {
	log.Debug("Retrieving all projects")
	if cred.token == "" {
		cred.token = getAuthToken(cred)
	}
	if cred.apiuri == "" {
		cred.apiuri = "/rest/api/latest/project.json?max-results=500"
	}

	resp := httpclient(cred, "GET")
	var pr AllProjects
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error reading response body")
	}
	json.Unmarshal(body, &pr)
	return pr.Projects.Projects
}

func getAllPlansInProject(cred credentials, projectName string) plans {
	log.Debug("Retrieving all plans in the project")
	if cred.token == "" {
		cred.token = getAuthToken(cred)
	}
	cred.apiuri = "/rest/api/latest/project/" + projectName + ".json?expand=plans.plan&max-results=500"

	resp := httpclient(cred, "GET")

	var pr project
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Error("Error reading response body")
	}
	json.Unmarshal(body, &pr)
	return pr.Plans
}
