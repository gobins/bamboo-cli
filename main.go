package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "bamboocli"
	app.Usage = "Command Line Interfact for Atlassian Bamboo"

	var username string
	var password string
	var bamboo_url string
	var plan_key string
	var project_key string
	var debug bool

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "username",
			Usage:       "username for bamboo",
			Destination: &username,
		},
		cli.StringFlag{
			Name:        "password",
			Usage:       "password for bamboo",
			Destination: &password,
		},
		cli.StringFlag{
			Name:        "url",
			Usage:       "url for bamboo",
			Destination: &bamboo_url,
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Set debug to see debug logs",
			Destination: &debug,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "get-projects",
			Usage: "Lists all projects",
			Action: func(C *cli.Context) {
				if debug {
					log.SetLevel(log.DebugLevel)
				}
				var errorFlag bool
				if username == "" {
					fmt.Println("username is required")
					errorFlag = true
				}
				if password == "" {
					fmt.Println("password is required")
					errorFlag = true
				}
				if bamboo_url == "" {
					fmt.Println("url is required")
					errorFlag = true
				}
				if errorFlag == true {
					os.Exit(1)
				}
				log.Debug("Calling function getAllProjects")
				CallGetAllProjects(username, password, bamboo_url)
			},
		},
		{
			Name:  "get-plans",
			Usage: "Lists all plans in a project",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "key",
					Usage:       "Key of the Project",
					Destination: &project_key,
				},
			},
			Action: func(C *cli.Context) {
				var errorFlag bool
				if debug {
					log.SetLevel(log.DebugLevel)
				}
				if username == "" {
					fmt.Println("username is required")
					errorFlag = true
				}
				if password == "" {
					fmt.Println("password is required")
					errorFlag = true
				}
				if bamboo_url == "" {
					fmt.Println("url is required")
					errorFlag = true
				}
				if project_key == "" {
					fmt.Println("key is required")
					errorFlag = true
				}
				if errorFlag == true {
					os.Exit(1)
				}
				log.Debug("Calling function CallGetAllPlansInProject")
				CallGetAllPlansInProject(username, password, bamboo_url, project_key)
			},
		},
		{
			Name:  "get-branches",
			Usage: "Lists all branches in a plan",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "key",
					Usage:       "Key of the Plan",
					Destination: &plan_key,
				},
			},
			Action: func(C *cli.Context) {
				var errorFlag bool
				if debug {
					log.SetLevel(log.DebugLevel)
				}
				if username == "" {
					fmt.Println("username is required")
					errorFlag = true
				}
				if password == "" {
					fmt.Println("password is required")
					errorFlag = true
				}
				if bamboo_url == "" {
					fmt.Println("url is required")
					errorFlag = true
				}
				if plan_key == "" {
					fmt.Println("Plan Key is required")
					errorFlag = true
				}
				if errorFlag == true {
					os.Exit(1)
				}
				log.Debug("Calling function CallGetAllBranches")
				CallGetAllBranches(username, password, bamboo_url, plan_key)
			},
		},
	}

	app.Run(os.Args)
}
