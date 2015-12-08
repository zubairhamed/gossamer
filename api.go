package gossamer

import "time"

type ProtocolType int

const (
	HTTP ProtocolType = 0
	COAP ProtocolType = 1
)

type EntityType string

const (
	ENTITY_THING               EntityType = "Thing"
	ENTITY_THINGS              EntityType = "Things"
	ENTITY_OBSERVEDPROPERTY    EntityType = "ObservedProperty"
	ENTITY_OBSERVEDPROPERTIES  EntityType = "ObservedProperties"
	ENTITY_LOCATION            EntityType = "Location"
	ENTITY_LOCATIONS           EntityType = "Locations"
	ENTITY_DATASTREAM          EntityType = "Datastream"
	ENTITY_DATASTREAMS         EntityType = "Datastreams"
	ENTITY_SENSOR              EntityType = "Sensor"
	ENTITY_SENSORS             EntityType = "Sensors"
	ENTITY_OBSERVATION         EntityType = "Observation"
	ENTITY_OBSERVATIONS        EntityType = "Observations"
	ENTITY_FEATURESOFINTEREST  EntityType = "FeaturesOfInterest"
	ENTITY_FEATURESOFINTERESTS EntityType = "FeaturesOfInterests"
	ENTITY_HISTORICALLOCATION  EntityType = "HistoricalLocation"
	ENTITY_HISTORICALLOCATIONS EntityType = "HistoricalLocations"
	ENTITY_UNKNOWN             EntityType = "UNKNOWN"
)

type QueryOptionType string

const (
	QUERYOPT_EXPAND  QueryOptionType = "$expand"
	QUERYOPT_SELECT  QueryOptionType = "$select"
	QUERYOPT_ORDERBY QueryOptionType = "$orderby"
	QUERYOPT_TOP     QueryOptionType = "$top"
	QUERYOPT_SKIP    QueryOptionType = "$skip"
	QUERYOPT_COUNT   QueryOptionType = "$count"
	QUERYOPT_FILTER  QueryOptionType = "$filter"
	QUERYOPT_UNKNOWN QueryOptionType = "UNKNOWN"
)

type QueryOptions interface {
	ExpandSet() bool
	SelectSet() bool
	OrderBySet() bool
	TopSet() bool
	SkipSet() bool
	CountSet() bool
	FilterSet() bool

	GetExpandOption() ExpandOption
	GetSelectOption() SelectOption
	GetOrderByOption() OrderByOption
	GetTopOption() TopOption
	GetSkipOption() SkipOption
	GetCountOption() CountOption
	GetFilterOption() FilterOption
}

type QueryOption interface {
	GetType() QueryOptionType
}

type ExpandOption interface {
	QueryOption
	GetValue() []string
}

type SelectOption interface {
	QueryOption
	GetValue() []string
}

type OrderByOption interface {
	QueryOption
	GetValue() []OrderByOptionValue
}

// asc, desc
type OrderByOptionValue interface {
}

type TopOption interface {
	QueryOption
	GetValue() int
}

type SkipOption interface {
	QueryOption
	GetValue() int
}

type CountOption interface {
	QueryOption
	GetValue() bool
}

type FilterOption interface {
	QueryOption
}

type Server interface {
	Stop()
	Start()
	UseSensingProfile(SensingProfileHandler)
	UseTaskingProfile(TaskingProfileHandler)
	UseStore(Datastore)
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

type Datastore interface {
	//	GetThings() []Thing
	//	GetThing(string) Thing

	//	GetObservedProperties() []ObservedProperty
	//	GetObservedProperty(string)
	//
	//	GetLocations() []Location
	//	GetLocation(string) Location
	//
	//	GetDatastreams() []Datastream
	//	GetDatastream(string) Datastream
	//
	//	GetSensors() []Sensor
	//	GetSensor(string) Sensor
	//
	//	GetObservations() []Observation
	//	GetObservation(string) Observation
	//
	//	GetFeatureOfInterests() []FeatureOfInterest
	//	GetFeatureOfInterest(string) FeatureOfInterest
	Get(EntityType, string, QueryOptions, EntityType, string) interface{}
	Init()
	Shutdown()
}

// Entities
type SensorThing interface {
	//	GetId() string
	//	GetSelfLink() string
	//	GetNavigationLink() string
	GetType() EntityType
}

type Thing interface {
	SensorThing
	GetDescription() string
	GetProperties() map[string]string

	GetLocations() []Location
	GetHistoricalLocations() []HistoricalLocation
	GetDatastreams() []Datastream
}

type LocationEncodingType string

const (
	LOCATION_ENCTYPE_GEOJSON LocationEncodingType = "application/vnd.geo+json"
)

type Location interface {
	SensorThing
	GetDescription() string
	GetEncodingType() LocationEncodingType
	GetLocationType() // !! depending on GetEncodingType()

	GetThings() []Thing
	GetHistoricalLocations() []HistoricalLocation
}

type HistoricalLocation interface {
	SensorThing
	GetTime() time.Time

	GetLocations() []Location
	GetThing() Thing
}

type DatastreamObservationType string

const (
	DATASTREAM_OBSTYPE_CATEGORY DatastreamObservationType = "http://www.opengis.net/def/observationType/ OGC-OM/2.0/OM_CategoryObservation"
	// IRI

	/*
		OM_CountObservation
		http://www.opengis.net/def/observationType/ OGC-OM/2.0/OM_CountObservation
		integer

		OM_Measurement
		http://www.opengis.net/def/observationType/ OGC-OM/2.0/OM_Measurement
		double

		OM_Observation
		http://www.opengis.net/def/observationType/ OGC-OM/2.0/OM_Observation
		Any

		OM_TruthObservation
		http://www.opengis.net/def/observationType/ OGC-OM/2.0/OM_TruthObservation
		boolean
	*/
)

type Datastream interface {
	SensorThing

	GetDescription() string
	GetUnitOfMeasurement() // UnitOfMeasure !!
	GetObservationType()   // !!
	GetObservedArea()      // !!
	GetPhenomenonTime() time.Time
	GetResultTime() time.Time

	GetThing() Thing
	GetSensor() Sensor
	GetObservedProperty() ObservedProperty
	GetObservations() []Observation
}

type SensorEncodingType string

const (
	SENSOR_ENCTYPE_PDF      SensorEncodingType = "application/pdf"
	SENSOR_ENCTYPE_SENSORML SensorEncodingType = "http://www.opengis.net/doc/IS/SensorML/2.0"
)

type Sensor interface {
	SensorThing
	GetDescription() string
	GetEncodingType() // !!
	GetMetadata()     // !!

	GetDatastreams() []Datastream
}

type ObservedProperty interface {
	SensorThing

	GetName() string
	GetDefinition() // !!
	GetDescription() string

	GetDatastreams() []Datastream
}

type Observation interface {
	SensorThing

	GetPhenomenonTime() time.Time
	GetResultTime() time.Time
	GetResult()        // !!
	GetResultQuality() // !!
	GetValidTime() time.Time
	GetParameters() map[string]string

	GetFeatureOfInterest() FeatureOfInterest
	GetDatastream() Datastream
}

type FeatureOfInterest interface {
	SensorThing

	GetDescription() string
	GetEncodingType() LocationEncodingType
	GetFeature() // !!

	GetObservations() []Observation
}

// Globals
var GLOB_ENV_HOST = "localhost:8000"
