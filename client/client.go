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

func (c *GossamerClient) InsertObservation(o *ObservationEntity) error {
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}

	r := bytes.NewReader(b)

	resp, err := http.Post(c.url+"/v1.0/Observations", "application/json", r)

	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		contents, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(contents))
	}

	return nil
}

func (c *GossamerClient) InsertThing(*ThingEntity) error {
	return nil
}
func (c *GossamerClient) InsertObservedProperty(*ObservedPropertyEntity) error    { return nil }
func (c *GossamerClient) InsertLocation(*LocationEntity) error                    { return nil }
func (c *GossamerClient) InsertDatastream(*DatastreamEntity) error                { return nil }
func (c *GossamerClient) InsertSensor(*SensorEntity) error                        { return nil }
func (c *GossamerClient) InsertFeaturesOfInterest(*FeatureOfInterestEntity) error { return nil }

func (c *GossamerClient) DeleteObservation(string) error        { return nil }
func (c *GossamerClient) DeleteThing(string) error              { return nil }
func (c *GossamerClient) DeleteObservedProperty(string) error   { return nil }
func (c *GossamerClient) DeleteLocation(string) error           { return nil }
func (c *GossamerClient) DeleteDatastream(string) error         { return nil }
func (c *GossamerClient) DeleteSensor(string) error             { return nil }
func (c *GossamerClient) DeleteFeaturesOfInterest(string) error { return nil }

func (c *GossamerClient) UpdateObservation(*ObservationEntity) error              { return nil }
func (c *GossamerClient) UpdateThing(*ThingEntity) error                          { return nil }
func (c *GossamerClient) UpdateObservedProperty(*ObservedPropertyEntity) error    { return nil }
func (c *GossamerClient) UpdateLocation(*LocationEntity) error                    { return nil }
func (c *GossamerClient) UpdateDatastream(*DatastreamEntity) error                { return nil }
func (c *GossamerClient) UpdateSensor(*SensorEntity) error                        { return nil }
func (c *GossamerClient) UpdateFeaturesOfInterest(*FeatureOfInterestEntity) error { return nil }

func (c *GossamerClient) PatchObservation(*ObservationEntity) error              { return nil }
func (c *GossamerClient) PatchThing(*ThingEntity) error                          { return nil }
func (c *GossamerClient) PatchObservedProperty(*ObservedPropertyEntity) error    { return nil }
func (c *GossamerClient) PatchLocation(*LocationEntity) error                    { return nil }
func (c *GossamerClient) PatchDatastream(*DatastreamEntity) error                { return nil }
func (c *GossamerClient) PatchSensor(*SensorEntity) error                        { return nil }
func (c *GossamerClient) PatchFeaturesOfInterest(*FeatureOfInterestEntity) error { return nil }

func (c *GossamerClient) FindObservation(string) (*ObservationEntity, error) { return nil, nil }
func (c *GossamerClient) FindThing(string) (*ThingEntity, error)             { return nil, nil }
func (c *GossamerClient) FindObservedProperty(string) (*ObservedPropertyEntity, error) {
	return nil, nil
}
func (c *GossamerClient) FindLocation(string) (*LocationEntity, error)     { return nil, nil }
func (c *GossamerClient) FindDatastream(string) (*DatastreamEntity, error) { return nil, nil }
func (c *GossamerClient) FindSensor(string) (*SensorEntity, error)         { return nil, nil }
func (c *GossamerClient) FindFeaturesOfInterest(string) (*FeatureOfInterestEntity, error) {
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
