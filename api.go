package gossamer

import (
	"time"
)

type EncodingType string

const (
	ENCODINGTYPE_PDF      EncodingType = "application/pdf"
	ENCODINGTYPE_SENSORML EncodingType = "http://www.opengis.net/doc/IS/SensorML/2.0"
)

type ObservationType string

type ProtocolType int

const (
	HTTP ProtocolType = 0
	COAP ProtocolType = 1
)

type ResourcePath interface {
	All() []ResourcePathItem
	Current() ResourcePathItem
	Next() ResourcePathItem
	Prev() ResourcePathItem
	First() ResourcePathItem
	Last() ResourcePathItem

	// Get the containing resource for the target resource
	// e.g. returns
	Containing() ResourcePathItem

	IsLast() bool
	IsFirst() bool
	HasNext() bool

	CurrentIndex() int
	Size() int
	Add(ResourcePathItem)
	At(int) ResourcePathItem
}

type ResourcePathItem interface {
	GetEntity() EntityType
	GetId() string
	GetQueryOptions() QueryOptions
}

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

// Determines a list of Query Options that has been set for a request
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

// Query Option
type QueryOption interface {
	GetType() QueryOptionType
}

// The $expand system query option indicates the related entities to be represented inline. The value of the $expand
// query option must be a comma separated list of navigation property names. Additionally each navigation property
// can be followed by a forward slash and another navigation property to enable identifying a multi-level relationship.
type ExpandOption interface {
	QueryOption
	GetValue() []string
}

// The $select system query option requests that the service to return only the properties explicitly requested by
// the client. The value of a $select query option is a comma-separated list of selection clauses. Each selection
// clause may be a property name (including navigation property names). The service returns the specified content, if available, along with any available expanded navigation properties.
type SelectOption interface {
	QueryOption
	GetValue() []string
}

// The $orderby system query option specifies the order in which items are returned from the service.
type OrderByOption interface {
	QueryOption
	GetValue() []OrderByOptionValue
	GetSortProperties() []string
}

// asc, desc
type OrderByOptionValue interface {
	GetSortProperty() string
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
	UseStore(Datastore)
}

type Request interface {
	GetProtocol() ProtocolType
	GetQueryOptions() QueryOptions
	GetResourcePath() ResourcePath
}

type Datastore interface {
	Query(ResourcePath, QueryOptions) (interface{}, error)
	Insert(ResourcePath, SensorThing) error
	Delete(EntityType, string) error
	Init()
	Shutdown()
}

// Entities
type SensorThing interface {
	GetId() string
	GetSelfLink() string
	GetType() EntityType
}

// The OGC SensorThings API follows the ITU-T definition, i.e., with regard to the Internet of Things,
// a thing is an object of the physical world (physical things) or the information world (virtual things)
// that is capable of being identified and integrated into communication networks [ITU-T Y.2060].
type Thing interface {
	SensorThing

	// This is a short description of the corresponding Thing entity.
	GetDescription() string

	// An Object containing user-annotated properties as key-value pairs.
	GetProperties() map[string]string
}

type LocationEncodingType string

const (
	LOCATION_ENCTYPE_GEOJSON LocationEncodingType = "application/vnd.geo+json"
)

// The Location entity locates the Thing or the Things it associated with. A Thing’s Location entity is
// defined as the last known location of the Thing.
// A Thing’s Location may be identical to the Thing’s Observations’ FeatureOfInterest. In the context of the IoT,
// the principle location of interest is usually associated with the location of the Thing, especially for in-situ
// sensing applications. For example, the location of interest of a wifi-connected thermostat should be the building
// or the room in which the smart thermostat is located. And the FeatureOfInterest of the Observations made by the
// thermostat (e.g., room temperature readings) should also be the building or the room. In this case, the content
// of the smart thermostat’s location should be the same as the content of the temperature readings’ feature of interest.
type Location interface {
	SensorThing
	GetDescription() string
	GetEncodingType() LocationEncodingType
	// GetLocationType() // !! depending on GetEncodingType()

	// GetThings() []Thing
	// GetHistoricalLocations() []HistoricalLocation
}

// A Thing’s HistoricalLocation entity set provides the current (i.e., last known) and previous locations of the
// Thing with their time.
type HistoricalLocation interface {
	SensorThing

	// The time when the Thing is known at the Location.
	GetTime() time.Time

	GetLocations() []Location
	GetThing() Thing
}

type DatastreamObservationType string

const (
	DATASTREAM_OBSTYPE_CATEGORY         DatastreamObservationType = "http://www.opengis.net/def/observationType/ OGC-OM/2.0/OM_CategoryObservation"
	DATASTREAM_OBSTYPE_COUNT            DatastreamObservationType = "http://www.opengis.net/def/observationType/ OGC-OM/2.0/OM_CountObservation"
	DATASTREAM_OBSTYPE_MEASUREMENT      DatastreamObservationType = "http://www.opengis.net/def/observationType/ OGC-OM/2.0/OM_Measurement"
	DATASTREAM_OBSTYPE_OBSERVATION      DatastreamObservationType = "http://www.opengis.net/def/observationType/ OGC-OM/2.0/OM_Observation"
	DATASTREAM_OBSTYPE_TRUTHOBSERVATION DatastreamObservationType = "http://www.opengis.net/def/observationType/ OGC-OM/2.0/OM_TruthObservation"
)

