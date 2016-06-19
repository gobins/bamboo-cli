package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
)

func describePlan(cred credentials, planKey string) plan {
	log.Debug("Retrieving plan details for key: ", planKey)
	if cred.token == "" {
		cred.token = getAuthToken(cred)
	}
	cred.apiuri = "/rest/api/latest/plan/" + planKey + ".json"

	resp := httpclient(cred, "GET")

	var myplan plan
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error reading response body")
	}

	json.Unmarshal(body, &myplan)
	log.Debug(myplan)
	return myplan
}
