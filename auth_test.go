package bamboocli

import "testing"

func TestGetAuthToken(t *testing.T) {
	cred := new(credentials)

	cred.baseurl = "http://192.168.56.101:8085"
	cred.username = "admin"
	cred.password = "password"
	resp := getAuthToken(*cred)

	if resp == "" {
		t.Error("Error")
	}
}
