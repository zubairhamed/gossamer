package server

import (
	"github.com/stretchr/testify/assert"
	. "github.com/zubairhamed/gossamer"
	"testing"
)

func TestServer(t *testing.T) {
	server := NewServer()
	assert.NotNil(t, server)
}

func TestDiscoverEntityTypeFunction(t *testing.T) {
	assert.Equal(t, ENTITY_THINGS, DiscoverEntityType("Things"))
	assert.Equal(t, ENTITY_THING, DiscoverEntityType("Thing"))
	assert.Equal(t, ENTITY_LOCATIONS, DiscoverEntityType("Locations"))
	assert.Equal(t, ENTITY_LOCATION, DiscoverEntityType("Location"))
	assert.Equal(t, ENTITY_DATASTREAMS, DiscoverEntityType("Datastreams"))
	assert.Equal(t, ENTITY_DATASTREAM, DiscoverEntityType("Datastream"))
	assert.Equal(t, ENTITY_SENSORS, DiscoverEntityType("Sensors"))
	assert.Equal(t, ENTITY_SENSOR, DiscoverEntityType("Sensor"))
	assert.Equal(t, ENTITY_OBSERVATIONS, DiscoverEntityType("Observations"))
	assert.Equal(t, ENTITY_OBSERVATION, DiscoverEntityType("Observation"))
	assert.Equal(t, ENTITY_OBSERVEDPROPERTIES, DiscoverEntityType("ObservedProperties"))
	assert.Equal(t, ENTITY_OBSERVEDPROPERTY, DiscoverEntityType("ObservedProperty"))
	assert.Equal(t, ENTITY_FEATURESOFINTEREST, DiscoverEntityType("FeaturesOfInterest"))
	assert.Equal(t, ENTITY_HISTORICALLOCATIONS, DiscoverEntityType("HistoricalLocations"))
	assert.Equal(t, ENTITY_HISTORICALLOCATION, DiscoverEntityType("HistoricalLocation"))
	assert.Equal(t, ENTITY_UNKNOWN, DiscoverEntityType("SomethingUnknown"))
}

func TestIsEntityFunction(t *testing.T) {
	assert.True(t, IsEntity("Thing"))
	assert.True(t, IsEntity("Things"))
	assert.True(t, IsEntity("ObservedProperty"))
	assert.True(t, IsEntity("ObservedProperties"))
	assert.True(t, IsEntity("Location"))
	assert.True(t, IsEntity("Locations"))
	assert.True(t, IsEntity("Datastream"))
	assert.True(t, IsEntity("Datastreams"))
	assert.True(t, IsEntity("Sensor"))
	assert.True(t, IsEntity("Sensors"))
	assert.True(t, IsEntity("Observation"))
	assert.True(t, IsEntity("Observations"))
	assert.True(t, IsEntity("FeaturesOfInterest"))
	assert.True(t, IsEntity("FeaturesOfInterests"))
	assert.True(t, IsEntity("HistoricalLocation"))
	assert.True(t, IsEntity("HistoricalLocations"))
	assert.False(t, IsEntity("UNKNOWN"))
}

func TestIsSingularEntityFunction(t *testing.T) {
	assert.True(t, IsSingularEntity("Thing"))
	assert.False(t, IsSingularEntity("Things"))
	assert.True(t, IsSingularEntity("ObservedProperty"))
	assert.False(t, IsSingularEntity("ObservedProperties"))
	assert.True(t, IsSingularEntity("Location"))
	assert.False(t, IsSingularEntity("Locations"))
	assert.True(t, IsSingularEntity("Datastream"))
	assert.False(t, IsSingularEntity("Datastreams"))
	assert.True(t, IsSingularEntity("Sensor"))
	assert.False(t, IsSingularEntity("Sensors"))
	assert.True(t, IsSingularEntity("Observation"))
	assert.False(t, IsSingularEntity("Observations"))
	assert.False(t, IsSingularEntity("FeaturesOfInterest"))
	assert.True(t, IsSingularEntity("HistoricalLocation"))
	assert.False(t, IsSingularEntity("HistoricalLocations"))
	assert.False(t, IsSingularEntity("UNKNOWN"))
}
