package gossamer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strings"
)

func NewMongoStore(hosts string) *MongoStore {
	ms := &MongoStore{
		hosts: hosts,
	}
	ms.Init()

	return ms
}

type MongoStore struct {
	session *mgo.Session
	hosts   string
}

func (m *MongoStore) Init() {
	session, err := mgo.Dial(m.hosts)
	if err != nil {
		log.Println(err)
	}
	session.SetMode(mgo.Monotonic, true)
	m.session = session
}

func (m *MongoStore) Shutdown() {
	m.session.Close()
}

func (m *MongoStore) cloneSession() *mgo.Session {
	return m.session.Clone()
}

func (m *MongoStore) Get(ent EntityType, entityId string, opts QueryOptions, lastEntity EntityType, lastEntityId string) interface{} {

	/*
		Apply Order
			# server-driven pagination
			$filter

			$count: returns number of records within collection with result, false do not return
			http://example.org/v1.0/Things?$count=true
			{
				"@iot.count": 2, "value": [
					{...},
					{...}
				]
			}

			$orderby: 	specifies order of collection returned
			http://example.org/v1.0/Observations?$orderby=result

			$skip: 	return entities starting with skip+1. if top and skip applied together, apply top first then skip
			http://example.org/v1.0/Things?$skip=5

			$top:	limits number of records returns by collection (up to but not greater)
			http://example.org/v1.0/Things?$top=5

			# post server-driven pagination
			$expand
			$select
	*/
	session := m.cloneSession()
	defer session.Close()

	c := session.DB("sensorthings").C(ResolveMongoCollectionName(ent))
	entityIdIsEmpty := entityId == ""
	lastEntityIdIsEmpty := lastEntityId == ""
	bsonMap := bson.M{}
	var results interface{}

	if !entityIdIsEmpty {
		bsonMap["@iot_id"] = entityId
	}

	if !lastEntityIdIsEmpty {
		bsonMap["@iot_"+strings.ToLower(string(lastEntity))+"_id"] = lastEntityId
	}
	query := c.Find(bsonMap)

	if !entityIdIsEmpty || IsSingularEntity(string(ent)) {
		switch {
		case ent == ENTITY_THINGS || ent == ENTITY_THING:
			var r ThingEntity
			query.One(&r)
			m.postHandleThing(&r)
			return r

		case ent == ENTITY_OBSERVEDPROPERTIES || ent == ENTITY_OBSERVEDPROPERTY:
			var r ObservedPropertyEntity
			query.One(&r)
			m.postHandleObservedProperty(&r)
			return r

		case ent == ENTITY_LOCATIONS || ent == ENTITY_LOCATION:
			var r LocationEntity
			query.One(&r)
			m.postHandleLocation(&r)
			return r

		case ent == ENTITY_DATASTREAMS || ent == ENTITY_DATASTREAM:
			var r DatastreamEntity
			query.One(&r)
			m.postHandleDatastream(&r)
			return r

		case ent == ENTITY_SENSORS || ent == ENTITY_SENSOR:
			var r SensorEntity
			query.One(&r)
			m.postHandleSensor(&r)
			return r

		case ent == ENTITY_OBSERVATIONS || ent == ENTITY_OBSERVATION:
			var r ObservationEntity
			query.One(&r)
			m.postHandleObservation(&r)
			return r

		case ent == ENTITY_FEATURESOFINTERESTS || ent == ENTITY_FEATURESOFINTEREST:
			var r FeatureOfInterestEntity
			query.One(&r)
			m.postHandleFeatureOfInterest(&r)
			return r

		case ent == ENTITY_HISTORICALLOCATIONS || ent == ENTITY_HISTORICALLOCATION:
			var r HistoricalLocationEntity
			query.One(&r)
			m.postHandleHistoricalLocation(&r)
			return r
		}
	} else { // Find all
		iter := query.Iter()

		switch ent {
		case ENTITY_THINGS:
			rs := []ThingEntity{}
			var r ThingEntity
			for iter.Next(&r) {
				m.postHandleThing(&r)
				rs = append(rs, r)
			}
			return rs

		case ENTITY_OBSERVEDPROPERTIES:
			rs := []ObservedPropertyEntity{}
			var r ObservedPropertyEntity
			for iter.Next(&r) {
				m.postHandleObservedProperty(&r)
				rs = append(rs, r)
			}
			return rs

		case ENTITY_LOCATIONS:
			rs := []LocationEntity{}
			var r LocationEntity
			for iter.Next(&r) {
				m.postHandleLocation(&r)
				rs = append(rs, r)
			}
			return rs

		case ENTITY_DATASTREAMS:
			rs := []DatastreamEntity{}
			var r DatastreamEntity
			for iter.Next(&r) {
				m.postHandleDatastream(&r)
				rs = append(rs, r)
			}
			return rs

		case ENTITY_SENSORS:
			rs := []SensorEntity{}
			var r SensorEntity
			for iter.Next(&r) {
				m.postHandleSensor(&r)
				rs = append(rs, r)
			}
			return rs

		case ENTITY_OBSERVATIONS:
			rs := []ObservationEntity{}
			var r ObservationEntity
			for iter.Next(&r) {
				m.postHandleObservation(&r)
				rs = append(rs, r)
			}
			return rs

		case ENTITY_FEATURESOFINTERESTS:
			rs := []FeatureOfInterestEntity{}
			var r FeatureOfInterestEntity
			for iter.Next(&r) {
				m.postHandleFeatureOfInterest(&r)
				rs = append(rs, r)
			}
			return rs

		case ENTITY_HISTORICALLOCATIONS:
			rs := []HistoricalLocationEntity{}
			var r HistoricalLocationEntity
			for iter.Next(&r) {
				m.postHandleHistoricalLocation(&r)
				rs = append(rs, r)
			}
			return rs
		}
	}

	log.Println("Retutning results", results)
	return results
}

