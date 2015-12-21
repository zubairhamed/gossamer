package main

import (
	"github.com/zubairhamed/gossamer"
	"github.com/zubairhamed/gossamer/client"
	"log"
	"time"
)

func main() {
	c := client.NewClient("http://localhost:8000")

	// Inserts
	InsertObservation(c)
	InsertDatastream(c)
	InsertFeatureOfInterest(c)
	InsertLocation(c)
	InsertObservedProperty(c)
	InsertThing(c)
	InsertSensor(c)

	o, e := c.GetObservation("b8c26429-4f53-4a02-93e3-29fef8bd4455")
	log.Println(o, e)
}

func InsertObservation(c gossamer.Client) {
	obs := &gossamer.ObservationEntity{}
	obs.PhenomenonTime = gossamer.NewTimePeriod(time.Now(), time.Now())
	obs.Result = "123"
	obs.ResultTime = gossamer.NewTimeInstant(time.Now())
	ds := &gossamer.DatastreamEntity{}
	ds.Id = "Datastream-1"
	obs.Datastream = ds

	err := c.InsertObservation(obs)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted New Observation")
}

func InsertDatastream(c gossamer.Client) {
	e := gossamer.NewDatastreamEntity()
	e.PhenomenonTime = gossamer.NewTimePeriod(time.Now(), time.Now())
	e.ResultTime = gossamer.NewTimePeriod(time.Now(), time.Now())
	e.Description = "XXX"
	e.ObservationType = gossamer.DATASTREAM_OBSTYPE_OBSERVATION
	e.UnitOfMeasurement = "XXX"

	thing := gossamer.NewThingEntity()
	thing.Id = "ABC123"
	e.Thing = thing

	sensor := gossamer.NewSensorEntity()
	sensor.Id = "DEF312"
	e.Sensor = sensor

	obsProp := gossamer.NewObservedPropertyEntity()
	obsProp.Id = "GHI987"
	e.ObservedProperty = obsProp

	err := c.InsertDatastream(e)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted New Datastream")
}

func InsertFeatureOfInterest(c gossamer.Client) {
	e := gossamer.NewFeatureOfInterestEntity()
	e.Description = "XXXX"
	e.EncodingType = gossamer.LOCATION_ENCTYPE_GEOJSON
	e.Feature = "Feature ABC 1 2 3"

	err := c.InsertFeaturesOfInterest(e)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted New Feature of Interest")
}

func InsertLocation(c gossamer.Client) {
	e := gossamer.NewLocationEntity()
	e.Description = "XXXXX"
	e.EncodingType = gossamer.LOCATION_ENCTYPE_GEOJSON
	e.Location = "XOXOXO"

	err := c.InsertLocation(e)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted New Location")
}

func InsertObservedProperty(c gossamer.Client) {
	e := gossamer.NewObservedPropertyEntity()
	e.Name = "XXXX"
	e.Description = "XXXXXX"
	e.Definition = "XXXXX"

	err := c.InsertObservedProperty(e)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted New ObservedProperty")
}

func InsertSensor(c gossamer.Client) {
	e := gossamer.NewSensorEntity()
	e.Description = "XXXXXX"
	e.EncodingType = gossamer.SENSOR_ENCTYPE_SENSORML
	e.Metadata = "XXXXX"

	err := c.InsertSensor(e)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted New Sensor")
}

func InsertThing(c gossamer.Client) {
	e := gossamer.NewThingEntity()
	e.Description = "XXXXX"

	err := c.InsertThing(e)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted New Thing")
}

