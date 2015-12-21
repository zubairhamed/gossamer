package main

import (
	"github.com/zubairhamed/gossamer"
	"github.com/zubairhamed/gossamer/client"
	"log"
	"time"
)

func main() {

	c := client.NewClient("http://localhost:8000")

	obs := &gossamer.ObservationEntity{}
	obs.PhenomenonTime = gossamer.NewTimePeriod(time.Now(), time.Now())
	obs.Result = "123"
	obs.ResultTime = gossamer.NewTimeInstant(time.Now())
	ds := &gossamer.DatastreamEntity{}
	ds.Id = "Datastream-1"
	obs.Datastream = ds

	err := c.InsertObservation(obs)
	if err != nil {
		log.Println(err)
	}
}
