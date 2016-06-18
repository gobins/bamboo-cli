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
