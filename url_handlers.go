package gossamer

import (
	"encoding/json"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
)

type ResourceUrlType struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (s *DefaultServer) handleRootResource(c web.C, w http.ResponseWriter, r *http.Request) {
	/*
	ResolveSelfLinkUrl("", )
	 */

	data :=
		[]ResourceUrlType{
			{"Things", ResolveSelfLinkUrl("", ENTITY_THINGS)},
			{"Locations", ResolveSelfLinkUrl("", ENTITY_LOCATIONS)},
			{"Datastreams", ResolveSelfLinkUrl("", ENTITY_DATASTREAMS)},
			{"Sensors", ResolveSelfLinkUrl("", ENTITY_SENSORS)},
			{"Observations", ResolveSelfLinkUrl("", ENTITY_OBSERVATIONS)},
			{"ObservedProperties", ResolveSelfLinkUrl("", ENTITY_OBSERVEDPROPERTIES)},
			{"FeaturesOfInterest", ResolveSelfLinkUrl("", ENTITY_FEATURESOFINTEREST)},
	}

	v := &ValueList{
		Value: data,
	}
	out, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		log.Println(err)
	}
	w.Write(out)
}
