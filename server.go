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
	sensingHandler SensingProfileHandler
	taskingHandler TaskingProfileHandler
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

	case strings.HasPrefix(e, "Locations"):
		return ENTITY_LOCATIONS

	case strings.HasPrefix(e, "Datastreams"):
		return ENTITY_DATASTREAMS

	case strings.HasPrefix(e, "Sensors"):
		return ENTITY_SENSORS

	case strings.HasPrefix(e, "Observations"):
		return ENTITY_OBSERVATIONS

	case strings.HasPrefix(e, "ObservedProperties"):
		return ENTITY_OBSERVEDPROPERTIES

	case strings.HasPrefix(e, "FeaturesOfInterests"):
		return ENTITY_FEATURESOFINTERESTS

	case strings.HasPrefix(e, "HistoricalLocations"):
		return ENTITY_HISTORICALLOCATIONS

	default:
		return ENTITY_UNKNOWN
	}
}

func IsEntity(e string) bool {
	if strings.HasPrefix(e, "Things") ||
		strings.HasPrefix(e, "Locations") ||
		strings.HasPrefix(e, "HistoricalLocations") ||
		strings.HasPrefix(e, "Datastreams") ||
		strings.HasPrefix(e, "Sensors") ||
		strings.HasPrefix(e, "Observations") ||
		strings.HasPrefix(e, "ObservedProperties") ||
		strings.HasPrefix(e, "FeaturesOfInterests") {
		return true
	}
	return false
}

var ERR_INVALID_ENTITY = errors.New("Invalid Entity")

func (s *DefaultServer) handleGetEntity(c web.C, w http.ResponseWriter, r *http.Request) {

	req, err := CreateRequest(r.URL)
	if err != nil {
		log.Println(err)
	}

	navPath := req.GetNavigation().GetItems()
	l := len(navPath)

	if l == 0 {
		// TODO: Throw error
	}

	c1 := make(chan bool)
	var lastEntity EntityType
	var lastEntityId string
	go func() {

		for _, v := range navPath {
			c2 := make(chan bool)
			go func() {
				// log.Println(s.dataStore, v.GetEntity(), v.GetId(), v.GetQueryOptions(), lastEntity)
				s.dataStore.Get(v.GetEntity(), v.GetId(), v.GetQueryOptions(), lastEntity, lastEntityId)
				lastEntity = v.GetEntity()
				lastEntityId = v.GetId()
				c2 <- true
			}()
			<- c2


		}
		c1 <- true
	}()
	<- c1
	/*
		/Things(1)/Observations
		query(navItem, lastQueryVal) {
			if lastQueryVal == nil {
				switch navItem {
					case navItem.type && id not set:
					case navItem.type && id set:
				}
			} else {

			}
		}

		var lastQueryVal interface{}
		for navItem := range navItems
			lastQueryVal = query(navItem, lastQueryVal)
			if !hasNextNav {

			}


		query(navEntity, lastQueryVal) {
			switch navEntity.type
				case Entity:
					if
		}


		var lastQueryVal interface{}
		var valueOut interface{}
		for navEntity in navigationItems
			lastQueryVal = query(navEntity, lastQueryVal)

			if lastItem
				if hasProperty
				if hasValue


		for navigationItems
			fetch
				- all
				- one with id

			if lasItem
				if has property || has value {
					writeOut property or Value
				} else {
					writeOut lastQueryVal
				}
			else
				continue

	*/
}

func (s *DefaultServer) query() {

}

func (s *DefaultServer) Start() {
	goji.Get("/v1.0", s.handleRootResource)
	goji.Get("/v1.0/*", s.handleGetEntity)

	log.Println("Start Server")
	goji.Serve()
}

func (s *DefaultServer) Stop() {
	log.Println("Stopped Server")
}

func (s *DefaultServer) UseSensingProfile(h SensingProfileHandler) {
	s.sensingHandler = h
}

func (s *DefaultServer) UseTaskingProfile(h TaskingProfileHandler) {
	s.taskingHandler = h
}
