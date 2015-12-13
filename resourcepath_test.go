package gossamer

import (
	"testing"
	"net/url"
	"github.com/stretchr/testify/assert"
)

func TestResourcePath(t *testing.T) {
	var url *url.URL
	var req Request
	var err error
	var rp ResourcePath
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


}