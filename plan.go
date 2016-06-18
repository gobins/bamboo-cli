package bamboocli

import (
	log "github.com/Sirupsen/logrus"
)

type plan struct {
	Name                      string   `json:"name"`
	ProjectKey                string   `json:"projectKey"`
	ProjectName               string   `json:"projectName"`
	BuildName                 string   `json:"buildName"`
	Key                       string   `json:"key"`
	Link                      Link     `json:"link"`
	IsActive                  string   `json:"isActive"`
	IsBuilding                string   `json:"isBuilding"`
	AverageBuildTimeInSeconds string   `json:"averageBuildTimeInSeconds"`
	Stages                    stages   `json:"stages"`
	Branches                  branches `json:"branches"`
}

type stages struct {
	Size string `json:"size"`
}

type branches struct {
	Size string `json:"size"`
}

func describePlan(cred credentials, planKey string) {
	log.Debug("Retrieving plan details for key: ", planKey)
}
