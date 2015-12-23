package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	. "github.com/zubairhamed/gossamer"
	"io/ioutil"
	"net/http"
)

func NewClient(url string) Client {
	return &GossamerClient{
		url: url,
	}
}

type GossamerClient struct {
	url string
}

func (c *GossamerClient) doInsert(o interface{}, pathFragment string) error {
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}

	r := bytes.NewReader(b)

	resp, err := http.Post(c.url+"/v1.0"+pathFragment, "application/json", r)

	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		contents, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(contents))
	}

	return nil
}

func (c *GossamerClient) InsertObservation(o Observation) error {
	return c.doInsert(o, "/Observations")
}

func (c *GossamerClient) InsertThing(o Thing) error {
	return c.doInsert(o, "/Things")
}

func (c *GossamerClient) InsertObservedProperty(o ObservedProperty) error {
	return c.doInsert(o, "/ObservedProperties")
}

func (c *GossamerClient) InsertLocation(o Location) error {
	return c.doInsert(o, "/Locations")
}

func (c *GossamerClient) InsertDatastream(o Datastream) error {
	return c.doInsert(o, "/Datastreams")
}
func (c *GossamerClient) InsertSensor(o Sensor) error {
	return c.doInsert(o, "/Sensors")
}

func (c *GossamerClient) InsertFeaturesOfInterest(o FeatureOfInterest) error {
	return c.doInsert(o, "/FeaturesOfInterest")
}

func (c *GossamerClient) doDelete(pathFragment string, id string) error {
	u := c.url + "/v1.0" + pathFragment + "(" + id + ")"
	req, err := http.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		contents, _ := ioutil.ReadAll(resp.Body)

		return errors.New(string(contents))
	}
	return nil
}

func (c *GossamerClient) DeleteObservation(id string) error {
	return c.doDelete("/Observations", id)
}

func (c *GossamerClient) DeleteThing(id string) error {
	return c.doDelete("/Things", id)
}

func (c *GossamerClient) DeleteObservedProperty(id string) error {
	return c.doDelete("/ObservedProperties", id)
}

func (c *GossamerClient) DeleteLocation(id string) error {
	return c.doDelete("/Locations", id)
}

func (c *GossamerClient) DeleteDatastream(id string) error {
	return c.doDelete("/Datastreams", id)
}

func (c *GossamerClient) DeleteSensor(id string) error {
	return c.doDelete("/Sensors", id)
}

func (c *GossamerClient) DeleteFeaturesOfInterest(id string) error {
	return c.doDelete("/FeaturesOfInterest", id)
}

/*
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}

	r := bytes.NewReader(b)

	resp, err := http.Post(c.url+"/v1.0"+pathFragment, "application/json", r)

	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		contents, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(contents))
	}

	return nil
*/

func (c *GossamerClient) doUpdate(o interface{}, pathFragment string) error {
	st := o.(SensorThing)
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}

	r := bytes.NewReader(b)

	u := c.url + "/v1.0" + pathFragment + "(" + st.GetId() + ")"
	req, err := http.NewRequest("PUT", u, r)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		contents, _ := ioutil.ReadAll(resp.Body)

		return errors.New(string(contents))
	}
	return nil
}

func (c *GossamerClient) UpdateObservation(o Observation) error {
	return c.doUpdate(o, "/Observations")
}
func (c *GossamerClient) UpdateThing(o Thing) error {
	return c.doUpdate(o, "/Observations")
}

func (c *GossamerClient) UpdateObservedProperty(o ObservedProperty) error {
	return c.doUpdate(o, "/ObservedProperties")
}

func (c *GossamerClient) UpdateLocation(o Location) error {
	return c.doUpdate(o, "/Locations")
}

func (c *GossamerClient) UpdateDatastream(o Datastream) error {
	return c.doUpdate(o, "/Datastreams")
}

func (c *GossamerClient) UpdateSensor(o Sensor) error {
	return c.doUpdate(o, "/Sensors")
}

