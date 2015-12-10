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

	log.Println(req)

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
	currIndex 	int
	items 		[]ResourcePathItem
	property      string
	propertyValue string
}

func (r *SensorThingsResourcePath) Next() ResourcePathItem {
	return r.At(r.currIndex+1)
}

func (r *SensorThingsResourcePath) Prev() ResourcePathItem {
	return r.At(r.currIndex-1)
}

func (r *SensorThingsResourcePath) Current() ResourcePathItem {
	return r.At(r.currIndex)
}

func (r *SensorThingsResourcePath) First() ResourcePathItem {
	return r.At(0)
}

func (r *SensorThingsResourcePath) Last() ResourcePathItem {
	return r.At(r.Size()-1)
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

func (r *SensorThingsResourcePath) GoNext() {
	r.currIndex++
}

func (r *SensorThingsResourcePath) GoPrev() {
	r.currIndex--
}
func (r *SensorThingsResourcePath) GoFirst() {
	r.currIndex = 0
}
func (r *SensorThingsResourcePath) GoLast() {
	r.currIndex = r.Size()-1
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