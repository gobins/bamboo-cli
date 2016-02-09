package bamboocli

import "testing"

func TestGetAuthToken(t *testing.T) {
	cred := new(credentials)

	cred.baseurl = "http://localhost:8085"
	cred.username = "admin"
	cred.password = "password"
	resp := getAuthToken(*cred)

	if resp == "" {
		t.Error("Error")
	}
}
