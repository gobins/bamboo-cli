package bamboocli

import "testing"

func TestDescribePlan(t *testing.T) {
	cred := new(credentials)

	cred.baseurl = "http://192.168.99.100:8085"
	cred.username = "admin"
	cred.password = "password"
	cred.apiuri = "/rest/api/latest/project.json"

	describePlan(*cred, "SP-MSP")
}
