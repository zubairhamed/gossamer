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

type SensorThingsEntity struct {
	Id       string `bson:"@iot_id" json:"@iot.id"`
	SelfLink string `json:"@iot.selfLink"`
}

func (e SensorThingsEntity) GetId() string {
	return e.Id
}

func (e SensorThingsEntity) GetSelfLink() string {
	return e.SelfLink
}

type ThingEntity struct {
	SensorThingsEntity         `bson:",inline"`
	NavLinkLocations           string               `json:"Locations@iot.navigationLink"`
	NavLinkDatastreams         string               `json:"Datastreams@iot.navigationLink"`
	NavLinkHistoricalLocations string               `json:"HistoricalLocations@iot.navigationLink"`
	Description                string               `json:"description"`
	Properties                 map[string]string    `json:"properties,omitempty"`
	IdLocations                []string             `json:"-" bson:"@iot_locations_id"`
	Locations                  []Location           `json:",omitempty"`
	HistoricalLocations        []HistoricalLocation `json:",omitempty"`
	Datastreams                []Datastream         `json:",omitempty"`
}

func (e ThingEntity) GetType() EntityType {
	return ENTITY_THINGS
}

func (e ThingEntity) GetAssociatedEntityId(ent EntityType) string {
	return ""
}

type LocationEntity struct {
	SensorThingsEntity         `bson:",inline"`
	NavLinkHistoricalLocations string               `json:"HistoricalLocations@iot.navigationLink"`
	NavLinkThings              string               `json:"Things@iot.navigationLink"`
	Description                string               `json:"description"`
	EncodingType               EncodingType         `bson:"encodingType" json:"encodingType"`
	Things                     []Thing              `json:",omitempty"`
	HistoricalLocations        []HistoricalLocation `json:",omitempty"`
}

func (e LocationEntity) GetType() EntityType {
	return ENTITY_LOCATIONS
}

func (e LocationEntity) GetAssociatedEntityId(ent EntityType) string {
	return ""
}

type HistoricalLocationEntity struct {
	SensorThingsEntity         `bson:",inline"`
	NavLinkHistoricalLocations string `json:"HistoricalLocations@iot.navigationLink"`
	NavLinkThing               string `json:"Thing@iot.navigationLink"`
	Time                       time.Time
	EncodingType               EncodingType `json:"encodingType,omitempty"`
	IdThing                    string       `json:"-" bson:"@iot_things_id"`
	IdLocations                []string     `json:"-" bson:"@iot_locations_id"`
	Thing                      Thing        `json:",omitempty"`
	Locations                  []Location   `json:",omitempty"`
}

func (e HistoricalLocationEntity) GetType() EntityType {
	return ENTITY_HISTORICALLOCATIONS
}

func (e HistoricalLocationEntity) GetAssociatedEntityId(ent EntityType) string {
	return ""
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
	IdThing                 string    `json:"-" bson:"@iot_things_id"`
	IdObservedProperty      string    `json:"-" bson:"@iot_observedproperties_id"`
	IdSensor                string    `json:"-" bson:"@iot_sensors_id"`
}

func (e DatastreamEntity) GetType() EntityType {
	return ENTITY_DATASTREAMS
}

func (e DatastreamEntity) GetAssociatedEntityId(ent EntityType) string {
	switch {
	case ent == ENTITY_SENSOR || ent == ENTITY_SENSORS:
		return e.IdSensor

	case ent == ENTITY_THING || ent == ENTITY_THINGS:
		return e.IdThing

	case ent == ENTITY_OBSERVEDPROPERTIES || ent == ENTITY_OBSERVEDPROPERTY:
		return e.IdObservedProperty
	}
	return ""
}

type SensorEntity struct {
	SensorThingsEntity `bson:",inline"`
	NavLinkDatastreams string       `json:"Datastreams@iot.navigationLink"`
	Description        string       `json:"description"`
	EncodingType       EncodingType `json:"encodingType"`
	Metadata           string       `json:"metadata"`
}

func (e SensorEntity) GetType() EntityType {
	return ENTITY_SENSORS
}

func (e SensorEntity) GetAssociatedEntityId(ent EntityType) string {
	return ""
}

type ObservedPropertyEntity struct {
	SensorThingsEntity `bson:",inline"`
	NavLinkDatastreams string `json:"Datastreams@iot.navigationLink"`
	Description        string `json:"description"`
	Name               string `json:"name"`
	Definition         string `json:"definition"`
}

func (e ObservedPropertyEntity) GetType() EntityType {
	return ENTITY_OBSERVEDPROPERTIES
}

func (e ObservedPropertyEntity) GetAssociatedEntityId(ent EntityType) string {
	return ""
}

type ObservationEntity struct {
	SensorThingsEntity       `bson:",inline"`
	NavLinkFeatureOfInterest string    `json:"FeatureOfInterest@iot.navigationLink"`
	NavLinkDatastream        string    `json:"Datastream@iot.navigationLink"`
	PhenomenonTime           time.Time `json:"phenomenonTime"`
	ResultTime               time.Time `json:"resultTime"`
	IdDatastream             string    `json:"-" bson:"@iot_datastreams_id"`
	IdFeatureOfInterest      string    `json:"-" bson:"@iot_featureofinterests_id"`
}

func (e ObservationEntity) GetType() EntityType {
	return ENTITY_OBSERVATIONS
}

func (e ObservationEntity) GetAssociatedEntityId(ent EntityType) string {
	return ""
}

type FeatureOfInterestEntity struct {
	SensorThingsEntity  `bson:",inline"`
	NavLinkObservations string       `json:"Observations@iot.navigationLink"`
	Description         string       `json:"description"`
	EncodingType        EncodingType `json:"encodingType"`
}

func (e FeatureOfInterestEntity) GetType() EntityType {
	return ENTITY_FEATURESOFINTERESTS
}

func (e FeatureOfInterestEntity) GetAssociatedEntityId(ent EntityType) string {
	return ""
}
