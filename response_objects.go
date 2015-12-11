package gossamer

type RootResourceResponse struct {
	value []ResourceUrlType `json: "value"`
}

type ResourceUrlType struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
