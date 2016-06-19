package main

import (
	"encoding/json"
	"fmt"
)

type planKey struct {
	Key string `json:"key"`
}

//Link attribute which holds the url to the object
type Link struct {
	Href string `json:"href"`
}

type project struct {
	Name  string `json:"name"`
	Key   string `json:"key"`
	Link  Link   `json:"link"`
	Plans plans  `json:"plans"`
}

type plans struct {
	Plans []plan `json:"plan"`
}

//Projects containing multiple project definition
type Projects struct {
	Projects []project `json:"project"`
}

//AllProjects parent element for all projects
type AllProjects struct {
	Projects Projects `json:"projects"`
}

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
	Size      int64    `json:"size"`
	AllBranch []branch `json:"branch"`
}

type branch struct {
	Shortname string `json:"shortName"`
	ShortKey  string `json:"shortKey"`
	Enabled   bool   `json:"enabled"`
	Link      Link   `json:"link"`
	Key       string `json:"key"`
}

type postresp struct {
	PlanKey        string `json:"planKey"`
	BuildNumber    int64  `json:"buildNumber"`
	BuildResultKey string `json:"buildResultKey"`
	TriggerReason  string `json:"triggerReason"`
	Link           Link   `json:"link"`
}

func js(what string, data interface{}) {
	final_struct := make(map[string]interface{})
	final_struct[what] = data
	js, _ := json.MarshalIndent(final_struct, "", "  ")
	fmt.Println(string(js))
}
