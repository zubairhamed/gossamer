package gossamer

import (
	"errors"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
	"strings"
)

func NewServer() Server {
	return &DefaultServer{}
}

type DefaultServer struct {
	dataStore      Datastore
}

func (s *DefaultServer) handleNotImplemented(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Println("Not implemented")
}

func (s *DefaultServer) UseStore(ds Datastore) {
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

func (s *DefaultServer) Start() {
	goji.Get("/v1.0", s.handleRootResource)
	goji.Get("/v1.0/", s.handleRootResource)

	goji.Get("/v1.0/*", s.handleGet)
	goji.Get("/v1.0/*", s.handlePost)
	goji.Get("/v1.0/*", s.handlePut)
	goji.Get("/v1.0/*", s.handleDelete)
	goji.Get("/v1.0/*", s.handlePatch)

	log.Println("Start Server")
	goji.Serve()
}

func (s *DefaultServer) Stop() {
	log.Println("Stopped Server")
}
