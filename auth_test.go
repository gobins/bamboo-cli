package main

import "testing"

func TestGetAuthToken(t *testing.T) {
	cred := new(credentials)

	cred.baseurl = "http://192.168.99.100:8085"
	cred.username = "admin"
	cred.password = "password"
	resp := getAuthToken(*cred)

	if resp == "" {
		t.Error("Error")
	}
}
