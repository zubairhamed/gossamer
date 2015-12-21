package client

import (
	"bytes"
	"encoding/json"
	"errors"
	. "github.com/zubairhamed/gossamer"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func NewClient(url string) Client {
	return &GossamerClient{
		url: url,
	}
}

type GossamerClient struct {
	url string
}

func (c *GossamerClient) QueryAll(EntityType, QueryOptions) ([]SensorThing, error) {
	return nil, nil
}

func (c *GossamerClient) QueryOne(EntityType, QueryOptions) (SensorThing, error) {
	return nil, nil
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

func (c *GossamerClient) DeleteObservation(string) error        { return nil }
func (c *GossamerClient) DeleteThing(string) error              { return nil }
func (c *GossamerClient) DeleteObservedProperty(string) error   { return nil }
func (c *GossamerClient) DeleteLocation(string) error           { return nil }
func (c *GossamerClient) DeleteDatastream(string) error         { return nil }
func (c *GossamerClient) DeleteSensor(string) error             { return nil }
func (c *GossamerClient) DeleteFeaturesOfInterest(string) error { return nil }

func (c *GossamerClient) UpdateObservation(o Observation) error              { return nil }
func (c *GossamerClient) UpdateThing(o Thing) error                          { return nil }
func (c *GossamerClient) UpdateObservedProperty(o ObservedProperty) error    { return nil }
func (c *GossamerClient) UpdateLocation(o Location) error                    { return nil }
func (c *GossamerClient) UpdateDatastream(o Datastream) error                { return nil }
func (c *GossamerClient) UpdateSensor(o Sensor) error                        { return nil }
func (c *GossamerClient) UpdateFeaturesOfInterest(o FeatureOfInterest) error { return nil }

func (c *GossamerClient) PatchObservation(o Observation) error              { return nil }
func (c *GossamerClient) PatchThing(o Thing) error                          { return nil }
func (c *GossamerClient) PatchObservedProperty(o ObservedProperty) error    { return nil }
func (c *GossamerClient) PatchLocation(o Location) error                    { return nil }
func (c *GossamerClient) PatchDatastream(o Datastream) error                { return nil }
func (c *GossamerClient) PatchSensor(o Sensor) error                        { return nil }
func (c *GossamerClient) PatchFeaturesOfInterest(o FeatureOfInterest) error { return nil }

func (c *GossamerClient) FindObservation(string) ([]Observation, error)           { return nil, nil }
func (c *GossamerClient) FindThing(string) ([]Thing, error)                       { return nil, nil }
func (c *GossamerClient) FindObservedProperty(string) ([]ObservedProperty, error) { return nil, nil }
func (c *GossamerClient) FindLocation(string) ([]Location, error)                 { return nil, nil }
func (c *GossamerClient) FindDatastream(string) ([]Datastream, error)             { return nil, nil }
func (c *GossamerClient) FindSensor(string) ([]Sensor, error)                     { return nil, nil }
func (c *GossamerClient) FindFeaturesOfInterest(string) ([]FeatureOfInterest, error) {
	return nil, nil
}

type DefaultClientQuery struct {
	entityClient EntityClient
	optOrderBy   []string
	optExpand    []string
	optSelect    []string
	optTop       int
	optCount     bool
	optSkip      int
}

func (cq *DefaultClientQuery) All() ([]SensorThing, error) {
	return nil, nil
}

func (cq *DefaultClientQuery) One() (SensorThing, error) {
	log.Println(cq.GetUrlString())
	return nil, nil
}

func (cq *DefaultClientQuery) Filter() ClientQuery {
	return cq
}

func (cq *DefaultClientQuery) Count(b bool) ClientQuery {
	cq.optCount = b
	return cq
}

func (cq *DefaultClientQuery) OrderBy(v ...string) ClientQuery {
	cq.optOrderBy = v
	return cq
}

func (cq *DefaultClientQuery) Skip(n int) ClientQuery {
	cq.optSkip = n
	return cq
}

func (cq *DefaultClientQuery) Top(n int) ClientQuery {
	cq.optTop = n
	return cq
}

func (cq *DefaultClientQuery) Expand(v ...string) ClientQuery {
	cq.optExpand = v
	return cq
}

func (cq *DefaultClientQuery) Select(v ...string) ClientQuery {
	cq.optSelect = v
	return cq
}

func (cq *DefaultClientQuery) GetUrlString() string {
	buf := bytes.NewBufferString("/v1.0/" + string(cq.entityClient.GetType()))

	if cq.entityClient.GetId() != "" {
		buf.WriteString("(" + cq.entityClient.GetId() + ")")
	}
	buf.WriteString("?")

	if len(cq.optExpand) > 0 {
		buf.WriteString("$expand=" + strings.Join(cq.optExpand, ",") + "&")
	}

	if len(cq.optOrderBy) > 0 {
		buf.WriteString("$orderby=" + strings.Join(cq.optOrderBy, ",") + "&")
	}

	if len(cq.optSelect) > 0 {
		buf.WriteString("$select=" + strings.Join(cq.optSelect, ",") + "&")
	}

	if cq.optCount {
		buf.WriteString("$count=true" + "&")
	} else {
		buf.WriteString("$count=false" + "&")
	}

	if cq.optSkip != 0 {
		buf.WriteString("$skip=" + string(cq.optSkip) + "&")
	}

	if cq.optTop != 0 {
		buf.WriteString("$top=" + string(cq.optTop) + "&")
	}
	return buf.String()
}

type DefaultClientDelete struct {
	entityClient EntityClient
}

type DefaultClientPatch struct {
	entityClient EntityClient
}

type DefaultClientUpdate struct {
	entityClient EntityClient
}

type DefaultClientInsert struct {
	entityClient EntityClient
}

type DefaultEntityClient struct {
	client     Client
	entityType EntityType
	id         string
}

func (c *DefaultEntityClient) GetType() EntityType {
	return c.entityType
}

func (c *DefaultEntityClient) GetId() string {
	return c.id
}

func (c *DefaultEntityClient) Query() ClientQuery {
	return &DefaultClientQuery{
		entityClient: c,
	}
}

func (c *DefaultEntityClient) Delete() ClientDelete {
	return &DefaultClientDelete{
		entityClient: c,
	}
}

func (c *DefaultEntityClient) Patch() ClientPatch {
	return &DefaultClientPatch{
		entityClient: c,
	}
}

func (c *DefaultEntityClient) Update() ClientUpdate {
	return &DefaultClientUpdate{
		entityClient: c,
	}
}

func (c *DefaultEntityClient) Insert() ClientInsert {
	return &DefaultClientInsert{
		entityClient: c,
	}
}
