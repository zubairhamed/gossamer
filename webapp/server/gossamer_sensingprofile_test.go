// This tests against a Mongo persistence all the SensingProfile
// tasks
package server_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/zenazn/goji/web"
	"github.com/zubairhamed/gossamer"
	"github.com/zubairhamed/gossamer/webapp/server"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"testing"
	"time"
	_ "time"
	"github.com/zubairhamed/gossamer/store"
)

func NewMockResponseWriter() *MockResponseWriter {
	return &MockResponseWriter{
		buf: bytes.NewBufferString(""),
	}
}

type MockResponseWriter struct {
	buf *bytes.Buffer
}

func (m *MockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (m *MockResponseWriter) Write(b []byte) (int, error) {
	m.buf.Write(b)
	return len(b), nil
}

func (m *MockResponseWriter) WriteHeader(h int) {

}

func (m *MockResponseWriter) GetJSON() map[string]interface{} {
	val := make(map[string]interface{})
	json.Unmarshal(m.buf.Bytes(), &val)

	return val
}

func TestCrudSensingProfile(t *testing.T) {

	var w *MockResponseWriter
	var ret map[string]interface{}
	var req *http.Request
	var l int
	c := web.C{}
	var entityTypes []string = []string{
		"/FeaturesOfInterest",
		"/Locations",
		"/Sensors",
		"/Observations",
		"/Datastreams",
		"/Things",
		"/ObservedProperties",
	}

	s := &server.GossamerServer{}
	s.UseStore(store.NewMongoStore("localhost", "sensorthings"))
	DropCollection()

	// ####### BASIC INSERT #######
	//	Create Location
	req, w = NewMockHttp("POST", "/Locations", NewDefaultLocation())
	s.HandlePost(c, w, req)
	req, w = NewMockHttp("GET", "/Locations", "")
	s.HandleGet(c, w, req)
	ret = w.GetJSON()
	assert.Equal(t, 1, len(ret["value"].([]interface{})))

	//	Create Sensor
	req, w = NewMockHttp("POST", "/Sensors", NewDefaultSensor())
	s.HandlePost(c, w, req)
	req, w = NewMockHttp("GET", "/Sensors", "")
	s.HandleGet(c, w, req)
	ret = w.GetJSON()
	assert.Equal(t, 1, len(ret["value"].([]interface{})))
	sensorId := GetMapProperty(0, "@iot.id", ret)

	//	Create ObservedProperty
	req, w = NewMockHttp("POST", "/ObservedProperties", NewDefaultObservedProperty())
	s.HandlePost(c, w, req)
	req, w = NewMockHttp("GET", "/ObservedProperties", "")
	s.HandleGet(c, w, req)
	ret = w.GetJSON()
	assert.Equal(t, 1, len(ret["value"].([]interface{})))
	observedPropertyId := GetMapProperty(0, "@iot.id", ret)

	//	Create FeatureOfInterest
	req, w = NewMockHttp("POST", "/FeaturesOfInterest", NewDefaultFeaturesOfInterest())
	s.HandlePost(c, w, req)
	req, w = NewMockHttp("GET", "/FeaturesOfInterest", "")
	s.HandleGet(c, w, req)
	ret = w.GetJSON()
	assert.Equal(t, 1, len(ret["value"].([]interface{})))
	featureOfInterestId := GetMapProperty(0, "@iot.id", ret)

	//	Create Thing
	req, w = NewMockHttp("POST", "/Things", NewDefaultThing())
	s.HandlePost(c, w, req)
	req, w = NewMockHttp("GET", "/Things", "")
	s.HandleGet(c, w, req)
	ret = w.GetJSON()
	assert.Equal(t, 1, len(ret["value"].([]interface{})))
	thingId := GetMapProperty(0, "@iot.id", ret)

	// Create Datastream
	var ds *gossamer.DatastreamEntity
	ds = NewDefaultDatastream()
	ds.Thing = &gossamer.ThingEntity{}
	ds.Thing.Id = thingId
	ds.Sensor = &gossamer.SensorEntity{}
	ds.Sensor.Id = sensorId
	ds.ObservedProperty = &gossamer.ObservedPropertyEntity{}
	ds.ObservedProperty.Id = observedPropertyId

	req, w = NewMockHttp("POST", "/Datastreams", ds)
	s.HandlePost(c, w, req)
	req, w = NewMockHttp("GET", "/Datastreams", "")
	s.HandleGet(c, w, req)
	ret = w.GetJSON()
	assert.Equal(t, 1, len(ret["value"].([]interface{})))
	datastreamId := GetMapProperty(0, "@iot.id", ret)

	// Create Observation
	var obs *gossamer.ObservationEntity
	obs = NewDefaultObservation()
	ds = &gossamer.DatastreamEntity{}
	ds.Id = datastreamId
	obs.Datastream = ds
	foi := &gossamer.FeatureOfInterestEntity{}
	foi.Id = featureOfInterestId
	obs.FeatureOfInterest = foi
	req, w = NewMockHttp("POST", "/Observations", obs)
	s.HandlePost(c, w, req)
	req, w = NewMockHttp("GET", "/Observations", "")
	s.HandleGet(c, w, req)
	ret = w.GetJSON()
	assert.Equal(t, 1, len(ret["value"].([]interface{})))

	// ####### ASSOCIATIVE INSERTS #######

	// ####### UPDATE #######

	// ####### UPDATE (PATCH) #######

	// ####### ADVANCED QUERIES #######
	for _, v := range entityTypes {
		u := v + "?$top=1"
		req, w = NewMockHttp("GET", u, "")
		s.HandleGet(c, w, req)
		ret = w.GetJSON()
		assert.NotNil(t, ret)
		assert.Equal(t, 1, len(ret["value"].([]interface{})))
		id := ret["value"].([]interface{})[0].(map[string]interface{})["@iot.id"]

		u = v + "(" + id.(string) + ")"
		req, w = NewMockHttp("GET", u, "")
		s.HandleGet(c, w, req)
		ret = w.GetJSON()
		assert.NotNil(t, ret)
	}

	// ####### DELETE #######
	for _, v := range entityTypes {
		req, w = NewMockHttp("GET", v, "")
		s.HandleGet(c, w, req)
		ret = w.GetJSON()
		l = len(ret["value"].([]interface{}))

		i := 0
		for i < l {
			id := GetMapProperty(i, "@iot.id", ret)
			req, w = NewMockHttp("DELETE", v+"("+id+")", "")
			s.HandleDelete(c, w, req)
			i++
		}
	}

	// ####### CHECK ZERO-ED COLLECTIONS #######
	for _, v := range entityTypes {
		req, w = NewMockHttp("GET", v, "")
		s.HandleGet(c, w, req)
		ret = w.GetJSON()
		l = len(ret["value"].([]interface{}))
		assert.Equal(t, 0, l)
	}
}

func GetMapProperty(idx int, prop string, val map[string]interface{}) string {
	arr := val["value"].([]interface{})
	ent := arr[idx].(map[string]interface{})

	return ent[prop].(string)
}

func DropCollection() {
	log.Println("Drop Collection")
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Println(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	err = session.DB("sensorthings").DropDatabase()
	if err != nil {
		log.Println(err)
	}
}

func NewMockHttp(method string, u string, p interface{}) (*http.Request, *MockResponseWriter) {
	b, _ := json.Marshal(p)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(method, "http://localhost:8000/v1.0"+u, body)

	return req, NewMockResponseWriter()
}

func NewDefaultLocation() *gossamer.LocationEntity {
	e := gossamer.NewLocationEntity()
	e.Description = "Description for Test Location"
	e.EncodingType = gossamer.LOCATION_ENCTYPE_GEOJSON
	e.Location = map[string]interface{}{
		"type":        "Point",
		"coordinates": []interface{}{-117.123, 54.123},
	}
	return e
}

func NewDefaultSensor() *gossamer.SensorEntity {
	e := gossamer.NewSensorEntity()
	e.Description = "Description for Sensor"
	e.EncodingType = gossamer.SENSOR_ENCTYPE_PDF
	e.Metadata = "Calibration date:  Jan 1, 2014"
	return e
}

func NewDefaultObservedProperty() *gossamer.ObservedPropertyEntity {
	e := gossamer.NewObservedPropertyEntity()
	e.Name = "Name Observed Property"
	e.Description = "Description for ObservedProperty"
	e.Definition = "Calibration date:  Jan 1, 2014"
	return e
}

func NewDefaultFeaturesOfInterest() *gossamer.FeatureOfInterestEntity {
	e := gossamer.NewFeatureOfInterestEntity()
	e.Description = "Description for Features of Interest"
	e.EncodingType = gossamer.LOCATION_ENCTYPE_GEOJSON
	e.Feature = "FEATURE"
	return e
	//		"feature": {
	//			"coordinates": [51.08386,-114.13036],
	//			"type": "Point"
	//		}
}

func NewDefaultThing() *gossamer.ThingEntity {
	e := gossamer.NewThingEntity()
	e.Description = "Description for Thing Entity"
	e.Properties = map[string]string{
		"property1": "value1",
		"property2": "value2",
		"property3": "value3",
	}
	return e
}

func NewDefaultDatastream() *gossamer.DatastreamEntity {
	e := gossamer.NewDatastreamEntity()

	e.UnitOfMeasurement = "UOM"
	e.ObservationType = gossamer.DATASTREAM_OBSTYPE_OBSERVATION
	e.Description = "Description for Datastream"

	return e
	//		"unitOfMeasurement": {
	//			"symbol": "%",
	//			"name": "Percentage",
	//			"definition": "http://www.qudt.org/qudt/owl/1.0.0/unit/Instances.html"
	//		},
	//		"Thing": {"@iot.id": 5394817},
	//		"ObservedProperty": {"@iot.id": 5394816},
	//		"Sensor": {"@iot.id": 5394815}
}

func NewDefaultObservation() *gossamer.ObservationEntity {
	e := gossamer.NewObservationEntity()
	e.PhenomenonTime = gossamer.NewTimePeriod(time.Now(), time.Now())
	e.ResultTime = gossamer.NewTimeInstant(time.Now())
	e.Result = 123

	return e
	//		"Datastream":{"@iot.id":100}
}
