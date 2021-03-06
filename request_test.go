package gossamer

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestParseRequestUrl(t *testing.T) {
	cases1 := []struct {
		in string
	}{
		/*
			EntityType
			ID
			Query Opt Values
				expand
				top
				skip
				count
				filter
			property
			propertyValue

		*/
		{"http://example.org/v1.0/ObservedProperties"},
		{"http://example.org/v1.0/Things(1)"},
		{"http://example.org/v1.0/Things(1)/Locations"},
		{"http://example.org/v1.0/Things?$expand=Datastreams"},
		{"http://example.org/v1.0/Things?$top=5"},
		{"http://example.org/v1.0/Things?$skip=5"},
		{"http://example.org/v1.0/Things?$count=true"},
		{"http://example.org/v1.0/Things?$filter=geo.distance(Locations/location,￼geography'POINT(-122, 43)') gt 1"},
		{"http://example.org/v1.0/Things?$expand=Datastreams/Observations/FeatureOfInterest&$filter=Datastreams/Observations/FeatureOfInterest/id eq 'FOI_1' and Datastreams/Observations/resultTime ge 2010-06-01T00:00:00Z and Datastreams/Observations/resultTime le 2010-07-01T00:00:00Z"},

		{"http://example.org/v1.0/Observations?$select=result,resultTime"},
		{"http://example.org/v1.0/Observations?$orderby=result"},
		{"http://example.org/v1.0/Observations?$expand=Datastream&$orderby=Datastreams/id desc, phenomenonTime"},
		{"http://example.org/v1.0/Observations(1)/resultTime"},
		{"http://example.org/v1.0/Observations(1)/resultTime/$value"},
		{"http://example.org/v1.0/Datastreams(1)/Observations"},
		{"http://example.org/v1.0/Datastreams(1)/Observations/$ref"},
		{"http://example.org/v1.0/Datastreams(1)/Observations(1)"},
		{"http://example.org/v1.0/Datastreams(1)/Observations(1)/resultTime"},
		{"http://example.org/v1.0/Datastreams(1)/Observations(1)/FeatureOfInterest"},
		{"http://example.org/v1.0/Datastreams(1)?$expand=Observations,ObservedProperty"},
		{"http://example.org/v1.0/Datastreams(1)?$expand=Observations($filter=result eq 1)"},
		{"http://example.org/v1.0/Datastreams(1)?$select=id,Observations&$expand=Observations/FeatureOfInterest"},
		{"http://example.org/v1.0/Datastreams(1)?$expand=Observations($select=result)"},
		{"http://example.org/v1.0/Observations?$top=5&$orderby=phenomenonTime desc"},
		{"http://example.org/v1.0/Observations?$skip=2&$top=2&$orderby=resultTime"},
		{"http://example.org/v1.0/Observations?$filter=result lt 10.00"},
		{"http://example.org/v1.0/Observations?$filter=Datastream/id eq '1'"},
	}

	for _, c := range cases1 {
		var err error
		reqUrl, err := url.Parse(c.in)
		assert.Nil(t, err)
		assert.NotNil(t, reqUrl)

		req, err := CreateIncomingRequest(reqUrl, HTTP)
		assert.Nil(t, err, "Error occured parsing URL: "+c.in)
		assert.NotNil(t, req)
	}

	var reqUrl *url.URL
	var req Request

	reqUrl, _ = url.Parse("http://example.org/v1.0/ObservedProperties")
	req, _ = CreateIncomingRequest(reqUrl, HTTP)
	assert.NotNil(t, req)

	assert.NotNil(t, req.GetResourcePath())
	assert.Equal(t, 1, len(req.GetResourcePath().All()))
	assert.Equal(t, ENTITY_OBSERVEDPROPERTIES, req.GetResourcePath().First().GetEntity())
	assert.Empty(t, req.GetResourcePath().First().GetId())

	reqUrl, _ = url.Parse("http://example.org/v1.0/Things(1)")
	req, _ = CreateIncomingRequest(reqUrl, HTTP)
	assert.NotNil(t, req)
	assert.NotNil(t, req.GetResourcePath())
	assert.Equal(t, 1, len(req.GetResourcePath().All()))
	assert.Equal(t, ENTITY_THINGS, req.GetResourcePath().First().GetEntity())
	assert.Equal(t, "1", req.GetResourcePath().First().GetId())

	reqUrl, _ = url.Parse("http://example.org/v1.0/Things?$expand=Datastreams")
	req, _ = CreateIncomingRequest(reqUrl, HTTP)
	assert.NotNil(t, req)
	assert.NotNil(t, req.GetResourcePath())
	assert.Equal(t, 1, len(req.GetResourcePath().All()))
	assert.Equal(t, ENTITY_THINGS, req.GetResourcePath().First().GetEntity())
	assert.Empty(t, req.GetResourcePath().First().GetId())

	reqUrl, _ = url.Parse("http://example.org/v1.0/Things(1)/Observations")
	req, _ = CreateIncomingRequest(reqUrl, HTTP)
	assert.NotNil(t, req)
	assert.NotNil(t, req.GetResourcePath())
	assert.Equal(t, 2, len(req.GetResourcePath().All()))
	assert.Equal(t, ENTITY_THINGS, req.GetResourcePath().First().GetEntity())
}
