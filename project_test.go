package bamboocli

import "testing"

func TestGetAllProjects(t *testing.T) {
	cred := new(credentials)

	cred.baseurl = "http://192.168.56.101:8085"
	cred.username = "admin"
	cred.password = "password"
	cred.apiuri = "/rest/api/latest/project"
	token := getAuthToken(*cred)

	if token == "" {
		t.Error("Error")
	}
	getAllProjects(*cred)

}