func (c *GossamerClient) UpdateFeaturesOfInterest(o FeatureOfInterest) error {
	return c.doUpdate(o, "/FeaturesOfInterest")
}

func (c *GossamerClient) doPatch(o interface{}, pathFragment string) error {
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}

	r := bytes.NewReader(b)

	u := c.url + "/v1.0" + pathFragment
	req, err := http.NewRequest("PATCH", u, r)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		contents, _ := ioutil.ReadAll(resp.Body)

		return errors.New(string(contents))
	}
	return nil
}

func (c *GossamerClient) PatchObservation(o Observation) error {
	return c.doPatch(o, "/Observations")
}

func (c *GossamerClient) PatchThing(o Thing) error {
	return c.doPatch(o, "/Things")
}

func (c *GossamerClient) PatchObservedProperty(o ObservedProperty) error {
	return c.doPatch(o, "/ObservedProperties")
}

func (c *GossamerClient) PatchLocation(o Location) error {
	return c.doPatch(o, "/Locations")
}

func (c *GossamerClient) PatchDatastream(o Datastream) error {
	return c.doPatch(o, "/Datastreams")
}

func (c *GossamerClient) PatchSensor(o Sensor) error {
	return c.doPatch(o, "/Sensors")
}

func (c *GossamerClient) PatchFeaturesOfInterest(o FeatureOfInterest) error {
	return c.doPatch(o, "/FeaturesOfInterest")
}

func (c *GossamerClient) doGet(pathFragment string, id string) ([]byte, error) {
	resp, err := http.Get(c.url + "/v1.0" + pathFragment + "(" + id + ")")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		contents, _ := ioutil.ReadAll(resp.Body)

		return nil, errors.New(string(contents))
	} else {
		contents, _ := ioutil.ReadAll(resp.Body)

		return contents, nil
	}
}

func (c *GossamerClient) GetObservation(id string) (Observation, error) {
	b, err := c.doGet("/Observations", id)

	var o ObservationEntity
	json.Unmarshal(b, &o)

	return o, err
}

func (c *GossamerClient) GetThing(id string) (Thing, error) {
	b, err := c.doGet("/Things", id)

	var o ThingEntity
	json.Unmarshal(b, &o)

	return o, err
}

func (c *GossamerClient) GetObservedProperty(id string) (ObservedProperty, error) {
	b, err := c.doGet("/ObservedProperties", id)

	var o ObservedPropertyEntity
	json.Unmarshal(b, &o)

	return o, err

}

func (c *GossamerClient) GetLocation(id string) (Location, error) {
	b, err := c.doGet("/Locations", id)

	var o LocationEntity
	json.Unmarshal(b, &o)

	return o, err

}

func (c *GossamerClient) GetDatastream(id string) (Datastream, error) {
	b, err := c.doGet("/Datastreams", id)

	var o DatastreamEntity
	json.Unmarshal(b, &o)

	return o, err

}

func (c *GossamerClient) GetSensor(id string) (Sensor, error) {
	b, err := c.doGet("/Sensors", id)

	var o SensorEntity
	json.Unmarshal(b, &o)

	return o, err

}

func (c *GossamerClient) GetFeaturesOfInterest(id string) (FeatureOfInterest, error) {
	b, err := c.doGet("/FeaturesOfInterest", id)

	var o FeatureOfInterestEntity
	json.Unmarshal(b, &o)

	return o, err
}

func (c *GossamerClient) GetHistoricalLocation(id string) (HistoricalLocation, error) {
	b, err := c.doGet("/HistoricalLocations", id)

	var o HistoricalLocationEntity
	json.Unmarshal(b, &o)

	return o, err
}

func (c *GossamerClient) doFind(pathFragment string) ([]byte, error) {
	resp, err := http.Get(c.url + "/v1.0" + pathFragment)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		contents, _ := ioutil.ReadAll(resp.Body)

		return nil, errors.New(string(contents))
	} else {
		contents, _ := ioutil.ReadAll(resp.Body)

		return contents, nil
	}
}

