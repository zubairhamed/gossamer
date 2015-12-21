package main

import (
	"github.com/zubairhamed/gossamer/client"
	"log"
	"github.com/zubairhamed/gossamer"
	"time"
)

func main() {

	c := client.NewClient("http://localhost:8000")

	obs := &gossamer.ObservationEntity{}
	obs.PhenomenonTime = gossamer.NewTimePeriod(time.Now(), time.Now())
	obs.Result = "123"
	obs.ResultTime = gossamer.NewTimeInstant(time.Now())

	err := c.InsertObservation(obs)
	if err != nil {
		log.Println(err)
	}
}
