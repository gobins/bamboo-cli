package bamboocli

import (
	"encoding/json"
	"fmt"
)

type planKey struct {
	Key string `json:"key"`
}

//Link attribute which holds the url to the object
type Link struct {
	Href string `json:"href"`
}

func js(what string, data interface{}) {
	final_struct := make(map[string]interface{})
	final_struct[what] = data
	js, _ := json.MarshalIndent(final_struct, "", "  ")
	fmt.Println(string(js))
}
