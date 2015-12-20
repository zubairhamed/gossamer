package gossamer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strings"
)

func NewMongoStore(hosts string, db string) *MongoStore {
	ms := &MongoStore{
		hosts: hosts,
		db:    db,
	}
	ms.Init()

	return ms
}

type MongoStore struct {
	session *mgo.Session
	db      string
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

func (m *MongoStore) doQuery(query *mgo.Query, ent EntityType, findMultiple bool, opts QueryOptions) interface{} {
	if findMultiple {
		iter := query.Iter()
		switch ent {
		case ENTITY_THINGS:
			rs := []ThingEntity{}
			var r ThingEntity
			for iter.Next(&r) {
				m.postHandleThing(&r, opts)
				rs = append(rs, r)
			}
			return rs

		case ENTITY_OBSERVEDPROPERTIES:
			rs := []ObservedPropertyEntity{}
			var r ObservedPropertyEntity
			for iter.Next(&r) {
				m.postHandleObservedProperty(&r, opts)
				rs = append(rs, r)
			}
			return rs

		case ENTITY_LOCATIONS:
			rs := []LocationEntity{}
			var r LocationEntity
			for iter.Next(&r) {
				m.postHandleLocation(&r, opts)
				rs = append(rs, r)
			}
			return rs

		case ENTITY_DATASTREAMS:
			rs := []DatastreamEntity{}
			var r DatastreamEntity
			for iter.Next(&r) {
				m.postHandleDatastream(&r, opts)
				rs = append(rs, r)
			}
			return rs

		case ENTITY_SENSORS:
			rs := []SensorEntity{}
			var r SensorEntity
			for iter.Next(&r) {
				m.postHandleSensor(&r, opts)
				rs = append(rs, r)
			}
			return rs

		case ENTITY_OBSERVATIONS:
			rs := []ObservationEntity{}
			var r ObservationEntity
			for iter.Next(&r) {
				m.postHandleObservation(&r, opts)
				rs = append(rs, r)
			}
			return rs

		case ENTITY_FEATURESOFINTEREST:
			rs := []FeatureOfInterestEntity{}
			var r FeatureOfInterestEntity
			for iter.Next(&r) {
				m.postHandleFeatureOfInterest(&r, opts)
				rs = append(rs, r)
			}
			return rs

		case ENTITY_HISTORICALLOCATIONS:
			rs := []HistoricalLocationEntity{}
			var r HistoricalLocationEntity
			for iter.Next(&r) {
				m.postHandleHistoricalLocation(&r, opts)
				rs = append(rs, r)
			}
			return rs
		}
	} else {
		switch {
		case ent == ENTITY_THINGS || ent == ENTITY_THING:
			var r ThingEntity
			query.One(&r)
			m.postHandleThing(&r, opts)
			return r

		case ent == ENTITY_OBSERVEDPROPERTIES || ent == ENTITY_OBSERVEDPROPERTY:
			var r ObservedPropertyEntity
			query.One(&r)
			m.postHandleObservedProperty(&r, opts)
			return r

		case ent == ENTITY_LOCATIONS || ent == ENTITY_LOCATION:
			var r LocationEntity
			query.One(&r)
			m.postHandleLocation(&r, opts)
			return r

		case ent == ENTITY_DATASTREAMS || ent == ENTITY_DATASTREAM:
			var r DatastreamEntity
			query.One(&r)
			m.postHandleDatastream(&r, opts)
			return r

		case ent == ENTITY_SENSORS || ent == ENTITY_SENSOR:
			var r SensorEntity
			query.One(&r)
			m.postHandleSensor(&r, opts)
			return r

		case ent == ENTITY_OBSERVATIONS || ent == ENTITY_OBSERVATION:
			var r ObservationEntity
			query.One(&r)
			m.postHandleObservation(&r, opts)
			return r

		case ent == ENTITY_FEATURESOFINTEREST:
			var r FeatureOfInterestEntity
			query.One(&r)
			m.postHandleFeatureOfInterest(&r, opts)
			return r

		case ent == ENTITY_HISTORICALLOCATIONS || ent == ENTITY_HISTORICALLOCATION:
			var r HistoricalLocationEntity
			query.One(&r)
			m.postHandleHistoricalLocation(&r, opts)
			return r
		}
	}
	return nil
}

func (m *MongoStore) createQuery(c *mgo.Collection, rp ResourcePath, opts QueryOptions, lastResult interface{}) (query *mgo.Query, findMultiple bool) {
	last := rp.At(rp.CurrentIndex() - 1)
	curr := rp.At(rp.CurrentIndex())
	currEntity := curr.GetEntity()
	isFirst := rp.IsFirst()

	bsonMap := bson.M{}
	if curr.GetId() != "" {
		bsonMap["@iot_id"] = curr.GetId()
		findMultiple = false
	} else if IsSingularEntity(string(currEntity)) {
		bsonMap["@iot_id"] = GetAssociatedEntityId(lastResult.(SensorThing), currEntity)
		findMultiple = false
	} else {
		findMultiple = true
		if !isFirst {
			lastEntity := last.GetEntity()

			switch {
			case lastEntity == ENTITY_LOCATIONS && currEntity == ENTITY_THINGS,
				lastEntity == ENTITY_LOCATIONS && currEntity == ENTITY_HISTORICALLOCATIONS:

				bsonMap["@iot_locations_id"] = map[string]interface{}{
					"$in": []string{last.GetId()},
				}

			case lastEntity == ENTITY_HISTORICALLOCATIONS && currEntity == ENTITY_LOCATIONS:
				bsonMap["@iot_id"] = map[string]interface{}{
					"$in": lastResult.(HistoricalLocationEntity).IdLocations,
				}

			case lastEntity == ENTITY_THINGS && currEntity == ENTITY_LOCATIONS:
				bsonMap["@iot_id"] = map[string]interface{}{
					"$in": lastResult.(ThingEntity).IdLocations,
				}

			default:
				bsonMap["@iot_"+strings.ToLower(string(lastEntity))+"_id"] = last.GetId()
			}
		}
	}

	query = c.Find(bsonMap)

	// Filter

	// OrderBy
	if opts.OrderBySet() {
		opt := opts.GetOrderByOption()
		query.Sort(strings.Replace(strings.Join(opt.GetSortProperties(), ","), "@iot.", "@iot_", -1))
	}

	// Skip
	if opts.SkipSet() {
		skip := opts.GetSkipOption().GetValue()
		query.Skip(skip)
	}

	// Top
	if opts.TopSet() {
		top := opts.GetTopOption().GetValue()
		query.Limit(top)
	}

	// Expand
	if opts.ExpandSet() {
		// vals := opts.GetExpandOption().GetValue()
	}

	// Select
	if opts.SelectSet() {
		opt := opts.GetSelectOption().GetValue()
		selectBsonMap := bson.M{}

		if len(opt) > 0 {
			for _, v := range opt {
				selectBsonMap[v] = 1
			}
			query.Select(selectBsonMap)
		}
	}

	return
}

func (m *MongoStore) Query(rp ResourcePath, opts QueryOptions) (interface{}, error) {
	queryComplete := make(chan bool)
	var results interface{}
	go func() {
		session := m.cloneSession()
		defer session.Close()

		for rp.HasNext() {
			curr := rp.Next()
			currEntity := curr.GetEntity()

			c := session.DB(m.db).C(ResolveMongoCollectionName(currEntity))
			query, findMultiple := m.createQuery(c, rp, opts, results)

			resourceQueryComplete := make(chan bool)
			go func() {
				results = m.doQuery(query, currEntity, findMultiple, opts)
				resourceQueryComplete <- true
			}()
			<-resourceQueryComplete

			if rp.IsLast() {
				break
			}
		}
		queryComplete <- true
	}()
	<-queryComplete
	return results, nil
}

func (m *MongoStore) Insert(rp ResourcePath, payload SensorThing) error {
	queryComplete := make(chan bool)
	var results interface{}

	go func() {
		session := m.cloneSession()
		defer session.Close()

		opts, _ := CreateQueryOptions("")
		for rp.HasNext() {
			curr := rp.Next()
			currEntity := curr.GetEntity()
			c := session.DB(m.db).C(ResolveMongoCollectionName(currEntity))
			query, findMultiple := m.createQuery(c, rp, opts, results)

			resourceQueryComplete := make(chan bool)
			go func() {
				results = m.doQuery(query, currEntity, findMultiple, opts)
				resourceQueryComplete <- true
			}()
			<-resourceQueryComplete

			if rp.IsLast() {
				break
			}
		}

		// Generate IoT ID and Insert
		err := m.doInsert(payload, results)
		if err != nil {
			log.Println(err)
		}

		queryComplete <- true
	}()
	<-queryComplete

	return nil
}

func (m *MongoStore) Delete(ent EntityType, id string) (err error) {
	opComplete := make(chan bool)

	go func() {
		session := m.cloneSession()
		defer session.Close()

		c := session.DB(m.db).C(ResolveMongoCollectionName(ent))
		err = c.Remove(bson.M{"@iot_id": id})

		opComplete <- true
	}()
	<-opComplete

	return
}

func (m *MongoStore) doInsert(payload SensorThing, results interface{}) (err error) {
	var insertData interface{}

	switch payload.GetType() {
	case ENTITY_LOCATIONS:
		e := payload.(*LocationEntity)
		e.Id = GenerateEntityId()
		insertData = e

	case ENTITY_OBSERVATIONS:
		e := payload.(*ObservationEntity)
		e.Id = GenerateEntityId()
		if e.Datastream != nil {
			if e.Datastream.Id == "" {
				// TODO: Insert New DataStream in Datastream Collection
				log.Println("TODO: Insert New DataStream in Datastream Collection")
			}
			e.IdDatastream = e.Datastream.Id
		}

		if e.FeatureOfInterest != nil {
			if e.FeatureOfInterest.Id == "" {
				// TODO: Insert New FeatureOfInterest in FeatureOfInterest Collection
				log.Println("TODO: Insert New FeatureOfInterest in FeatureOfInterest Collection")
			}
			e.IdFeatureOfInterest = e.FeatureOfInterest.Id
		}
		insertData = e

	case ENTITY_THINGS:
		e := payload.(*ThingEntity)
		e.Id = GenerateEntityId()
		e.IdLocations = []string{}
		for _, v := range e.Locations {
			locId := v.Id
			if locId == "" {
				locId = GenerateEntityId()
				// TODO: Insert New Location
				log.Println("TODO: Insert New Location")
			}
			e.IdLocations = append(e.IdLocations, locId)
		}

		for _, v := range e.Datastreams {
			dsId := v.Id
			if dsId == "" {
				dsId = GenerateEntityId()
				// TODO: Insert New Datastream
				log.Println("TODO: Insert New Datastream")
			} else {
				// Unlink existing datastream.thing and relink to this one? Wut wut?
			}
		}

		insertData = e

	case ENTITY_HISTORICALLOCATIONS:
		// Shouldn't be allowed to insert

	case ENTITY_DATASTREAMS:
		e := payload.(*DatastreamEntity)
		e.Id = GenerateEntityId()

		if e.ObservedProperty != nil {
			if e.ObservedProperty.Id == "" {
				// TODO: Insert New ObservedProperty
				log.Println("TODO: Insert New ObservedProperty")
			}
			e.IdObservedProperty = e.ObservedProperty.Id
		}

		if e.Sensor != nil {
			if e.Sensor.Id == "" {
				// TODO: Insert New Sensor
				log.Println("TODO: Insert New Sensor")
			}
			e.IdSensor = e.Sensor.Id
		}

		if e.Thing != nil {
			if e.Thing.Id == "" {
				// TODO: Insert New Thing
				log.Println("TODO: Insert New Thing")
			}
			e.IdThing = e.Thing.Id
		}

		insertData = e

	case ENTITY_SENSORS:
		e := payload.(*SensorEntity)
		e.Id = GenerateEntityId()

		insertData = e

	case ENTITY_OBSERVEDPROPERTIES:
		e := payload.(*ObservedPropertyEntity)
		e.Id = GenerateEntityId()

		insertData = e

	case ENTITY_FEATURESOFINTEREST:
		e := payload.(*FeatureOfInterestEntity)
		e.Id = GenerateEntityId()

		insertData = e
	}

	session := m.cloneSession()
	defer session.Close()

	c := session.DB(m.db).C(ResolveMongoCollectionName(payload.GetType()))
	err = c.Insert(insertData)
	if err != nil {
		log.Println(err)
		return
	}
	return nil
}

func (m *MongoStore) postHandleThing(e *ThingEntity, opts QueryOptions) {

	if !opts.SelectSet() {
		e.SelfLink = ResolveSelfLinkUrl(e.Id, ENTITY_THINGS)

		e.NavLinkLocations = ResolveEntityLink(e.Id, ENTITY_THINGS) + "/Locations"
		e.NavLinkHistoricalLocations = ResolveEntityLink(e.Id, ENTITY_THINGS) + "/HistoricalLocations"
		e.NavLinkDatastreams = ResolveEntityLink(e.Id, ENTITY_THINGS) + "/Datastreams"
	}
}

func (m *MongoStore) postHandleObservedProperty(e *ObservedPropertyEntity, opts QueryOptions) {
	if !opts.SelectSet() {
		e.SelfLink = ResolveSelfLinkUrl(e.Id, ENTITY_OBSERVEDPROPERTIES)

		e.NavLinkDatastreams = ResolveEntityLink(e.Id, ENTITY_OBSERVEDPROPERTIES) + "/Datastreams"
	}
}

func (m *MongoStore) postHandleLocation(e *LocationEntity, opts QueryOptions) {
	if !opts.SelectSet() {
		e.SelfLink = ResolveSelfLinkUrl(e.Id, ENTITY_LOCATIONS)

		e.NavLinkHistoricalLocations = ResolveEntityLink(e.Id, ENTITY_LOCATIONS) + "/HistoricalLocations"
		e.NavLinkThings = ResolveEntityLink(e.Id, ENTITY_LOCATIONS) + "/Things"
	}
}

func (m *MongoStore) postHandleDatastream(e *DatastreamEntity, opts QueryOptions) {
	if !opts.SelectSet() {
		e.SelfLink = ResolveSelfLinkUrl(e.Id, ENTITY_DATASTREAMS)

		e.NavLinkObservations = ResolveEntityLink(e.Id, ENTITY_DATASTREAMS) + "/Observations"
		e.NavLinkObservedProperty = ResolveEntityLink(e.Id, ENTITY_DATASTREAMS) + "/ObservedProperty"
		e.NavLinkSensor = ResolveEntityLink(e.Id, ENTITY_DATASTREAMS) + "/Sensor"
		e.NavLinkThing = ResolveEntityLink(e.Id, ENTITY_DATASTREAMS) + "/Thing"
	}
}

func (m *MongoStore) postHandleSensor(e *SensorEntity, opts QueryOptions) {
	if !opts.SelectSet() {
		e.SelfLink = ResolveSelfLinkUrl(e.Id, ENTITY_SENSORS)

		e.NavLinkDatastreams = ResolveEntityLink(e.Id, ENTITY_SENSORS) + "/Datastreams"
	}
}

func (m *MongoStore) postHandleObservation(e *ObservationEntity, opts QueryOptions) {
	if !opts.SelectSet() {
		e.SelfLink = ResolveSelfLinkUrl(e.Id, ENTITY_OBSERVATIONS)

		e.NavLinkDatastream = ResolveEntityLink(e.Id, ENTITY_OBSERVATIONS) + "/Datastream"
		e.NavLinkFeatureOfInterest = ResolveEntityLink(e.Id, ENTITY_OBSERVATIONS) + "/FeatureOfInterest"
	}
}

func (m *MongoStore) postHandleFeatureOfInterest(e *FeatureOfInterestEntity, opts QueryOptions) {
	if !opts.SelectSet() {
		e.SelfLink = ResolveSelfLinkUrl(e.Id, ENTITY_FEATURESOFINTEREST)

		e.NavLinkObservations = ResolveEntityLink(e.Id, ENTITY_FEATURESOFINTEREST) + "/Observations"
	}
}

func (m *MongoStore) postHandleHistoricalLocation(e *HistoricalLocationEntity, opts QueryOptions) {
	if !opts.SelectSet() {
		e.SelfLink = ResolveSelfLinkUrl(e.Id, ENTITY_HISTORICALLOCATIONS)

		e.NavLinkHistoricalLocations = ResolveEntityLink(e.Id, ENTITY_HISTORICALLOCATIONS) + "/HistoricalLocations"
		e.NavLinkThing = ResolveEntityLink(e.Id, ENTITY_HISTORICALLOCATIONS) + "/Thing"
	}
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

	case ent == ENTITY_FEATURESOFINTEREST:
		return "featuresofinterest"

	case ent == ENTITY_HISTORICALLOCATION || ent == ENTITY_HISTORICALLOCATIONS:
		return "historicallocations"
	}
	return ""
}
