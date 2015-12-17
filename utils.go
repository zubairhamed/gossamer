package gossamer

import (
	"errors"
	"reflect"
	"log"
)

func ResolveEntityLink(id string, ent EntityType) string {
	s := string(ent)

	if id != "" {
		s += "(" + id + ")"
	}
	return s
}

func ResolveSelfLinkUrl(id string, ent EntityType) string {
	return "http://" + GLOB_ENV_HOST + "/v1.0/" + ResolveEntityLink(id, ent)
}

func ValidateMandatoryProperties(e SensorThing) error {
	elem := reflect.TypeOf(e).Elem().Name()

	switch elem {
	case "ThingEntity":
		ent := e.(*ThingEntity)
		if ent.Description == "" {
			return errors.New("Missing mandatory property for Thing entity: 'description'")
		}

	case "LocationEntity":
		ent := e.(*LocationEntity)
		if ent.Description == "" {
			return errors.New("Missing mandatory property for Location entity: 'description'")
		}

		if ent.EncodingType == "" {
			return errors.New("Missing mandatory property for Location entity: 'encodingType'")
		}

		if ent.Location == nil {
			return errors.New("Missing mandatory property for Location entity: 'location'")
		}

	case "ObservedPropertyEntity":
		ent := e.(*ObservedPropertyEntity)
		if ent.Name == "" {
			return errors.New("Missing mandatory property for ObservedProperty entity: 'name'")
		}

		if ent.Definition == "" {
			return errors.New("Missing mandatory property for ObservedProperty entity: 'definition'")
		}

		if ent.Description == "" {
			return errors.New("Missing mandatory property for ObservedProperty entity: 'description'")
		}

	case "DatastreamEntity":
		ent := e.(*DatastreamEntity)
		if ent.Description == "" {
			return errors.New("Missing mandatory property for Datastream entity: 'description'")
		}

		if ent.UnitOfMeasurement == nil {
			return errors.New("Missing mandatory property for Datastream entity: 'unitOfMeasurement'")
		}

		if ent.ObservationType == "" {
			return errors.New("Missing mandatory property for Datastream entity: 'observationType'")
		}

	case "SensorEntity":
		ent := e.(*SensorEntity)
		if ent.Description == "" {
			return errors.New("Missing mandatory property for Sensor entity: 'description'")
		}

		if ent.EncodingType == "" {
			return errors.New("Missing mandatory property for Sensor entity: 'encodingType'")
		}

		if ent.Metadata == "" {
			return errors.New("Missing mandatory property for Sensor entity: 'metadata'")
		}

	case "ObservationEntity":
		ent := e.(*ObservationEntity)
		log.Println("ent", ent.PhenomenonTime.IsZero())
		if ent.PhenomenonTime.IsZero() {
			return errors.New("Missing mandatory property for Observation entity: 'phenomenonTime'")
		}

		if ent.ResultTime.IsZero() {
			return errors.New("Missing mandatory property for Observation entity: 'resultTime'")
		}

		if ent.Result == nil {
			return errors.New("Missing mandatory property for Observation entity: 'result'")
		}

	case "FeatureOfInterestEntity":
		ent := e.(*FeatureOfInterestEntity)
		if ent.Description == "" {
			return errors.New("Missing mandatory property for FeaturesOfInterest entity: 'description'")
		}

		if ent.EncodingType == "" {
			return errors.New("Missing mandatory property for FeaturesOfInterest entity: 'encodingType'")
		}

		if ent.Feature == nil {
			return errors.New("Missing mandatory property for FeaturesOfInterest entity: 'feature'")
		}

	case "HistoricalLocationEntity":
		ent := e.(*HistoricalLocationEntity)
		if ent.Time.IsZero() {
			return errors.New("Missing mandatory property for HistoricalLocation entity: 'time'")
		}
	}
	return nil
}

func ValidateDeep(e SensorThing) error {
	return nil
}

func ValidateIntegrityConstraints(e SensorThing) error {
	elem := reflect.TypeOf(e).Elem().Name()
	switch elem {
	case "DatastreamEntity":
		ent := e.(*DatastreamEntity)

		if ent.IdThing == "" && ent.Thing == nil {
			return errors.New("Missing constrains for Datastream Entity: 'Thing'")
		}

		if ent.IdSensor == "" && ent.Sensor == nil {
			return errors.New("Missing constrains for Datastream Entity: 'Sensor'")
		}

		if ent.IdObservedProperty == "" && ent.ObservedProperty == nil {
			return errors.New("Missing constrains for Datastream Entity: 'ObservedProperty'")
		}

	case "ObservationEntity":
		ent := e.(*ObservationEntity)
		if ent.IdDatastream == "" && &ent.Datastream == nil {
			return errors.New("Missing constrains for Observation Entity: 'Datastream'")
		}

	}
	return nil
}
