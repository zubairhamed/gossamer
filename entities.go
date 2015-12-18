package gossamer

import (
	"time"
)

type SensorThingsEntity struct {
	Id       string `json:"@iot.id,omitempty" bson:"@iot_id"`
	SelfLink string `json:"@iot.selfLink,omitempty" bson:"-"`
}

func (e SensorThingsEntity) GetId() string {
	return e.Id
}

func (e SensorThingsEntity) GetSelfLink() string {
	return e.SelfLink
}

func NewThingEntity() *ThingEntity {
	t := &ThingEntity{}

	return t
}

func NewLocationEntity() *LocationEntity {
	e := &LocationEntity{}

	return e
}

func NewHistoricalLocationEntity() *HistoricalLocationEntity {
	e := &HistoricalLocationEntity{}

	return e
}

func NewDatastreamEntity() *DatastreamEntity {
	e := &DatastreamEntity{}

	return e
}

func NewSensorEntity() *SensorEntity {
	e := &SensorEntity{}

	return e
}

func NewObservedPropertyEntity() *ObservedPropertyEntity {
	e := &ObservedPropertyEntity{}

	return e
}

func NewObservationEntity() *ObservationEntity {
	e := &ObservationEntity{}

	return e
}

func NewFeatureOfInterestEntity() *FeatureOfInterestEntity {
	e := &FeatureOfInterestEntity{}

	return e
}

// The OGC SensorThings API follows the ITU-T definition, i.e., with regard to the Internet of Things,
// a thing is an object of the physical world (physical things) or the information world (virtual things)
// that is capable of being identified and integrated into communication networks [ITU-T Y.2060].
type ThingEntity struct {
	SensorThingsEntity         `bson:",inline"`
	NavLinkLocations           string               `json:"Locations@iot.navigationLink,omitempty" bson:"-"`
	NavLinkDatastreams         string               `json:"Datastreams@iot.navigationLink,omitempty" bson:"-"`
	NavLinkHistoricalLocations string               `json:"HistoricalLocations@iot.navigationLink,omitempty" bson:"-"`
	Description                string               `json:"description,omitempty"`
	Properties                 map[string]string    `json:"properties,omitempty"`
	IdLocations                []string             `json:"-" bson:"@iot_locations_id"`
	Locations                  []Location           `json:",omitempty" bson:"-"`
	HistoricalLocations        []HistoricalLocation `json:",omitempty" bson:"-"`
	Datastreams                []Datastream         `json:",omitempty" bson:"-"`
}

func (e ThingEntity) GetType() EntityType {
	return ENTITY_THINGS
}

func (e ThingEntity) GetAssociatedEntityId(ent EntityType) string {
	return ""
}

func (e ThingEntity) SetAssociatedEntityId(et EntityType, id string) {
	if et == ENTITY_LOCATIONS {
		entity := &LocationEntity{}
		entity.Id = id
		e.Locations = []Location{ entity }
	}

	if et == ENTITY_DATASTREAMS {
		entity := &DatastreamEntity{}
		entity.Id = id
		e.Datastreams = []Datastream{ entity }
	}

}

// The Location entity locates the Thing or the Things it associated with. A Thing’s Location entity is
// defined as the last known location of the Thing.
// A Thing’s Location may be identical to the Thing’s Observations’ FeatureOfInterest. In the context of the IoT,
// the principle location of interest is usually associated with the location of the Thing, especially for in-situ
// sensing applications. For example, the location of interest of a wifi-connected thermostat should be the building
// or the room in which the smart thermostat is located. And the FeatureOfInterest of the Observations made by the
// thermostat (e.g., room temperature readings) should also be the building or the room. In this case, the content
// of the smart thermostat’s location should be the same as the content of the temperature readings’ feature of interest.
type LocationEntity struct {
	SensorThingsEntity         `bson:",inline"`
	NavLinkHistoricalLocations string               `json:"HistoricalLocations@iot.navigationLink,omitempty" bson:"-"`
	NavLinkThings              string               `json:"Things@iot.navigationLink,omitempty" bson:"-"`
	Description                string               `json:"description,omitempty"`
	EncodingType               EncodingType         `json:"encodingType,omitempty" bson:"encodingType"`
	Location                   interface{}          `json:"location,omitempty" bson:"location"`
	Things                     []Thing              `json:",omitempty" bson:"-"`
	HistoricalLocations        []HistoricalLocation `json:",omitempty" bson:"-"`
}

func (e LocationEntity) GetDescription() string {
	return e.Description
}

func (e LocationEntity) GetEncodingType() LocationEncodingType {
	return LOCATION_ENCTYPE_GEOJSON
}

func (e LocationEntity) GetType() EntityType {
	return ENTITY_LOCATIONS
}

func (e LocationEntity) GetAssociatedEntityId(ent EntityType) string {
	return ""
}

func (e LocationEntity) SetAssociatedEntityId(et EntityType, id string) {

}

