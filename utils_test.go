package gossamer

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUtilFunctions(t *testing.T) {
	assert.Equal(t, "Things(12345)", ResolveEntityLink("12345", ENTITY_THINGS))
	assert.Equal(t, "http://localhost:8000/v1.0/Things(12345)", ResolveSelfLinkUrl("12345", "Things"))
}

func TestValidateMandatoryPropertiesForThing(t *testing.T) {
	var err error

	thingEntity := NewThingEntity()
	err = ValidateMandatoryProperties(thingEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for Thing entity: 'description'", err.Error())

	thingEntity.Description = "XXXXXX"
	err = ValidateMandatoryProperties(thingEntity)
	assert.Nil(t, err)
}

func TestValidateMandatoryPropertiesForLocation(t *testing.T) {
	var err error

	locationEntity := NewLocationEntity()
	err = ValidateMandatoryProperties(locationEntity)
	assert.NotNil(t, err)

	locationEntity.Description = "XXXXXXX"
	err = ValidateMandatoryProperties(locationEntity)
	assert.NotNil(t, err)

	locationEntity.EncodingType = "XXXXXXX"
	err = ValidateMandatoryProperties(locationEntity)
	assert.NotNil(t, err)

	locationEntity.Location = "XXXXXXX"
	err = ValidateMandatoryProperties(locationEntity)
	assert.Nil(t, err)
}

func TestValidateMandatoryPropertiesForHistoricalLocation(t *testing.T) {
	var err error

	historicalLocation := NewHistoricalLocationEntity()
	err = ValidateMandatoryProperties(historicalLocation)
	assert.NotNil(t, err)

	historicalLocation.Time = time.Now()
	err = ValidateMandatoryProperties(historicalLocation)
	assert.Nil(t, err)
}

func TestValidateMandatoryPropertiesForDatastream(t *testing.T) {
	var err error

	datastreamEntity := NewDatastreamEntity()
	err = ValidateMandatoryProperties(datastreamEntity)
	assert.NotNil(t, err)

	datastreamEntity.Description = "XXXXXXX"
	err = ValidateMandatoryProperties(datastreamEntity)
	assert.NotNil(t, err)

	datastreamEntity.UnitOfMeasurement = "XXXXXXX"
	err = ValidateMandatoryProperties(datastreamEntity)
	assert.NotNil(t, err)

	datastreamEntity.ObservationType = "XXXXXXX"
	err = ValidateMandatoryProperties(datastreamEntity)
	assert.Nil(t, err)
}

func TestValidateMandatoryPropertiesForSensor(t *testing.T) {
	var err error

	sensorEntity := NewSensorEntity()
	err = ValidateMandatoryProperties(sensorEntity)
	assert.NotNil(t, err)

	sensorEntity.Description = "XXXXXXX"
	err = ValidateMandatoryProperties(sensorEntity)
	assert.NotNil(t, err)

	sensorEntity.EncodingType = "XXXXXXX"
	err = ValidateMandatoryProperties(sensorEntity)
	assert.NotNil(t, err)

	sensorEntity.Metadata = "XXXXXXX"
	err = ValidateMandatoryProperties(sensorEntity)
	assert.Nil(t, err)
}

func TestValidateMandatoryPropertiesForObservedProperty(t *testing.T) {
	var err error

	observedPropertyEntity := NewObservedPropertyEntity()
	err = ValidateMandatoryProperties(observedPropertyEntity)
	assert.NotNil(t, err)

	observedPropertyEntity.Name = "XXXXXXX"
	err = ValidateMandatoryProperties(observedPropertyEntity)
	assert.NotNil(t, err)

	observedPropertyEntity.Definition = "XXXXXXX"
	err = ValidateMandatoryProperties(observedPropertyEntity)
	assert.NotNil(t, err)

	observedPropertyEntity.Description = "XXXXXXX"
	err = ValidateMandatoryProperties(observedPropertyEntity)
	assert.Nil(t, err)
}

func TestValidateMandatoryPropertiesForObservation(t *testing.T) {
	var err error

	observationEntity := NewObservationEntity()
	err = ValidateMandatoryProperties(observationEntity)
	assert.NotNil(t, err)

	observationEntity.PhenomenonTime = time.Now()
	err = ValidateMandatoryProperties(observationEntity)
	assert.NotNil(t, err)

	observationEntity.ResultTime = time.Now()
	err = ValidateMandatoryProperties(observationEntity)
	assert.NotNil(t, err)

	observationEntity.Result = "XXXXXXX"
	err = ValidateMandatoryProperties(observationEntity)
	assert.Nil(t, err)
}

func TestValidateMandatoryPropertiesForFeatureOfInterest(t *testing.T) {
	var err error

	featureOfInterestEntity := NewFeatureOfInterestEntity()
	err = ValidateMandatoryProperties(featureOfInterestEntity)
	assert.NotNil(t, err)

	featureOfInterestEntity.Description = "XXXXXXX"
	err = ValidateMandatoryProperties(featureOfInterestEntity)
	assert.NotNil(t, err)

	featureOfInterestEntity.EncodingType = "XXXXXXX"
	err = ValidateMandatoryProperties(featureOfInterestEntity)
	assert.NotNil(t, err)

	featureOfInterestEntity.Feature = "XXXXXXX"
	err = ValidateMandatoryProperties(featureOfInterestEntity)
	assert.Nil(t, err)
}

func TestValidateIntegrityConstraints(t *testing.T) {

}