func (m *MongoStore) postHandleThing(e *ThingEntity) {
	e.SelfLink = ResolveSelfLinkUrl(e.Id, ENTITY_THINGS)

	e.NavLinkLocations = ResolveEntityLink(e.Id, ENTITY_THINGS) + "/Locations"
	e.NavLinkHistoricalLocations = ResolveEntityLink(e.Id, ENTITY_THINGS) + "/HistoricalLocations"
	e.NavLinkDatastreams = ResolveEntityLink(e.Id, ENTITY_THINGS) + "/Datastreams"
}

func (m *MongoStore) postHandleObservedProperty(e *ObservedPropertyEntity) {
	e.SelfLink = ResolveSelfLinkUrl(e.Id, ENTITY_OBSERVEDPROPERTIES)

	e.NavLinkDatastreams = ResolveEntityLink(e.Id, ENTITY_OBSERVEDPROPERTIES) + "/Datastreams"
}

func (m *MongoStore) postHandleLocation(e *LocationEntity) {
	e.SelfLink = ResolveSelfLinkUrl(e.Id, ENTITY_LOCATIONS)

	e.NavLinkHistoricalLocations = ResolveEntityLink(e.Id, ENTITY_LOCATIONS) + "/HistoricalLocations"
	e.NavLinkThings = ResolveEntityLink(e.Id, ENTITY_LOCATIONS) + "/Things"
}

func (m *MongoStore) postHandleDatastream(e *DatastreamEntity) {
	e.SelfLink = ResolveSelfLinkUrl(e.Id, ENTITY_DATASTREAMS)

	e.NavLinkObservations = ResolveEntityLink(e.Id, ENTITY_DATASTREAMS) + "/Observations"
	e.NavLinkObservedProperty = ResolveEntityLink(e.Id, ENTITY_DATASTREAMS) + "/ObservedProperty"
	e.NavLinkSensor = ResolveEntityLink(e.Id, ENTITY_DATASTREAMS) + "/Sensor"
	e.NavLinkThing = ResolveEntityLink(e.Id, ENTITY_DATASTREAMS) + "/Thing"
}

func (m *MongoStore) postHandleSensor(e *SensorEntity) {
	e.SelfLink = ResolveSelfLinkUrl(e.Id, ENTITY_SENSORS)

	e.NavLinkDatastreams = ResolveEntityLink(e.Id, ENTITY_SENSORS) + "/Datastreams"
}

func (m *MongoStore) postHandleObservation(e *ObservationEntity) {
	e.SelfLink = ResolveSelfLinkUrl(e.Id, ENTITY_OBSERVATIONS)

	e.NavLinkDatastream = ResolveEntityLink(e.Id, ENTITY_OBSERVATIONS) + "/Datastream"
	e.NavLinkFeatureOfInterest = ResolveEntityLink(e.Id, ENTITY_OBSERVATIONS) + "/FeatureOfInterest"
}

func (m *MongoStore) postHandleFeatureOfInterest(e *FeatureOfInterestEntity) {
	e.SelfLink = ResolveSelfLinkUrl(e.Id, ENTITY_FEATURESOFINTERESTS)

	e.NavLinkObservations = ResolveEntityLink(e.Id, ENTITY_FEATURESOFINTERESTS) + "/Observations"
}

func (m *MongoStore) postHandleHistoricalLocation(e *HistoricalLocationEntity) {
	e.SelfLink = ResolveSelfLinkUrl(e.Id, ENTITY_HISTORICALLOCATIONS)

	e.NavLinkHistoricalLocations = ResolveEntityLink(e.Id, ENTITY_HISTORICALLOCATIONS) + "/HistoricalLocations"
	e.NavLinkThing = ResolveEntityLink(e.Id, ENTITY_HISTORICALLOCATIONS) + "/Thing"
}

func ResolveMongoCollectionName(ent EntityType) string {
	switch {
	case ent == ENTITY_THINGS || ent == ENTITY_THING:
		return "things"

	case ent == ENTITY_OBSERVEDPROPERTY || ent == ENTITY_OBSERVEDPROPERTIES:
		return "observedproperties"

	case ent == ENTITY_LOCATION || ent == ENTITY_LOCATIONS:
		return "locations"

	case ent == ENTITY_DATASTREAM || ent == ENTITY_DATASTREAMS:
		return "datastreams"

	case ent == ENTITY_SENSOR || ent == ENTITY_SENSORS:
		return "sensors"

	case ent == ENTITY_OBSERVATION || ent == ENTITY_OBSERVATIONS:
		return "observations"

	case ent == ENTITY_FEATURESOFINTEREST || ent == ENTITY_FEATURESOFINTERESTS:
		return "featureofinterests"

	case ent == ENTITY_HISTORICALLOCATION || ent == ENTITY_HISTORICALLOCATIONS:
		return "historicallocations"
	}
	return ""
}
