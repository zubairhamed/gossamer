package gossamer

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestResolveEntityLinkFunctions(t *testing.T) {
	assert.Equal(t, "Things(12345)", ResolveEntityLink("12345", ENTITY_THINGS))
}

func TestResolveSelfLinkUrlFunctions(t *testing.T) {
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
	assert.Equal(t, "Missing mandatory property for Location entity: 'description'", err.Error())

	locationEntity.Description = "XXXXXXX"
	err = ValidateMandatoryProperties(locationEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for Location entity: 'encodingType'", err.Error())

	locationEntity.EncodingType = "XXXXXXX"
	err = ValidateMandatoryProperties(locationEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for Location entity: 'location'", err.Error())

	locationEntity.Location = "XXXXXXX"
	err = ValidateMandatoryProperties(locationEntity)
	assert.Nil(t, err)
}

func TestValidateMandatoryPropertiesForHistoricalLocation(t *testing.T) {
	var err error

	historicalLocation := NewHistoricalLocationEntity()
	err = ValidateMandatoryProperties(historicalLocation)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for HistoricalLocation entity: 'time'", err.Error())

	historicalLocation.Time = time.Now()
	err = ValidateMandatoryProperties(historicalLocation)
	assert.Nil(t, err)
}

func TestValidateMandatoryPropertiesForDatastream(t *testing.T) {
	var err error

	datastreamEntity := NewDatastreamEntity()
	err = ValidateMandatoryProperties(datastreamEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for Datastream entity: 'description'", err.Error())

	datastreamEntity.Description = "XXXXXXX"
	err = ValidateMandatoryProperties(datastreamEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for Datastream entity: 'unitOfMeasurement'", err.Error())

	datastreamEntity.UnitOfMeasurement = "XXXXXXX"
	err = ValidateMandatoryProperties(datastreamEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for Datastream entity: 'observationType'", err.Error())

	datastreamEntity.ObservationType = "XXXXXXX"
	err = ValidateMandatoryProperties(datastreamEntity)
	assert.Nil(t, err)
}

func TestValidateMandatoryPropertiesForSensor(t *testing.T) {
	var err error

	sensorEntity := NewSensorEntity()
	err = ValidateMandatoryProperties(sensorEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for Sensor entity: 'description'", err.Error())

	sensorEntity.Description = "XXXXXXX"
	err = ValidateMandatoryProperties(sensorEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for Sensor entity: 'encodingType'", err.Error())

	sensorEntity.EncodingType = "XXXXXXX"
	err = ValidateMandatoryProperties(sensorEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for Sensor entity: 'metadata'", err.Error())

	sensorEntity.Metadata = "XXXXXXX"
	err = ValidateMandatoryProperties(sensorEntity)
	assert.Nil(t, err)
}

func TestValidateMandatoryPropertiesForObservedProperty(t *testing.T) {
	var err error

	observedPropertyEntity := NewObservedPropertyEntity()
	err = ValidateMandatoryProperties(observedPropertyEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for ObservedProperty entity: 'name'", err.Error())

	observedPropertyEntity.Name = "XXXXXXX"
	err = ValidateMandatoryProperties(observedPropertyEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for ObservedProperty entity: 'definition'", err.Error())

	observedPropertyEntity.Definition = "XXXXXXX"
	err = ValidateMandatoryProperties(observedPropertyEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for ObservedProperty entity: 'description'", err.Error())

	observedPropertyEntity.Description = "XXXXXXX"
	err = ValidateMandatoryProperties(observedPropertyEntity)
	assert.Nil(t, err)
}

func TestValidateMandatoryPropertiesForObservation(t *testing.T) {
	var err error

	observationEntity := NewObservationEntity()
	err = ValidateMandatoryProperties(observationEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for Observation entity: 'phenomenonTime'", err.Error())

	observationEntity.PhenomenonTime = NewTimePeriod(time.Now(), time.Now())
	err = ValidateMandatoryProperties(observationEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for Observation entity: 'resultTime'", err.Error())

	observationEntity.ResultTime = TimeInstant(time.Now())
	err = ValidateMandatoryProperties(observationEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for Observation entity: 'result'", err.Error())

	observationEntity.Result = "XXXXXXX"
	err = ValidateMandatoryProperties(observationEntity)
	assert.Nil(t, err)
}

func TestValidateMandatoryPropertiesForFeatureOfInterest(t *testing.T) {
	var err error

	featureOfInterestEntity := NewFeatureOfInterestEntity()
	err = ValidateMandatoryProperties(featureOfInterestEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for FeaturesOfInterest entity: 'description'", err.Error())

	featureOfInterestEntity.Description = "XXXXXXX"
	err = ValidateMandatoryProperties(featureOfInterestEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for FeaturesOfInterest entity: 'encodingType'", err.Error())

	featureOfInterestEntity.EncodingType = "XXXXXXX"
	err = ValidateMandatoryProperties(featureOfInterestEntity)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing mandatory property for FeaturesOfInterest entity: 'feature'", err.Error())

	featureOfInterestEntity.Feature = "XXXXXXX"
	err = ValidateMandatoryProperties(featureOfInterestEntity)
	assert.Nil(t, err)
}

func TestValidateIntegrityConstraintsForThings(t *testing.T) {
	var err error

	thingEntity := NewThingEntity()

	err = ValidateIntegrityConstraints(thingEntity)
	assert.Nil(t, err)
}

func TestValidateIntegrityConstraintsForLocations(t *testing.T) {
	var err error

	locationEntity := NewLocationEntity()

	err = ValidateIntegrityConstraints(locationEntity)
	assert.Nil(t, err)
}

func TestValidateIntegrityConstraintsForDatastreams(t *testing.T) {
	var err error
	var datastream *DatastreamEntity
	datastream = NewDatastreamEntity()

	err = ValidateIntegrityConstraints(datastream)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing constrains for Datastream Entity: 'Thing'", err.Error())

	datastream.IdThing = "12345"
	err = ValidateIntegrityConstraints(datastream)
	assert.Equal(t, "Missing constrains for Datastream Entity: 'Sensor'", err.Error())

	datastream.IdSensor = "12345"
	err = ValidateIntegrityConstraints(datastream)
	assert.Equal(t, "Missing constrains for Datastream Entity: 'ObservedProperty'", err.Error())

	datastream.IdObservedProperty = "12345"
	err = ValidateIntegrityConstraints(datastream)
	assert.Nil(t, err)
}

func TestValidateIntegrityConstraintsForSensors(t *testing.T) {
	var err error

	sensorEntity := NewSensorEntity()

	err = ValidateIntegrityConstraints(sensorEntity)
	assert.Nil(t, err)
}

func TestValidateIntegrityConstraintsForObservedProperties(t *testing.T) {
	var err error

	observedProperty := NewObservedPropertyEntity()

	err = ValidateIntegrityConstraints(observedProperty)
	assert.Nil(t, err)
}

func TestValidateIntegrityConstraintsForObservations(t *testing.T) {
	var err error

	observation := NewObservationEntity()

	err = ValidateIntegrityConstraints(observation)
	assert.NotNil(t, err)
	assert.Equal(t, "Missing constrains for Observation Entity: 'Datastream'", err.Error())
}

func TestValidateIntegrityConstraintsForFeaturesOfInterest(t *testing.T) {
	var err error

	featuresOfInterest := NewFeatureOfInterestEntity()

	err = ValidateIntegrityConstraints(featuresOfInterest)
	assert.Nil(t, err)

}
