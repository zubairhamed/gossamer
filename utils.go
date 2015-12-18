package gossamer

import (
	"errors"
	"encoding/json"
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

func ValidateMandatoryProperties(entity SensorThing) error {
	switch entity.GetType() {
	case ENTITY_THINGS:
		e := (entity).(*ThingEntity)
		if e.Description == "" {
			return errors.New("Missing mandatory property for Thing entity: 'description'")
		}

	case ENTITY_OBSERVATIONS:
		e := (entity).(*ObservationEntity)
		if e.PhenomenonTime == nil {
			return errors.New("Missing mandatory property for Observation entity: 'phenomenonTime'")
		}

		if e.PhenomenonTime.IsZero() {
			return errors.New("Missing mandatory property for Observation entity: 'phenomenonTime'")
		}

		if e.ResultTime == nil {
			return errors.New("Missing mandatory property for Observation entity: 'resultTime'")
		}
		if e.ResultTime.IsZero() {
			return errors.New("Missing mandatory property for Observation entity: 'resultTime'")
		}

		if e.Result == nil {
			return errors.New("Missing mandatory property for Observation entity: 'result'")
		}

	case ENTITY_HISTORICALLOCATIONS:
		e := (entity).(*HistoricalLocationEntity)
		if e.Time.IsZero() {
			return errors.New("Missing mandatory property for HistoricalLocation entity: 'time'")
		}

	case ENTITY_SENSORS:
		e := (entity).(*SensorEntity)
		if e.Description == "" {
			return errors.New("Missing mandatory property for Sensor entity: 'description'")
		}

		if e.EncodingType == "" {
			return errors.New("Missing mandatory property for Sensor entity: 'encodingType'")
		}

		if e.Metadata == "" {
			return errors.New("Missing mandatory property for Sensor entity: 'metadata'")
		}

	case ENTITY_LOCATIONS:
		e := (entity).(*LocationEntity)
		if e.Description == "" {
			return errors.New("Missing mandatory property for Location entity: 'description'")
		}

		if e.EncodingType == "" {
			return errors.New("Missing mandatory property for Location entity: 'encodingType'")
		}

		if e.Location == nil {
			return errors.New("Missing mandatory property for Location entity: 'location'")
		}

	case ENTITY_FEATURESOFINTERESTS:
		e := (entity).(*FeatureOfInterestEntity)
		if e.Description == "" {
			return errors.New("Missing mandatory property for FeaturesOfInterest entity: 'description'")
		}

		if e.EncodingType == "" {
			return errors.New("Missing mandatory property for FeaturesOfInterest entity: 'encodingType'")
		}

		if e.Feature == nil {
			return errors.New("Missing mandatory property for FeaturesOfInterest entity: 'feature'")
		}

	case ENTITY_DATASTREAMS:
		e := (entity).(*DatastreamEntity)
		if e.Description == "" {
			return errors.New("Missing mandatory property for Datastream entity: 'description'")
		}

		if e.UnitOfMeasurement == nil {
			return errors.New("Missing mandatory property for Datastream entity: 'unitOfMeasurement'")
		}

		if e.ObservationType == "" {
			return errors.New("Missing mandatory property for Datastream entity: 'observationType'")
		}

	case ENTITY_OBSERVEDPROPERTIES:
		e := (entity).(*ObservedPropertyEntity)
		if e.Name == "" {
			return errors.New("Missing mandatory property for ObservedProperty entity: 'name'")
		}

		if e.Definition == "" {
			return errors.New("Missing mandatory property for ObservedProperty entity: 'definition'")
		}

		if e.Description == "" {
			return errors.New("Missing mandatory property for ObservedProperty entity: 'description'")
		}
	}
	return nil
}

func ValidateIntegrityConstraints(entity SensorThing) error {
	switch entity.GetType() {
	case ENTITY_OBSERVATIONS:
		e := (entity).(*ObservationEntity)
		if e.IdDatastream == "" && e.Datastream == nil {
			return errors.New("Missing constrains for Observation Entity: 'Datastream'")
		}

	case ENTITY_DATASTREAMS:
		e := (entity).(*DatastreamEntity)
		if e.IdThing == "" && e.Thing == nil {
			return errors.New("Missing constrains for Datastream Entity: 'Thing'")
		}

		if e.IdSensor == "" && e.Sensor == nil {
			return errors.New("Missing constrains for Datastream Entity: 'Sensor'")
		}

		if e.IdObservedProperty == "" && e.ObservedProperty == nil {
			return errors.New("Missing constrains for Datastream Entity: 'ObservedProperty'")
		}
	}
	return nil
}

func SetAssociatedEntityId(entity SensorThing, et EntityType, id string) {
	switch entity.GetType() {
	case ENTITY_THINGS:

	case ENTITY_OBSERVATIONS:
		e := (entity).(*ObservationEntity)
		if et == ENTITY_DATASTREAMS {
			e.Datastream = NewDatastreamEntity()
			e.Datastream.Id = id
		}

		if et == ENTITY_FEATURESOFINTERESTS {
			e.FeatureOfInterest = NewFeatureOfInterestEntity()
			e.FeatureOfInterest.Id = id
		}

	case ENTITY_HISTORICALLOCATIONS:
		e := (entity).(*HistoricalLocationEntity)
		log.Println(e)

	case ENTITY_SENSORS:
		e := (entity).(*SensorEntity)
		log.Println(e)

	case ENTITY_LOCATIONS:
		e := (entity).(*LocationEntity)
		log.Println(e)

	case ENTITY_FEATURESOFINTERESTS:
		e := (entity).(*FeatureOfInterestEntity)
		log.Println(e)

	case ENTITY_DATASTREAMS:
		e := (entity).(*DatastreamEntity)
		log.Println(e)

	case ENTITY_OBSERVEDPROPERTIES:
		e := (entity).(*ObservedPropertyEntity)
		log.Println(e)
	}
}

func DecodeJsonToEntityStruct(decoder *json.Decoder, et EntityType) (SensorThing, error) {
	var err error
	switch et {
	case ENTITY_THINGS:
		var e ThingEntity
		err = decoder.Decode(&e)
		return e, err

	case ENTITY_OBSERVATIONS:
		var e ObservationEntity
		err = decoder.Decode(&e)
		return e, err

	case ENTITY_HISTORICALLOCATIONS:
		return nil, errors.New("Adding Historical Locations not allowed")

	case ENTITY_SENSORS:
		var e SensorEntity
		err = decoder.Decode(&e)
		return e, err

	case ENTITY_LOCATION:
		var e LocationEntity
		err = decoder.Decode(&e)
		return e, err

	case ENTITY_FEATURESOFINTERESTS:
		var e FeatureOfInterestEntity
		err = decoder.Decode(&e)
		return e, err

	case ENTITY_DATASTREAMS:
		var e DatastreamEntity
		err = decoder.Decode(&e)
		return e, err

	case ENTITY_OBSERVEDPROPERTIES:
		var e ObservedPropertyEntity
		err = decoder.Decode(&e)
		return e, err
	}
	return nil, errors.New("Unknown Entity Type")
}
