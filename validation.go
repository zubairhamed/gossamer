package gossamer

import "errors"

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

	case ENTITY_FEATURESOFINTEREST:
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

// TODO: Refactor this into whitelists
func ValidatePostRequestUrl(req Request) error {
	rp := req.GetResourcePath()
	last := rp.Last()
	lastEntity := last.GetEntity()
	cont := rp.Containing()

	if cont == nil {
		if last.GetEntity() == ENTITY_HISTORICALLOCATIONS {
			return errors.New("Not allowed")
		}
		return nil
	}

	contEntity := cont.GetEntity()
	if contEntity == ENTITY_THINGS && lastEntity == ENTITY_LOCATIONS ||
	   contEntity == ENTITY_FEATURESOFINTEREST && lastEntity == ENTITY_DATASTREAMS ||
	   contEntity == ENTITY_DATASTREAMS && lastEntity == ENTITY_OBSERVATIONS {

		return nil
	}
	return errors.New("Not allowed")
}

func ValidateGetRequestUrl(req Request) error {
	return nil
}

func ValidatePutRequestUrl(req Request) error {
	return nil
}

func ValidateDeleteRequestUrl(req Request) error {
	return nil
}

func ValidatePatchRequestUrl(req Request) error {
	return nil
}