func (c *GossamerClient) FindObservations() ([]Observation, error) {
	b, err := c.doFind("/Observations")

	var valueList struct {
		Count int
		Value []ObservationEntity
	}

	json.Unmarshal(b, &valueList)

	ret := make([]Observation, len(valueList.Value))
	for i, v := range valueList.Value {
		ret[i] = v
	}

	return ret, err
}

func (c *GossamerClient) FindThings() ([]Thing, error) {
	b, err := c.doFind("/Things")

	var valueList struct {
		Count int
		Value []ThingEntity
	}

	json.Unmarshal(b, &valueList)

	ret := make([]Thing, len(valueList.Value))
	for i, v := range valueList.Value {
		ret[i] = v
	}

	return ret, err
}

func (c *GossamerClient) FindObservedProperties() ([]ObservedProperty, error) {
	b, err := c.doFind("/ObservedProperties")

	var valueList struct {
		Count int
		Value []ObservedPropertyEntity
	}

	json.Unmarshal(b, &valueList)

	ret := make([]ObservedProperty, len(valueList.Value))
	for i, v := range valueList.Value {
		ret[i] = v
	}

	return ret, err
}

func (c *GossamerClient) FindLocations() ([]Location, error) {
	b, err := c.doFind("/Locations")

	var valueList struct {
		Count int
		Value []LocationEntity
	}

	json.Unmarshal(b, &valueList)

	ret := make([]Location, len(valueList.Value))
	for i, v := range valueList.Value {
		ret[i] = v
	}

	return ret, err
}

func (c *GossamerClient) FindHistoricalLocations() ([]HistoricalLocation, error) {
	b, err := c.doFind("/HistoricalLocations")

	var valueList struct {
		Count int
		Value []HistoricalLocationEntity
	}

	json.Unmarshal(b, &valueList)

	ret := make([]HistoricalLocation, len(valueList.Value))
	for i, v := range valueList.Value {
		ret[i] = v
	}

	return ret, err
}

func (c *GossamerClient) FindDatastreams() ([]Datastream, error) {
	b, err := c.doFind("/Datastreams")

	var valueList struct {
		Count int
		Value []DatastreamEntity
	}

	json.Unmarshal(b, &valueList)

	ret := make([]Datastream, len(valueList.Value))
	for i, v := range valueList.Value {
		ret[i] = v
	}

	return ret, err
}

func (c *GossamerClient) FindSensors() ([]Sensor, error) {
	b, err := c.doFind("/Sensors")

	var valueList struct {
		Count int
		Value []SensorEntity
	}

	json.Unmarshal(b, &valueList)

	ret := make([]Sensor, len(valueList.Value))
	for i, v := range valueList.Value {
		ret[i] = v
	}

	return ret, err
}

func (c *GossamerClient) FindFeaturesOfInterest() ([]FeatureOfInterest, error) {
	b, err := c.doFind("/FeaturesOfInterest")

	var valueList struct {
		Count int
		Value []FeatureOfInterestEntity
	}

	json.Unmarshal(b, &valueList)

	ret := make([]FeatureOfInterest, len(valueList.Value))
	for i, v := range valueList.Value {
		ret[i] = v
	}

	return ret, err
}

func (c *GossamerClient) QueryObservations() ClientQuery {
	return &GossamerClientQuery{
		client:     c,
		entityType: ENTITY_OBSERVATIONS,
	}
}

func (c *GossamerClient) QueryThings() ClientQuery {
	return &GossamerClientQuery{
		client:     c,
		entityType: ENTITY_THINGS,
	}
}

func (c *GossamerClient) QueryObservedProperties() ClientQuery {
	return &GossamerClientQuery{
		client:     c,
		entityType: ENTITY_OBSERVEDPROPERTIES,
	}

}

func (c *GossamerClient) QueryLocations() ClientQuery {
	return &GossamerClientQuery{
		client:     c,
		entityType: ENTITY_LOCATIONS,
	}
}

