package gossamer

import (
	"encoding/json"
	"errors"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func NewServer() Server {
	return &GossamerServer{}
}

type GossamerServer struct {
	dataStore Datastore
}

func (s *GossamerServer) handleNotImplemented(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Println("Not implemented")
}

func (s *GossamerServer) UseStore(ds Datastore) {
	s.dataStore = ds
}

func DiscoverEntityType(e string) EntityType {
	switch {
	case strings.HasPrefix(e, "Things"):
		return ENTITY_THINGS

	case strings.HasPrefix(e, "Thing"):
		return ENTITY_THING

	case strings.HasPrefix(e, "Locations"):
		return ENTITY_LOCATIONS

	case strings.HasPrefix(e, "Location"):
		return ENTITY_LOCATION

	case strings.HasPrefix(e, "Datastreams"):
		return ENTITY_DATASTREAMS

	case strings.HasPrefix(e, "Datastream"):
		return ENTITY_DATASTREAM

	case strings.HasPrefix(e, "Sensors"):
		return ENTITY_SENSORS

	case strings.HasPrefix(e, "Sensor"):
		return ENTITY_SENSOR

	case strings.HasPrefix(e, "Observations"):
		return ENTITY_OBSERVATIONS

	case strings.HasPrefix(e, "Observation"):
		return ENTITY_OBSERVATION

	case strings.HasPrefix(e, "ObservedProperties"):
		return ENTITY_OBSERVEDPROPERTIES

	case strings.HasPrefix(e, "ObservedProperty"):
		return ENTITY_OBSERVEDPROPERTY

	case strings.HasPrefix(e, "FeaturesOfInterests"):
		return ENTITY_FEATURESOFINTERESTS

	case strings.HasPrefix(e, "FeaturesOfInterest"):
		return ENTITY_FEATURESOFINTEREST

	case strings.HasPrefix(e, "HistoricalLocations"):
		return ENTITY_HISTORICALLOCATIONS

	case strings.HasPrefix(e, "HistoricalLocation"):
		return ENTITY_HISTORICALLOCATION

	default:
		return ENTITY_UNKNOWN
	}
}

func IsEntity(e string) bool {
	if strings.HasPrefix(e, "Thing") ||
		strings.HasPrefix(e, "Things") ||
		strings.HasPrefix(e, "Location") ||
		strings.HasPrefix(e, "Locations") ||
		strings.HasPrefix(e, "HistoricalLocation") ||
		strings.HasPrefix(e, "HistoricalLocations") ||
		strings.HasPrefix(e, "Datastream") ||
		strings.HasPrefix(e, "Datastreams") ||
		strings.HasPrefix(e, "Sensor") ||
		strings.HasPrefix(e, "Sensors") ||
		strings.HasPrefix(e, "Observation") ||
		strings.HasPrefix(e, "Observations") ||
		strings.HasPrefix(e, "ObservedProperty") ||
		strings.HasPrefix(e, "ObservedProperties") ||
		strings.HasPrefix(e, "FeaturesOfInterest") ||
		strings.HasPrefix(e, "FeaturesOfInterests") {
		return true
	}
	return false
}

func IsSingularEntity(e string) bool {
	if (strings.HasPrefix(e, "Location") && !strings.HasPrefix(e, "Locations")) ||
		(strings.HasPrefix(e, "Thing") && !strings.HasPrefix(e, "Things")) ||
		(strings.HasPrefix(e, "HistoricalLocation") && !strings.HasPrefix(e, "HistoricalLocations")) ||
		(strings.HasPrefix(e, "Datastream") && !strings.HasPrefix(e, "Datastreams")) ||
		(strings.HasPrefix(e, "Sensor") && !strings.HasPrefix(e, "Sensors")) ||
		(strings.HasPrefix(e, "Observation") && !strings.HasPrefix(e, "Observations")) ||
		(strings.HasPrefix(e, "ObservedProperty") && !strings.HasPrefix(e, "ObservedProperties")) ||
		(strings.HasPrefix(e, "FeaturesOfInterest") && !strings.HasPrefix(e, "FeaturesOfInterests")) {
		return true
	}
	return false
}

var ERR_INVALID_ENTITY = errors.New("Invalid Entity")

func (s *GossamerServer) Start() {
	goji.Get("/v1.0", s.handleRootResource)
	goji.Get("/v1.0/", s.handleRootResource)

	goji.Get("/v1.0/*", s.handleGet)
	goji.Post("/v1.0/*", s.handlePost)
	goji.Put("/v1.0/*", s.handlePut)
	goji.Delete("/v1.0/*", s.handleDelete)
	goji.Patch("/v1.0/*", s.handlePatch)

	log.Println("Start Server")
	goji.Serve()
}

func (s *GossamerServer) Stop() {
	log.Println("Stopped Server")
}

func (s *GossamerServer) handleRootResource(c web.C, w http.ResponseWriter, r *http.Request) {
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
		ThrowHttpInternalServerError(MSG_ERR_HANDLING_REQUEST+err.Error(), w)
		return
	}
	w.Write(out)
}

