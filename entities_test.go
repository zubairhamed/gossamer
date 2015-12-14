package gossamer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntities(t *testing.T) {
	assert.Equal(t, ENTITY_THINGS, ThingEntity{}.GetType())
	assert.Equal(t, ENTITY_LOCATIONS, LocationEntity{}.GetType())
	assert.Equal(t, ENTITY_HISTORICALLOCATIONS, HistoricalLocationEntity{}.GetType())
	assert.Equal(t, ENTITY_DATASTREAMS, DatastreamEntity{}.GetType())
	assert.Equal(t, ENTITY_SENSORS, SensorEntity{}.GetType())
	assert.Equal(t, ENTITY_OBSERVEDPROPERTIES, ObservedPropertyEntity{}.GetType())
	assert.Equal(t, ENTITY_OBSERVATIONS, ObservationEntity{}.GetType())
	assert.Equal(t, ENTITY_FEATURESOFINTERESTS, FeatureOfInterestEntity{}.GetType())
}
