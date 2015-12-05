package gossamer

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
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
	hosts 	string
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

	c := m.session.DB("sensorthings").C(string(ent))
	entityIdIsEmpty := entityId == ""
	lastEntityIdIsEmpty := lastEntityId == ""
	bsonMap := bson.M{}
	var results interface{}

	if !entityIdIsEmpty {
		bsonMap["@iot_id"] = entityId
	}

	if !lastEntityIdIsEmpty {
		bsonMap["@iot_" + string(lastEntity) + "_id"] = lastEntityId
	}
	query := c.Find(bsonMap)

	// Find one
	if !entityIdIsEmpty {
		query.One(&results)
	} else {	// Find all
		var r []interface{}
		query.All(&r)
		results = r
	}

	return results
}
