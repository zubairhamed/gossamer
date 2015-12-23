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
	log.Println("========== INSERT ==========")
	InsertObservation(c)
	InsertDatastream(c)
	InsertFeaturesOfInterest(c)
	InsertLocation(c)
	InsertObservedProperty(c)
	InsertThing(c)
	InsertSensor(c)

	log.Println("========== FIND ==========")
	FindObservations(c)
	FindDatastreams(c)
	FindFeaturesOfInterest(c)
	FindLocations(c)
	FindObservedProperties(c)
	FindThings(c)
	FindSensors(c)

	log.Println("========== QUERY ==========")
	QueryObservations(c)

	log.Println("========== UPDATE ==========")
	UpdateObservations(c)
	UpdateDatastreams(c)
	UpdateFeaturesOfInterest(c)
	UpdateLocation(c)
	UpdateObservedProperties(c)
	UpdateThings(c)
	UpdateSensors(c)

	log.Println("========== PATCH ==========")
	PatchObservations(c)
	PatchDatastreams(c)
	PatchFeaturesOfInterest(c)
	PatchLocation(c)
	PatchObservedProperties(c)
	PatchThings(c)
	PatchSensors(c)

	log.Println("========== DELETE ==========")
	DeleteObservations(c)
	DeleteDatastreams(c)
	DeleteFeaturesOfInterest(c)
	DeleteLocations(c)
	DeleteObservedProperties(c)
	DeleteThings(c)
	DeleteSensors(c)
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

func InsertFeaturesOfInterest(c gossamer.Client) {
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

func FindObservations(c gossamer.Client) {
	ol, e := c.FindObservations()
	if e != nil {
		log.Fatal(e)
	}

	for _, v := range ol {
		o, err := c.GetObservation(v.GetId())
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Got Observation Entity ", o.GetId())
		}
	}
}

func FindDatastreams(c gossamer.Client) {
	ol, e := c.FindDatastreams()
	if e != nil {
		log.Fatal(e)
	}

	for _, v := range ol {
		o, err := c.GetDatastream(v.GetId())
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Got Datastream Entity ", o.GetId())
		}
	}
}

func FindFeaturesOfInterest(c gossamer.Client) {
	ol, e := c.FindFeaturesOfInterest()
	if e != nil {
		log.Fatal(e)
	}

	for _, v := range ol {
		o, err := c.GetFeaturesOfInterest(v.GetId())
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Got FeaturesofInterest Entity ", o.GetId())
		}
	}
}

func FindLocations(c gossamer.Client) {
	ol, e := c.FindLocations()
	if e != nil {
		log.Fatal(e)
	}

	for _, v := range ol {
		o, err := c.GetLocation(v.GetId())
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Got Location Entity ", o.GetId())
		}
	}
}

func FindObservedProperties(c gossamer.Client) {
	ol, e := c.FindObservedProperties()
	if e != nil {
		log.Fatal(e)
	}

	for _, v := range ol {
		o, err := c.GetObservedProperty(v.GetId())
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Got Observed Property Entity ", o.GetId())
		}
	}
}

func FindThings(c gossamer.Client) {
	ol, e := c.FindThings()
	if e != nil {
		log.Fatal(e)
	}

	for _, v := range ol {
		o, err := c.GetThing(v.GetId())
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Got Things Entity ", o.GetId())
		}
	}
}

func FindSensors(c gossamer.Client) {
	ol, e := c.FindSensors()
	if e != nil {
		log.Fatal(e)
	}

	for _, v := range ol {
		o, err := c.GetSensor(v.GetId())
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Got Sensor Entity ", o.GetId())
		}
	}
}

func QueryObservations(c gossamer.Client) {
	o, err := c.QueryObservations().All()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Query all Observations returned records # ", len(o))
}

func DeleteObservations(c gossamer.Client) {
	ol, e := c.FindObservations()
	if e != nil {
		log.Fatal(e)
	}

	for _, v := range ol {
		o, err := c.GetObservation(v.GetId())
		if err != nil {
			log.Fatal(err)
		} else {
			// Delete Observation
			e := c.DeleteObservation(o.GetId())
			if e != nil {
				log.Fatal(err)
			} else {
				log.Println("Deleted Observation ", o.GetId())
			}
		}
	}
}

func DeleteDatastreams(c gossamer.Client) {
	ol, e := c.FindDatastreams()
	if e != nil {
		log.Fatal(e)
	}

	for _, v := range ol {
		o, err := c.GetDatastream(v.GetId())
		if err != nil {
			log.Fatal(err)
		} else {
			// Delete Observation
			e := c.DeleteDatastream(o.GetId())
			if e != nil {
				log.Fatal(err)
			} else {
				log.Println("Deleted Datastream ", o.GetId())
			}
		}
	}
}

