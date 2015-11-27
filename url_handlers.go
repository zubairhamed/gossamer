package gossamer
import (
	"github.com/zenazn/goji/web"
	"net/http"
	"encoding/json"
	"log"
)

type ResourceUrlType struct {
	Name	string 	`json: "name"`
	Url 	string	`json: "url"`
}

func (s *DefaultServer) handleRootResource(c web.C, w http.ResponseWriter, r *http.Request) {
	data := struct {
		Value []ResourceUrlType `json:"value"`
	} {
		[]ResourceUrlType{
			{"Things", "http://example.org/v1.0/Things", },
			{"Locations", "http://example.org/v1.0/Locations", },
			{"Datastreams", "http://example.org/v1.0/Datastreams", },
			{"Sensors", "http://example.org/v1.0/Sensors", },
			{"Observations", "http://example.org/v1.0/Observations", },
			{"ObservedProperties", "http://example.org/v1.0/ObservedProperties", },
			{"FeaturesOfInterest", "http://example.org/v1.0/FeaturesOfInterest", },
		},
	}
	out, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))
}


