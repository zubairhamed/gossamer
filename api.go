package gossamer

type ProtocolType int
const (
	HTTP 	ProtocolType = 0
	COAP	ProtocolType = 1
)

type EntityType string
const (
	ENTITY_THINGS 				EntityType = "Things"
	ENTITY_OBSERVEDPROPERTIES 	EntityType = "ObservedProperties"
	ENTITY_LOCATIONS			EntityType = "Locations"
	ENTITY_DATASTREAMS			EntityType = "Datastreams"
	ENTITY_SENSORS				EntityType = "Sensors"
	ENTITY_OBSERVATIONS			EntityType = "Observations"
	ENTITY_FEATURESOFINTEREST	EntityType = "FeaturesOfInterest"
	ENTITY_UNKNOWN				EntityType = "UNKNOWN"
)

type QueryOptionType string
const (
	QUERYOPT_EXPAND		QueryOptionType = "$expand"
	QUERYOPT_SELECT		QueryOptionType = "$select"
	QUERYOPT_ORDERBY	QueryOptionType = "$orderby"
	QUERYOPT_TOP		QueryOptionType = "$top"
	QUERYOPT_SKIP		QueryOptionType = "$skip"
	QUERYOPT_COUNT		QueryOptionType = "$count"
	QUERYOPT_FILTER		QueryOptionType = "$filter"
	QUERYOPT_UNKNOWN 	QueryOptionType = "UNKNOWN"
)

type QueryOptions interface {
	ExpandSet() bool
	SelectSet() bool
	OrderBySet() bool
	TopSet() bool
	SkipSet() bool
	CountSet() bool
	FilterSet() bool
}

type Server interface {
	Stop()
	Start()
	UseSensingProfile(SensingProfileHandler)
	UseTaskingProfile(TaskingProfileHandler)
}

type NavigationItem interface {
	GetEntity() EntityType
	GetId() string
	GetQueryOptions() QueryOptions
}

type Navigation interface {
//	String()
	GetItems() []NavigationItem
//	GetProperty() string
//	GetPropertyValue() string
}

type Request interface {
	GetProtocol() ProtocolType
	GetQueryOptions() QueryOptions
	GetNavigation() Navigation
}

type SensingProfileHandler interface {
//	GetThings(Request)
//	GetThing(string, Request)
//
//	GetLocations(Request)
//	GetLocation(string, Request)
//
//	GetDatastreams(Request)
//	GetDatastream(string, Request)
//
//	GetSensors(Request)
//	GetSensor(string, Request)
//
//	GetObservations(Request)
//	GetObservation(string, Request)
//
//	GetObservedProperties(Request)
//	GetObservedProperty(string, Request)
//
//	GetFeaturesOfInterests(Request)
//	GetFeaturesOfInterest(string, Request)
}

type TaskingProfileHandler interface {

}