func DeleteFeaturesOfInterest(c gossamer.Client) {
	ol, e := c.FindFeaturesOfInterest()
	if e != nil {
		log.Fatal(e)
	}

	for _, v := range ol {
		o, err := c.GetFeaturesOfInterest(v.GetId())
		if err != nil {
			log.Fatal(err)
		} else {
			// Delete Observation
			e := c.DeleteFeaturesOfInterest(o.GetId())
			if e != nil {
				log.Fatal(err)
			} else {
				log.Println("Deleted FeatureOfInterest ", o.GetId())
			}
		}
	}
}

func DeleteLocations(c gossamer.Client) {
	ol, e := c.FindLocations()
	if e != nil {
		log.Fatal(e)
	}

	for _, v := range ol {
		o, err := c.GetLocation(v.GetId())
		if err != nil {
			log.Fatal(err)
		} else {
			// Delete Observation
			e := c.DeleteLocation(o.GetId())
			if e != nil {
				log.Fatal(err)
			} else {
				log.Println("Deleted Location ", o.GetId())
			}
		}
	}
}

func DeleteObservedProperties(c gossamer.Client) {
	ol, e := c.FindObservedProperties()
	if e != nil {
		log.Fatal(e)
	}

	for _, v := range ol {
		o, err := c.GetObservedProperty(v.GetId())
		if err != nil {
			log.Fatal(err)
		} else {
			// Delete Observation
			e := c.DeleteObservedProperty(o.GetId())
			if e != nil {
				log.Fatal(err)
			} else {
				log.Println("Deleted ObservedProperty ", o.GetId())
			}
		}
	}
}

func DeleteThings(c gossamer.Client) {
	ol, e := c.FindThings()
	if e != nil {
		log.Fatal(e)
	}

	for _, v := range ol {
		o, err := c.GetThing(v.GetId())
		if err != nil {
			log.Fatal(err)
		} else {
			// Delete Observation
			e := c.DeleteThing(o.GetId())
			if e != nil {
				log.Fatal(err)
			} else {
				log.Println("Deleted Thing ", o.GetId())
			}
		}
	}
}

func DeleteSensors(c gossamer.Client) {
	ol, e := c.FindSensors()
	if e != nil {
		log.Fatal(e)
	}

	for _, v := range ol {
		o, err := c.GetSensor(v.GetId())
		if err != nil {
			log.Fatal(err)
		} else {
			// Delete Observation
			e := c.DeleteSensor(o.GetId())
			if e != nil {
				log.Fatal(err)
			} else {
				log.Println("Deleted Sensor ", o.GetId())
			}
		}
	}
}

func UpdateObservations(c gossamer.Client) {
	l, err := c.QueryObservations().Top(1).All()
	if err != nil {
		log.Fatal(err)
	}

	var o gossamer.Observation
	if len(l) > 0 {
		o = l[0].(gossamer.Observation)

		n := gossamer.NewObservationEntity()
		n.Id = o.GetId()
		n.ResultQuality = "Patch this value only"
		n.ResultTime = gossamer.NewTimeInstant(time.Now())

		c.UpdateObservation(n)
	}
}

func UpdateDatastreams(c gossamer.Client) {

}

func UpdateFeaturesOfInterest(c gossamer.Client) {

}

func UpdateLocation(c gossamer.Client)           {}
func UpdateObservedProperties(c gossamer.Client) {}
func UpdateThings(c gossamer.Client)             {}
func UpdateSensors(c gossamer.Client)            {}

func PatchObservations(c gossamer.Client) {
	l, err := c.QueryObservations().Top(1).All()
	if err != nil {
		log.Fatal(err)
	}

	var o gossamer.Observation
	if len(l) > 0 {
		o = l[0].(gossamer.Observation)

		n := gossamer.CloneObservationEntity(o)
		n.ResultQuality = "Value was changed"
		n.ResultTime = gossamer.NewTimeInstant(time.Now())

		c.PatchObservation(n)
	}
}

func PatchDatastreams(c gossamer.Client)        {}
func PatchFeaturesOfInterest(c gossamer.Client) {}
func PatchLocation(c gossamer.Client)           {}
func PatchObservedProperties(c gossamer.Client) {}
func PatchThings(c gossamer.Client)             {}
func PatchSensors(c gossamer.Client)            {}