// A Thing’s HistoricalLocation entity set provides the current (i.e., last known) and previous locations of the
// Thing with their time.
type HistoricalLocationEntity struct {
	SensorThingsEntity         `bson:",inline"`
	NavLinkHistoricalLocations string `json:"HistoricalLocations@iot.navigationLink,omitempty"`
	NavLinkThing               string `json:"Thing@iot.navigationLink,omitempty"`
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

func (e HistoricalLocationEntity) SetAssociatedEntityId(et EntityType, id string) {

}

// A Datastream groups a collection of Observations and the Observations in a Datastream measure the
// same ObservedProperty and are produced by the same Sensor.
type DatastreamEntity struct {
	SensorThingsEntity      `bson:",inline"`
	NavLinkThing            string                  `json:"Thing@iot.navigationLink,omitempty" bson:"-"`
	NavLinkSensor           string                  `json:"Sensor@iot.navigationLink,omitempty" bson:"-"`
	NavLinkObservedProperty string                  `json:"ObservedProperty@iot.navigationLink,omitempty" bson:"-"`
	NavLinkObservations     string                  `json:"Observations@iot.navigationLink,omitempty" bson:"-"`
	PhenomenonTime          time.Time               `json:"phenomenonTime,omitempty"`
	ResultTime              time.Time               `json:"resultTime,omitempty"`
	UnitOfMeasurement       interface{}             `json:"unitOfMeasurement,omitempty" bson:"unitOfMeasurement"`
	ObservationType         ObservationType         `json:"observationType,omitempty" bson:"observationType"`
	Description             string                  `json:"description,omitempty"`
	IdThing                 string                  `json:"-" bson:"@iot_things_id"`
	IdObservedProperty      string                  `json:"-" bson:"@iot_observedproperties_id"`
	IdSensor                string                  `json:"-" bson:"@iot_sensors_id"`
	Thing                   *ThingEntity            `json:"Thing" bson:"-"`
	ObservedProperty        *ObservedPropertyEntity `json:"ObservedProperty" bson:"-"`
	Sensor                  *Sensor                 `json:"Sensor" bson:"-"`
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

func (e DatastreamEntity) SetAssociatedEntityId(et EntityType, id string) {

}

// A Sensor is an instrument that observes a property or phenomenon with the goal of producing an estimate of the
// value of the property.
type SensorEntity struct {
	SensorThingsEntity `bson:",inline"`
	NavLinkDatastreams string       `json:"Datastreams@iot.navigationLink,omitempty" bson:"-"`
	Description        string       `json:"description,omitempty"`
	EncodingType       EncodingType `json:"encodingType,omitempty"`
	Metadata           string       `json:"metadata,omitempty"`
}

func (e SensorEntity) GetType() EntityType {
	return ENTITY_SENSORS
}

func (e SensorEntity) GetAssociatedEntityId(ent EntityType) string {
	return ""
}

func (e SensorEntity) SetAssociatedEntityId(et EntityType, id string) {

}

// An ObservedProperty specifies the phenomenon of an Observation.
type ObservedPropertyEntity struct {
	SensorThingsEntity `bson:",inline"`
	NavLinkDatastreams string `json:"Datastreams@iot.navigationLink,omitempty" bson:"-"`
	Description        string `json:"description,omitempty"`
	Name               string `json:"name,omitempty"`
	Definition         string `json:"definition,omitempty"`
}

func (e ObservedPropertyEntity) GetType() EntityType {
	return ENTITY_OBSERVEDPROPERTIES
}

func (e ObservedPropertyEntity) GetAssociatedEntityId(ent EntityType) string {
	return ""
}

func (e ObservedPropertyEntity) SetAssociatedEntityId(et EntityType, id string) {

}

// An Observation is act of measuring or otherwise determining the value of a property
type ObservationEntity struct {
	SensorThingsEntity       `bson:",inline"`
	NavLinkFeatureOfInterest string                   `json:"FeatureOfInterest@iot.navigationLink,omitempty" bson:"-"`
	NavLinkDatastream        string                   `json:"Datastream@iot.navigationLink,omitempty" bson:"-"`
	PhenomenonTime           *TimePeriod              `json:"phenomenonTime,omitempty"`
	ResultTime               *TimeInstant             `json:"resultTime,omitempty"`
	Result                   interface{}              `json:"result,omitempty" bson:"result"`
	IdDatastream             string                   `json:"-" bson:"@iot_datastreams_id"`
	IdFeatureOfInterest      string                   `json:"-" bson:"@iot_featureofinterests_id"`
	Datastream               *DatastreamEntity        `json:"Datastream,omitempty" bson:"-"`
	FeatureOfInterest        *FeatureOfInterestEntity `json:"FeatureOfInterest,omitempty" bson:"-"`
}

func (e ObservationEntity) GetType() EntityType {
	return ENTITY_OBSERVATIONS
}

func (e ObservationEntity) GetAssociatedEntityId(ent EntityType) string {
	return ""
}

func (e *ObservationEntity) SetAssociatedEntityId(et EntityType, id string) {
	if et == ENTITY_DATASTREAMS {
		e.Datastream = NewDatastreamEntity()
		e.Datastream.Id = id
	}

	if et == ENTITY_FEATURESOFINTERESTS {
		e.FeatureOfInterest = NewFeatureOfInterestEntity()
		e.FeatureOfInterest.Id = id
	}
}

// An Observation results in a value being assigned to a phenomenon. The phenomenon is a property of a feature, the
// latter being the FeatureOfInterest of the Observation [OGC and ISO 19156:2001]. In the context of the Internet of
// Things, many Observations’ FeatureOfInterest can be the Location of the Thing. For example, the FeatureOfInterest
// of a wifi-connect thermostat can be the Location of the thermostat (i.e., the living room where the thermostat is
// located in). In the case of remote sensing, the FeatureOfInterest can be the geographical area or volume that is
// being sensed.
type FeatureOfInterestEntity struct {
	SensorThingsEntity  `bson:",inline"`
	NavLinkObservations string       `json:"Observations@iot.navigationLink,omitempty" bson:"-"`
	Description         string       `json:"description,omitempty"`
	EncodingType        EncodingType `json:"encodingType,omitempty"`
	Feature             interface{}  `json:"feature,omitempty" bson:"feature"`
}

func (e FeatureOfInterestEntity) GetType() EntityType {
	return ENTITY_FEATURESOFINTERESTS
}

func (e FeatureOfInterestEntity) GetAssociatedEntityId(ent EntityType) string {
	return ""
}

func (e FeatureOfInterestEntity) SetAssociatedEntityId(et EntityType, id string) {

}
