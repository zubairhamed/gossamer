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

	case strings.HasPrefix(e, "FeaturesOfInterest"):
		return ENTITY_FEATURESOFINTEREST

	default:
		return ENTITY_UNKNOWN
	}
}

func IsEntity(e string) bool {
	if 	strings.HasPrefix(e, "Things") ||
		strings.HasPrefix(e, "Locations") ||
		strings.HasPrefix(e, "Datastreams") ||
		strings.HasPrefix(e, "Sensors") ||
		strings.HasPrefix(e, "Observations") ||
		strings.HasPrefix(e, "ObservedProperties") ||
		strings.HasPrefix(e, "FeaturesOfInterest") {
		return true
	}
	return false
}

var ERR_INVALID_ENTITY = errors.New("Invalid Entity")

func (s *DefaultServer) handleGetEntity(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Println("=====================")
	CreateRequest(r.URL)
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
