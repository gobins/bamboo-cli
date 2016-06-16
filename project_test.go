package bamboocli

import "testing"

func TestGetAllProjects(t *testing.T) {
	cred := new(credentials)

	cred.baseurl = "http://192.168.99.100:8085"
	cred.username = "admin"
	cred.password = "password"
	cred.apiuri = "/rest/api/latest/project.json"

	getAllProjects(*cred)

}

func TestGetAllPlans(t *testing.T) {
	cred := new(credentials)

	cred.baseurl = "http://192.168.99.100:8085"
	cred.username = "admin"
	cred.password = "password"

	getAllPlansInProject(*cred, "First Project")

}
