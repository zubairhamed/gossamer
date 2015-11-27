package gossamer

//type ResourceUrlType struct {
//	name	string 	`json: "name"`
//	url 	string	`json: "url"`
//}

type RootResourceResponse struct {
	value 	[]ResourceUrlType 	`json: "value"`
}