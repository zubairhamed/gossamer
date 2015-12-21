package client

import (
	"bytes"
	"github.com/zubairhamed/gossamer"
	"log"
	"strings"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type Client interface {
	QueryAll(gossamer.EntityType, gossamer.QueryOptions) ([]gossamer.SensorThing, error)
	QueryOne(gossamer.EntityType, gossamer.QueryOptions) (gossamer.SensorThing, error)

	InsertObservation(*gossamer.ObservationEntity) (error)
}


func NewClient(url string) Client {
	return &GossamerClient{
		url: url,
	}
}

type GossamerClient struct {
	url		string
}

func (c *GossamerClient) QueryAll(gossamer.EntityType, gossamer.QueryOptions) ([]gossamer.SensorThing, error) {
	return nil, nil
}

func (c *GossamerClient) QueryOne(gossamer.EntityType, gossamer.QueryOptions) (gossamer.SensorThing, error) {
	return nil, nil
}

func (c *GossamerClient) InsertObservation(o *gossamer.ObservationEntity) (error) {
	b, err := json.Marshal(o)

	r := bytes.NewReader(b)

	resp, err := http.Post(c.url + "/v1.0/Observations", "application/json", r)
	contents, _ := ioutil.ReadAll(resp.Body)
	log.Println(resp, err, contents)



	return nil
}

type ClientQuery interface {
	All() ([]gossamer.SensorThing, error)
	One() (gossamer.SensorThing, error)
	Filter() ClientQuery
	Count(bool) ClientQuery
	OrderBy(...string) ClientQuery
	Skip(int) ClientQuery
	Top(int) ClientQuery
	Expand(...string) ClientQuery
	Select(...string) ClientQuery
	GetUrlString() string
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

func (cq *DefaultClientQuery) All() ([]gossamer.SensorThing, error) {
	return nil, nil
}

func (cq *DefaultClientQuery) One() (gossamer.SensorThing, error) {
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

type ClientDelete interface {
}

type DefaultClientDelete struct {
	entityClient EntityClient
}

type ClientPatch interface {
}

type DefaultClientPatch struct {
	entityClient EntityClient
}

type ClientUpdate interface {
}

type DefaultClientUpdate struct {
	entityClient EntityClient
}

type ClientInsert interface {

}

type DefaultClientInsert struct {
	entityClient EntityClient
}

type EntityClient interface {
	GetType() gossamer.EntityType
	GetId() string
	Query() ClientQuery
	Delete() ClientDelete
	Patch() ClientPatch
	Update() ClientUpdate
	Insert() ClientInsert
}

type DefaultEntityClient struct {
	client 	   Client
	entityType gossamer.EntityType
	id         string
}

func (c *DefaultEntityClient) GetType() gossamer.EntityType {
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
