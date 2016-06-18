package bamboocli

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
)

type plan struct {
	Name                      string   `json:"name"`
	ProjectKey                string   `json:"projectKey"`
	ProjectName               string   `json:"projectName"`
	BuildName                 string   `json:"buildName"`
	Key                       string   `json:"key"`
	Link                      Link     `json:"link"`
	IsActive                  bool     `json:"isActive"`
	IsBuilding                bool     `json:"isBuilding"`
	AverageBuildTimeInSeconds int64    `json:"averageBuildTimeInSeconds"`
	Stages                    stages   `json:"stages"`
	Branches                  branches `json:"branches"`
}

type stages struct {
	Size int64 `json:"size"`
}

type branches struct {
	Size int64 `json:"size"`
}

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