func (c *GossamerClient) QueryDatastreams() ClientQuery {
	return &GossamerClientQuery{
		client:     c,
		entityType: ENTITY_DATASTREAMS,
	}
}

func (c *GossamerClient) QuerySensors() ClientQuery {
	return &GossamerClientQuery{
		client:     c,
		entityType: ENTITY_SENSORS,
	}
}

func (c *GossamerClient) QueryFeaturesOfInterest() ClientQuery {
	return &GossamerClientQuery{
		client:     c,
		entityType: ENTITY_FEATURESOFINTEREST,
	}
}

func (c *GossamerClient) QueryHistoricalLocations() ClientQuery {
	return &GossamerClientQuery{
		client:     c,
		entityType: ENTITY_HISTORICALLOCATIONS,
	}
}

type GossamerClientQuery struct {
	client     Client
	entityType EntityType
	topOption  int
}

func (c *GossamerClientQuery) GetEntityType() EntityType {
	return c.entityType
}

func (c *GossamerClientQuery) Top(v int) ClientQuery {
	c.topOption = v
	return c
}

func (c *GossamerClientQuery) All() ([]SensorThing, error) {
	switch c.entityType {
	case ENTITY_DATASTREAMS:
		o, err := c.client.FindDatastreams()
		st := make([]SensorThing, len(o))
		for i, v := range o {
			st[i] = v
		}
		return st, err

	case ENTITY_LOCATIONS:
		o, err := c.client.FindLocations()
		st := make([]SensorThing, len(o))
		for i, v := range o {
			st[i] = v
		}
		return st, err

	case ENTITY_THINGS:
		o, err := c.client.FindThings()
		st := make([]SensorThing, len(o))
		for i, v := range o {
			st[i] = v
		}
		return st, err

	case ENTITY_OBSERVATIONS:
		o, err := c.client.FindObservations()
		st := make([]SensorThing, len(o))
		for i, v := range o {
			st[i] = v
		}
		return st, err

	case ENTITY_OBSERVEDPROPERTIES:
		o, err := c.client.FindObservedProperties()
		st := make([]SensorThing, len(o))
		for i, v := range o {
			st[i] = v
		}
		return st, err

	case ENTITY_SENSORS:
		o, err := c.client.FindSensors()
		st := make([]SensorThing, len(o))
		for i, v := range o {
			st[i] = v
		}
		return st, err

	case ENTITY_FEATURESOFINTEREST:
		o, err := c.client.FindFeaturesOfInterest()
		st := make([]SensorThing, len(o))
		for i, v := range o {
			st[i] = v
		}
		return st, err

	case ENTITY_HISTORICALLOCATIONS:
		o, err := c.client.FindHistoricalLocations()
		st := make([]SensorThing, len(o))
		for i, v := range o {
			st[i] = v
		}
		return st, err
	}
	return nil, errors.New("Unknown entity to query")
}

func (c *GossamerClientQuery) One(id string) (SensorThing, error) {
	switch c.entityType {
	case ENTITY_DATASTREAMS:
		o, err := c.client.GetDatastream(id)
		return o.(SensorThing), err

	case ENTITY_LOCATIONS:
		o, err := c.client.GetLocation(id)
		return o.(SensorThing), err

	case ENTITY_THINGS:
		o, err := c.client.GetThing(id)
		return o.(SensorThing), err

	case ENTITY_OBSERVATIONS:
		o, err := c.client.GetObservation(id)
		return o.(SensorThing), err

	case ENTITY_OBSERVEDPROPERTIES:
		o, err := c.client.GetObservedProperty(id)
		return o.(SensorThing), err

	case ENTITY_SENSORS:
		o, err := c.client.GetSensor(id)
		return o.(SensorThing), err

	case ENTITY_FEATURESOFINTEREST:
		o, err := c.client.GetFeaturesOfInterest(id)
		return o.(SensorThing), err

	case ENTITY_HISTORICALLOCATIONS:
		o, err := c.client.GetHistoricalLocation(id)
		return o.(SensorThing), err
	}
	return nil, errors.New("Unknown entity to query")
}