func (s *GossamerServer) handleGet(c web.C, w http.ResponseWriter, r *http.Request) {
	req, err := CreateIncomingRequest(r.URL, HTTP)
	if err != nil {
		ThrowHttpBadRequest(MSG_ERR_HANDLING_REQUEST+err.Error(), w)
	}

	rp := req.GetResourcePath()
	result, err := s.dataStore.Query(rp, req.GetQueryOptions())
	if err != nil {
		ThrowHttpInternalServerError(MSG_ERR_HANDLING_REQUEST+err.Error(), w)
		return
	}

	var jsonOut interface{}
	if reflect.TypeOf(result).Kind() == reflect.Slice {
		count := reflect.ValueOf(result).Len()
		if req.GetQueryOptions().CountSet() {
			opt := req.GetQueryOptions().GetCountOption()
			if !opt.GetValue() {
				count = 0
			}
		}
		// Check for $count and include

		jsonOut = &ValueList{
			Count: count,
			Value: result,
		}
	} else {
		jsonOut = result
	}

	b, err := json.MarshalIndent(jsonOut, "", "  ")
	if err != nil {
		ThrowHttpInternalServerError(MSG_ERR_HANDLING_REQUEST+err.Error(), w)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		ThrowHttpInternalServerError(MSG_ERR_HANDLING_REQUEST+err.Error(), w)
		return
	}
}

func (s *GossamerServer) handlePost(c web.C, w http.ResponseWriter, r *http.Request) {
	var err error
	var req Request

	// TODO: Verify Content-Type == "application/json"

	req, err = CreateIncomingRequest(r.URL, HTTP)
	if err != nil {
		ThrowHttpBadRequest(MSG_ERR_HANDLING_REQUEST + err.Error(), w)
	}

	if err = ValidatePostRequestUrl(req); err != nil {
		ThrowHttpMethodNotAllowed(err.Error(), w)
		return
	}

	rp := req.GetResourcePath()
	ent := rp.Last().GetEntity()
	cont := rp.Containing()

	if !IsSingularEntity(string(ent)) {
		decoder := json.NewDecoder(r.Body)
		e, err := DecodeJsonToEntityStruct(decoder, ent)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+err.Error(), w)
			return
		}

		if cont != nil {
			SetAssociatedEntityId(e, cont.GetEntity(), cont.GetId())
		}

		rp := req.GetResourcePath()

		err = ValidateMandatoryProperties(e)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+err.Error(), w)
			return
		}

		err = ValidateIntegrityConstraints(e)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+err.Error(), w)
			return
		}

		err = s.dataStore.Insert(rp, e)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+err.Error(), w)
			return
		}
	}
}

func (s *GossamerServer) handlePut(c web.C, w http.ResponseWriter, r *http.Request) {

}

func (s *GossamerServer) handleDelete(c web.C, w http.ResponseWriter, r *http.Request) {

}

func (s *GossamerServer) handlePatch(c web.C, w http.ResponseWriter, r *http.Request) {

}
