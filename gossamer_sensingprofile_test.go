// This tests against a Mongo persistence all the SensingProfile
// tasks
package gossamer

import (
	"testing"
	"github.com/zenazn/goji/web"
	"net/http"
	"net/url"
	"encoding/json"
	"bytes"
	"encoding/base64"
	"log"
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
	log.Println(m.buf)
	r := base64.NewDecoder(base64.StdEncoding, m.buf)
	decoder := json.NewDecoder(r)

	json.Unmarshal(m.buf.Bytes())

	val := make(map[string]interface{})
	decoder.Decode(val)

	return val
}

func TestCrudSensingProfile(t *testing.T) {
	var w *MockResponseWriter

	w = NewMockResponseWriter()

	server := &GossamerServer{}
	server.UseStore(NewMongoStore("localhost", "sensorthings"))

	// Assert zero-ed collections

	url, _ := url.Parse("http://localhost:8000/v1.0/Locations")

	req := &http.Request{
		Method: "GET",
		URL: url,
	}

	server.handleGet(web.C{}, w, req)
	log.Println(w.GetJSON())




//	Create Location
//	Create Sensor
//	Create ObservedProperty
//	Create FeatureOfInterest
//	Create Thing
//	Create Datastream (with xThing, xSensor, xObservedProperty)
//	Create Observation (with xDatastream, xFeatureOfInterest)

}
