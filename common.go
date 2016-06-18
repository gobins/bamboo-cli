package bamboocli

type planKey struct {
	Key string `json:"key"`
}

//Link attribute which holds the url to the object
type Link struct {
	Href string `json:"href"`
}
