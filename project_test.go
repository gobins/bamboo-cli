package bamboocli

import "testing"

func TestGetAllProjects(t *testing.T) {
	cred := new(credentials)

	cred.baseurl = "http://localhost:8085"
	cred.username = "admin"
	cred.password = "password"
	cred.apiuri = "/rest/api/latest/project"
	token := getAuthToken(*cred)

	if token == "" {
		t.Error("Error")
	}
	getAllProjects(*cred)

}
