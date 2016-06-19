package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
)

func runPlanBranch(cred credentials, build_key string) {
	log.Debug("Trigger Plan Branch with Build Key: " + build_key)

	cred.apiuri = "/rest/api/latest/queue/" + build_key //+ "?stage&executeAllStages&customRevision"
	log.Debug("Calling Bamboo with uri: " + cred.apiuri)

	resp := httpPost(cred)
	var response postresp
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error reading response body")
	}
	json.Unmarshal(body, &response)
	log.Info(response)
}
