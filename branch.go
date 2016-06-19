package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
)

func getAllBranches(cred credentials, plan_key string) branches {
	log.Debug("Retrieving all branches for plan: " + plan_key)

	if cred.token == "" {
		cred.token = getAuthToken(cred)
	}

	cred.apiuri = "/rest/api/latest/plan/" + plan_key + ".json?expand=branches"
	log.Debug("Calling Bamboo with uri: " + cred.apiuri)

	resp := httpclient(cred, "GET")
	var pl plan
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error reading response body")
	}
	json.Unmarshal(body, &pl)
	return pl.Branches
}

// func getAllStages(cred *credentials, plan_key string) {

// }

// func getAllJobs(cred *credentials, plan_key string) {

// }
