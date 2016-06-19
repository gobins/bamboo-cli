package main

import (
	"fmt"
	"github.com/apcera/termtables"
)

func CallGetAllProjects(username, password, bamboo_url string) {
	cred := new(credentials)

	cred.baseurl = bamboo_url
	cred.username = username
	cred.password = password

	allProjects := getAllProjects(*cred)
	table := termtables.CreateTable()
	table.AddHeaders("Name", "Key", "Link")
	if len(allProjects) != 0 {
		for _, proj := range allProjects {
			table.AddRow(proj.Name, proj.Key, proj.Link.Href)
		}
	}
	fmt.Println(table.Render())
}

func CallGetAllPlansInProject(username, password, bamboo_url, project_name string) {
	cred := new(credentials)

	cred.baseurl = bamboo_url
	cred.username = username
	cred.password = password

	allPlans := getAllPlansInProject(*cred, project_name)
	table := termtables.CreateTable()
	table.AddHeaders("Name", "Key", "Active", "Building", "Stages", "Branches", "Link")
	if len(allPlans.Plans) != 0 {
		for _, pl := range allPlans.Plans {
			table.AddRow(pl.BuildName, pl.Key, pl.IsActive, pl.IsBuilding, pl.Stages.Size, pl.Branches.Size, pl.Link.Href)
		}
	}
	fmt.Println(table.Render())
}

func CallGetAllBranches(username, password, bamboo_url, plan_key string) {
	cred := new(credentials)

	cred.baseurl = bamboo_url
	cred.username = username
	cred.password = password

	allBranches := getAllBranches(*cred, plan_key)
	table := termtables.CreateTable()
	table.AddHeaders("Name", "Enabled", "Key", "Link")
	if len(allBranches.AllBranch) != 0 {
		for _, br := range allBranches.AllBranch {
			table.AddRow(br.Shortname, br.Enabled, br.Key, br.Link.Href)
		}
	}
	fmt.Println(table.Render())
}
