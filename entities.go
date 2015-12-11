package gossamer

import "time"

type EncodingType string

const (
	ENCODINGTYPE_PDF      EncodingType = "application/pdf"
	ENCODINGTYPE_SENSORML EncodingType = "http://www.opengis.net/doc/IS/SensorML/2.0"
)

type ObservationType string

type SensorThingsEntity struct {
	Id       string `bson:"@iot_id" json:"@iot.id"`
	SelfLink string `json:"@iot.selfLink"`
}

func (s *SensorThingsEntity) GetId() string {
	return s.Id
}

func (s *SensorThingsEntity) GetSelfLink() string {
	return s.SelfLink
}

type ThingEntity struct {
	SensorThingsEntity         `bson:",inline"`
	NavLinkLocations           string            `json:"Locations@iot.navigationLink"`
	NavLinkDatastreams         string            `json:"Datastreams@iot.navigationLink"`
	NavLinkHistoricalLocations string            `json:"HistoricalLocations@iot.navigationLink"`
	Description                string            `json:"description"`
	Properties                 map[string]string `json:"properties,omitempty"`
	LocationsId                []string          `json:"-"`
	HistoricalLocationsId      []string          `json:"-"`
	DatastreamsId              []string          `json:"-"`

	Locations           []Location           `json:",omitempty"`
	HistoricalLocations []HistoricalLocation `json:",omitempty"`
	Datastreams         []Datastream         `json:",omitempty"`
}

func (e ThingEntity) GetType() EntityType {
	return ENTITY_LOCATIONS
}

type LocationEntity struct {
	SensorThingsEntity         `bson:",inline"`
	NavLinkHistoricalLocations string       `json:"HistoricalLocations@iot.navigationLink"`
	NavLinkThings              string       `json:"Things@iot.navigationLink"`
	Description                string       `json:"description"`
	EncodingType               EncodingType `bson:"encodingType" json:"encodingType"`

	Things              []Thing              `json:",omitempty"`
	HistoricalLocations []HistoricalLocation `json:",omitempty"`
}

func (e LocationEntity) GetType() EntityType {
	return ENTITY_THINGS
}

type HistoricalLocationEntity struct {
	SensorThingsEntity         `bson:",inline"`
	NavLinkHistoricalLocations string `json:"HistoricalLocations@iot.navigationLink"`
	NavLinkThing               string `json:"Thing@iot.navigationLink"`
	Time                       time.Time
	EncodingType               EncodingType `json:"encodingType"`

	Thing     Thing      `json:",omitempty"`
	Locations []Location `json:",omitempty"`
}

func (e *HistoricalLocationEntity) GetType() EntityType {
	return ENTITY_HISTORICALLOCATIONS
}

type DatastreamEntity struct {
	SensorThingsEntity      `bson:",inline"`
	NavLinkThing            string    `json:"Thing@iot.navigationLink"`
	NavLinkSensor           string    `json:"Sensor@iot.navigationLink"`
	NavLinkObservedProperty string    `json:"ObservedProperty@iot.navigationLink"`
	NavLinkObservations     string    `json:"Observations@iot.navigationLink"`
	PhenomenonTime          time.Time `json:"phenomenonTime"`
	ResultTime              time.Time `json:"resultTime"`
	Description             string    `json:"description"`

	ObservedProperty ObservedProperty `json:",omitempty"`
	Sensor           Sensor           `json:",omitempty"`
	Thing            Thing            `json:",omitempty"`
	Observations     []Observation    `json:",omitempty"`
}

func (e *DatastreamEntity) GetType() EntityType {
	return ENTITY_DATASTREAMS
}

type SensorEntity struct {
	SensorThingsEntity `bson:",inline"`
	NavLinkDatastreams string       `json:"Datastreams@iot.navigationLink"`
	Description        string       `json:"description"`
	EncodingType       EncodingType `json:"encodingType"`
	Metadata           string       `json:"metadata"`

	Datastreams []Datastream `json:",omitempty"`
}

func (e *SensorEntity) GetType() EntityType {
	return ENTITY_SENSORS
}

type ObservedPropertyEntity struct {
	SensorThingsEntity `bson:",inline"`
	NavLinkDatastreams string `json:"Datastreams@iot.navigationLink"`
	Description        string `json:"description"`
	Name               string `json:"name"`
	Definition         string `json:"definition"`

	Datastreams []Datastream `json:",omitempty"`
}

func (e *ObservedPropertyEntity) GetType() EntityType {
	return ENTITY_OBSERVEDPROPERTIES
}

type ObservationEntity struct {
	SensorThingsEntity       `bson:",inline"`
	NavLinkFeatureOfInterest string    `json:"FeatureOfInterest@iot.navigationLink"`
	NavLinkDatastream        string    `json:"Datastream@iot.navigationLink"`
	PhenomenonTime           time.Time `json:"phenomenonTime"`
	ResultTime               time.Time `json:"resultTime"`

	Datastream        Datastream        `json:",omitempty"`
	FeatureOfInterest FeatureOfInterest `json:",omitempty"`
}

func (e *ObservationEntity) GetType() EntityType {
	return ENTITY_OBSERVATIONS
}

type FeatureOfInterestEntity struct {
	SensorThingsEntity  `bson:",inline"`
	NavLinkObservations string       `json:"Observations@iot.navigationLink"`
	Description         string       `json:"description"`
	EncodingType        EncodingType `json:"encodingType"`

	Observations []Observation `json:",omitempty"`
}

func (e *FeatureOfInterestEntity) GetType() EntityType {
	return ENTITY_FEATURESOFINTERESTS
}

type ValueList struct {
	Count int         `json:"count,omitempty"`
	Value interface{} `json:"value"`
}
