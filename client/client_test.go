package client_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/zubairhamed/gossamer"
	"github.com/zubairhamed/gossamer/client"
	"testing"
	"time"
)

func TestClientInsert(t *testing.T) {
	c := client.NewClient("http://localhost:8000")
	var err error

	// Insert Observation
	obs := &gossamer.ObservationEntity{}
	obs.PhenomenonTime = gossamer.NewTimePeriod(time.Now(), time.Now())
	obs.Result = "123"
	obs.ResultTime = gossamer.NewTimeInstant(time.Now())
	ds := &gossamer.DatastreamEntity{}
	ds.Id = "Datastream-1"
	obs.Datastream = ds

	err = c.InsertObservation(obs)
	assert.Nil(t, err)

	// Insert Datastream
	dsEntity := gossamer.NewDatastreamEntity()
	dsEntity.PhenomenonTime = gossamer.NewTimePeriod(time.Now(), time.Now())
	dsEntity.ResultTime = gossamer.NewTimePeriod(time.Now(), time.Now())
	dsEntity.Description = "XXX"
	dsEntity.ObservationType = gossamer.DATASTREAM_OBSTYPE_OBSERVATION
	dsEntity.UnitOfMeasurement = "XXX"

	thing := gossamer.NewThingEntity()
	thing.Id = "ABC123"
	dsEntity.Thing = thing

	sensor := gossamer.NewSensorEntity()
	sensor.Id = "DEF312"
	dsEntity.Sensor = sensor

	obsProp := gossamer.NewObservedPropertyEntity()
	obsProp.Id = "GHI987"
	dsEntity.ObservedProperty = obsProp

	err = c.InsertDatastream(dsEntity)
	assert.Nil(t, err)

	// Insert Feature of Interest
}
