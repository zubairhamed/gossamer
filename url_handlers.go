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
	data := []ResourceUrlType{
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
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println(err)
	}
	w.Write(out)
}

func (s *DefaultServer) handleGet(c web.C, w http.ResponseWriter, r *http.Request) {
	req, err := CreateRequest(r.URL)
	if err != nil {
		log.Println(err)
	}

}

func (s *DefaultServer) handlePost(c web.C, w http.ResponseWriter, r *http.Request) {

}

func (s *DefaultServer) handlePut(c web.C, w http.ResponseWriter, r *http.Request) {

}

func (s *DefaultServer) handleDelete(c web.C, w http.ResponseWriter, r *http.Request) {

}

func (s *DefaultServer) handlePatch(c web.C, w http.ResponseWriter, r *http.Request) {

}

type ResourcePathItem interface {

}

type ResourcePath interface {
	Next() ResourcePathItem
	Prev() ResourcePathItem
	Current() ResourcePathItem
	First() ResourcePathItem
	Last() ResourcePathItem

	IsLast() bool
	IsFirst() bool

	GoNext()
	GoPrev()
	GoFirst()
	GoLast()

	CurrentIndex() int
	Size() int
	Add(ResourcePathItem)
	At(int) ResourcePathItem
}

type SensorThingsResourcePath struct {

}

/*
Add
At
Length
 */