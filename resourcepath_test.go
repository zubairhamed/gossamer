package gossamer

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestResourcePath(t *testing.T) {
	var url *url.URL
	var req Request
	var err error
	var rp ResourcePath
	var prev ResourcePathItem
	var next ResourcePathItem
	var firstItem ResourcePathItem

	url, _ = url.Parse("http://localhost:8000/v1.0/Things")
	req, err = CreateIncomingRequest(url, HTTP)
	assert.Nil(t, err)
	assert.NotNil(t, req)
	rp = req.GetResourcePath()
	assert.NotNil(t, rp)
	assert.Equal(t, 1, len(rp.All()))
	assert.Equal(t, -1, rp.CurrentIndex())
	firstItem = rp.First()
	assert.NotNil(t, firstItem)
	assert.Equal(t, ENTITY_THINGS, firstItem.GetEntity())
	assert.Equal(t, "", firstItem.GetId())
	assert.Nil(t, firstItem.GetQueryOptions())

	// http://localhost:8000/v1.0/Things(12345)
	// http://localhost:8000/v1.0/Things(12345)/Locations
	// http://localhost:8000/v1.0/Things(12345)/Locations(67890)

	url, _ = url.Parse("http://localhost:8000/v1.0/Datastreams(12345)/Sensor")
	req, err = CreateIncomingRequest(url, HTTP)
	assert.Nil(t, err)
	assert.NotNil(t, req)
	rp = req.GetResourcePath()
	assert.NotNil(t, rp)
	assert.Equal(t, 2, len(rp.All()))
	assert.Equal(t, -1, rp.CurrentIndex())
	assert.Equal(t, "12345", rp.First().GetId())
	assert.Equal(t, ENTITY_DATASTREAMS, rp.First().GetEntity())
	assert.True(t, rp.IsFirst())
	assert.NotNil(t, rp.Last())
	assert.True(t, rp.IsLast())
	assert.Equal(t, "", rp.Last().GetId())
	assert.Equal(t, ENTITY_SENSOR, rp.Last().GetEntity())
	prev = rp.Prev()
	assert.Equal(t, "12345", prev.GetId())
	assert.Equal(t, ENTITY_DATASTREAMS, prev.GetEntity())
	next = rp.Next()
	assert.Equal(t, "", next.GetId())
	assert.Equal(t, ENTITY_SENSOR, next.GetEntity())
	assert.False(t, rp.HasNext())
	assert.Nil(t, rp.At(3))
}
