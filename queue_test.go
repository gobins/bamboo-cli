package main

import "testing"

func TestRunPlanBranch(t *testing.T) {
	cred := new(credentials)

	cred.baseurl = "http://192.168.99.100:8085"
	cred.username = "admin"
	cred.password = "password"

	runPlanBranch(*cred, "SP-MSP0")

}
