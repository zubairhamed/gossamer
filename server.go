package gossamer
import (
	"log"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
	"strings"
	"errors"
)

func NewServer() Server {
	return &DefaultServer{}
}

type DefaultServer struct {
	sensingHandler 	SensingProfileHandler
	taskingHandler 	TaskingProfileHandler
	dataStore 		Datastore
}

func (s *DefaultServer) handleNotImplemented(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Println("Not implemented")
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
	if 	strings.HasPrefix(e, "Things") ||
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

	// /Things(1)/Observations
	entities := []SensorThingEntity{}
	go func() {
		for idx, v := range navPath {
			switch v.GetEntity() {
			case ENTITY_THINGS:
				if v.GetId() == "" {
					s.dataStore.GetThings()
				} else {
					s.dataStore.GetThing(v.GetId())
				}
				break
			}

			if (idx+1) < l {
				log.Println("Has more items..:")
				if v.GetId() != "" {

				}
			} else {
				log.Println("Is last item..:")
			}
			log.Println(idx, v)
		}
		c1 <- true
	}()

	/*
		for navigationItems
			fetch
				- all
				- one with id

			if lasItem
				if has property
				if has value
				write out
			else
				continue

	 */
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
