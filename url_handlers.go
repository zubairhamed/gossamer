package gossamer

import (
	"encoding/json"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
	"reflect"
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

	rp := req.GetResourcePath()

	result, err := s.dataStore.Query(rp)
	if err != nil {

	}

	var jsonOut interface{}
	if reflect.TypeOf(result).Kind() == reflect.Slice {
		// Check for $count and include

		jsonOut = &ValueList{
			Value: result,
		}
	} else {
		jsonOut = result
	}

	b, err := json.MarshalIndent(jsonOut, "", "  ")
	if err != nil {
		log.Println("Error converting to JSON")
	}

	_, err = w.Write(b)
	if err != nil {
		log.Println(err)
	}
	// http.Error(w, err.Error(), http.StatusBadRequest)
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
	GetEntity() EntityType
	GetId() string
	GetQueryOptions() QueryOptions
}

type SensorThingsResourcePathItem struct {
	entityType   EntityType
	entityId     string
	queryOptions QueryOptions
}

func (n *SensorThingsResourcePathItem) GetQueryOptions() QueryOptions {
	return n.queryOptions
}

func (n *SensorThingsResourcePathItem) GetEntity() EntityType {
	return n.entityType
}

func (n *SensorThingsResourcePathItem) GetId() string {
	return n.entityId
}


type ResourcePath interface {
	Next() ResourcePathItem
	Prev() ResourcePathItem
	Current() ResourcePathItem
	First() ResourcePathItem
	Last() ResourcePathItem
	All() []ResourcePathItem

	IsLast() bool
	IsFirst() bool
	HasNext() bool

	CurrentIndex() int
	Size() int
	Add(ResourcePathItem)
	At(int) ResourcePathItem
}

type SensorThingsResourcePath struct {
	currIndex 	int
	items 		[]ResourcePathItem
	property      string
	propertyValue string
}

func (r *SensorThingsResourcePath) Next() ResourcePathItem {
	r.currIndex++
	return r.At(r.currIndex)
}

func (r *SensorThingsResourcePath) Prev() ResourcePathItem {
	r.currIndex--
	return r.At(r.currIndex)
}

func (r *SensorThingsResourcePath) Current() ResourcePathItem {
	return r.At(r.currIndex)
}

func (r *SensorThingsResourcePath) First() ResourcePathItem {
	r.currIndex = 0
	return r.At(r.currIndex)
}

func (r *SensorThingsResourcePath) Last() ResourcePathItem {
	r.currIndex = r.Size()-1
	return r.At(r.currIndex)
}

func (r *SensorThingsResourcePath) HasNext() bool {
	if r.IsLast() {
		return false
	}
	return true
}

func (r *SensorThingsResourcePath) IsLast() bool {
	sz := r.Size()
	if r.CurrentIndex() == sz-1 {
		return true
	}
	return false
}

func (r *SensorThingsResourcePath) IsFirst() bool {
	if r.CurrentIndex() == 0 {
		return true
	}
	return false
}

func (r *SensorThingsResourcePath) CurrentIndex() int {
	return r.currIndex
}

func (r *SensorThingsResourcePath) Size() int {
	return len(r.items)
}

func (r *SensorThingsResourcePath) Add(rsrc ResourcePathItem) {
	r.items = append(r.items, rsrc)
}

func (r *SensorThingsResourcePath) At(idx int) ResourcePathItem {
	sz := r.Size()-1
	if idx > sz || idx < 0 {
		return nil
	}
	return r.items[idx]
}

func (r *SensorThingsResourcePath) All() []ResourcePathItem {
	return r.items
}