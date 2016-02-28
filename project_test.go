package bamboocli

import "testing"

func TestGetAllProjects(t *testing.T) {
	cred := new(credentials)

	cred.baseurl = "http://192.168.56.101:8085"
	cred.username = "admin"
	cred.password = "password"
	cred.apiuri = "/rest/api/latest/project"

	getAllProjects(*cred)

	// if len(resp) == 0 {
	// 	t.Error("Error")
	// }

}