// A Datastream groups a collection of Observations and the Observations in a Datastream measure the
// same ObservedProperty and are produced by the same Sensor.
type Datastream interface {
	SensorThing

	// GetDescription() string

	// A JSON Object containing three key- value pairs. The name property presents the full name of the
	// unitOfMeasurement; the symbol property shows the textual form of the unit symbol; and the definition
	// contains the IRI defining the unitOfMeasurement.
	// The values of these properties SHOULD follow the Unified Code for Unit of Measure (UCUM).
	// GetUnitOfMeasurement() // UnitOfMeasure !!

	// The type of Observation (with unique result type), which is used by the service to encode observations.
	//	GetObservationType() // !!

	// The spatial bounding box of the spatial extent of all FeaturesOfInterest that belong to the
	// Observations associated with this Datastream.
	//	GetObservedArea() // !!

	// The temporal bounding box of the phenomenon times of all observations belonging to this Datastream.
	// GetPhenomenonTime() time.Time

	// The temporal bounding box of the result times of all observations belonging to this Datastream.
	// GetResultTime() time.Time

	// GetThing() Thing
	// GetSensor() Sensor
	// GetObservedProperty() ObservedProperty
	// GetObservations() []Observation
}

type SensorEncodingType string

const (
	SENSOR_ENCTYPE_PDF      SensorEncodingType = "application/pdf"
	SENSOR_ENCTYPE_SENSORML SensorEncodingType = "http://www.opengis.net/doc/IS/SensorML/2.0"
)

// A Sensor is an instrument that observes a property or phenomenon with the goal of producing an estimate of the
// value of the property.
type Sensor interface {
	SensorThing

	GetDescription() string
	GetEncodingType() SensorEncodingType
	GetMetadata() interface{}
}

// An ObservedProperty specifies the phenomenon of an Observation.
type ObservedProperty interface {
	SensorThing

	GetName() string
	GetDefinition() string
	GetDescription() string
}

// An Observation is act of measuring or otherwise determining the value of a property
type Observation interface {
	SensorThing

	// The time instant or period of when the Observation happens.
	GetPhenomenonTime() *TimePeriod

	// The time of the Observation's result was generated.
	GetResultTime() *TimeInstant

	// The estimated value of an ObservedProperty from the Observation.
	GetResult() interface{}

	// Describes the quality of the result.
	GetResultQuality() interface{}

	// The time period during which the result may be used.
	GetValidTime() *TimePeriod

	// Key-value pairs showing the environmental conditions during measurement.
	GetParameters() map[string]interface{}
}

// An Observation results in a value being assigned to a phenomenon. The phenomenon is a property of a feature, the
// latter being the FeatureOfInterest of the Observation [OGC and ISO 19156:2001]. In the context of the Internet of
// Things, many Observations’ FeatureOfInterest can be the Location of the Thing. For example, the FeatureOfInterest
// of a wifi-connect thermostat can be the Location of the thermostat (i.e., the living room where the thermostat is
// located in). In the case of remote sensing, the FeatureOfInterest can be the geographical area or volume that is
// being sensed.
type FeatureOfInterest interface {
	SensorThing

	GetDescription() string
	GetEncodingType() LocationEncodingType
	GetFeature() interface{}
}

// Client API
type Client interface {
	QueryAll(EntityType, QueryOptions) ([]SensorThing, error)
	QueryOne(EntityType, QueryOptions) (SensorThing, error)

	InsertObservation(Observation) error
	InsertThing(Thing) error
	InsertObservedProperty(ObservedProperty) error
	InsertLocation(Location) error
	InsertDatastream(Datastream) error
	InsertSensor(Sensor) error
	InsertFeaturesOfInterest(FeatureOfInterest) error

	DeleteObservation(string) error
	DeleteThing(string) error
	DeleteObservedProperty(string) error
	DeleteLocation(string) error
	DeleteDatastream(string) error
	DeleteSensor(string) error
	DeleteFeaturesOfInterest(string) error

	UpdateObservation(Observation) error
	UpdateThing(Thing) error
	UpdateObservedProperty(ObservedProperty) error
	UpdateLocation(Location) error
	UpdateDatastream(Datastream) error
	UpdateSensor(Sensor) error
	UpdateFeaturesOfInterest(FeatureOfInterest) error

	PatchObservation(Observation) error
	PatchThing(Thing) error
	PatchObservedProperty(ObservedProperty) error
	PatchLocation(Location) error
	PatchDatastream(Datastream) error
	PatchSensor(Sensor) error
	PatchFeaturesOfInterest(FeatureOfInterest) error

	FindObservation(string) ([]Observation, error)
	FindThing(string) ([]Thing, error)
	FindObservedProperty(string) ([]ObservedProperty, error)
	FindLocation(string) ([]Location, error)
	FindDatastream(string) ([]Datastream, error)
	FindSensor(string) ([]Sensor, error)
	FindFeaturesOfInterest(string) ([]FeatureOfInterest, error)
}

type ClientQuery interface {
	All() ([]SensorThing, error)
	One() (SensorThing, error)
	Filter() ClientQuery
	Count(bool) ClientQuery
	OrderBy(...string) ClientQuery
	Skip(int) ClientQuery
	Top(int) ClientQuery
	Expand(...string) ClientQuery
	Select(...string) ClientQuery
	GetUrlString() string
}

type ClientDelete interface {
}

type ClientPatch interface {
}

type ClientUpdate interface {
}

type ClientInsert interface {
}

type EntityClient interface {
	GetType() EntityType
	GetId() string
	Query() ClientQuery
	Delete() ClientDelete
	Patch() ClientPatch
	Update() ClientUpdate
	Insert() ClientInsert
}
