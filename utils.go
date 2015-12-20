package gossamer

import (
	"encoding/json"
	"errors"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
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

func GetAssociatedEntityId(entity SensorThing, et EntityType) string {
	switch entity.GetType() {
	case ENTITY_THINGS:
		// e := (entity).(*ThingEntity)

	case ENTITY_OBSERVATIONS:
		// e := (entity).(*ObservationEntity)

	case ENTITY_HISTORICALLOCATIONS:
		// e := (entity).(*HistoricalLocationEntity)

	case ENTITY_SENSORS:
		// e := (entity).(*SensorEntity)

	case ENTITY_LOCATIONS:
		// e := (entity).(*LocationEntity)

	case ENTITY_FEATURESOFINTEREST:
		// e := (entity).(*FeatureOfInterestEntity)

	case ENTITY_DATASTREAMS:
		e := (entity).(*DatastreamEntity)
		switch {
		case et == ENTITY_SENSOR || et == ENTITY_SENSORS:
			return e.IdSensor

		case et == ENTITY_THING || et == ENTITY_THINGS:
			return e.IdThing

		case et == ENTITY_OBSERVEDPROPERTIES || et == ENTITY_OBSERVEDPROPERTY:
			return e.IdObservedProperty
		}

	case ENTITY_OBSERVEDPROPERTIES:
		// e := (entity).(*ObservedPropertyEntity)
	}
	return ""
}

func SetAssociatedEntityId(entity SensorThing, et EntityType, id string) {
	switch entity.GetType() {
	case ENTITY_THINGS:
		e := (entity).(*ThingEntity)
		if et == ENTITY_LOCATIONS {
			entity := NewLocationEntity()
			entity.Id = id
			e.Locations = []*LocationEntity{entity}
		}

		if et == ENTITY_DATASTREAMS {
			entity := NewDatastreamEntity()
			entity.Id = id
			e.Datastreams = []*DatastreamEntity{entity}
		}

	case ENTITY_OBSERVATIONS:
		e := (entity).(*ObservationEntity)
		if et == ENTITY_DATASTREAMS {
			e.Datastream = NewDatastreamEntity()
			e.Datastream.Id = id
		}

		if et == ENTITY_FEATURESOFINTEREST {
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

	case ENTITY_FEATURESOFINTEREST:
		e := (entity).(*FeatureOfInterestEntity)
		log.Println(e)

	case ENTITY_DATASTREAMS:
		e := (entity).(*DatastreamEntity)
		if et == ENTITY_THINGS {
			e.Thing = NewThingEntity()
			e.Thing.Id = id
		}

		if et == ENTITY_OBSERVEDPROPERTIES {
			e.ObservedProperty = NewObservedPropertyEntity()
			e.ObservedProperty.Id = id
		}

		if et == ENTITY_SENSORS {
			e.Sensor = NewSensorEntity()
			e.Sensor.Id = id
		}

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
		return &e, err

	case ENTITY_OBSERVATIONS:
		var e ObservationEntity
		err = decoder.Decode(&e)
		return &e, err

	case ENTITY_HISTORICALLOCATIONS:
		return nil, errors.New("Adding Historical Locations not allowed")

	case ENTITY_SENSORS:
		var e SensorEntity
		err = decoder.Decode(&e)
		return &e, err

	case ENTITY_LOCATIONS:
		var e LocationEntity
		err = decoder.Decode(&e)
		return &e, err

	case ENTITY_FEATURESOFINTEREST:
		var e FeatureOfInterestEntity
		err = decoder.Decode(&e)
		return &e, err

	case ENTITY_DATASTREAMS:
		var e DatastreamEntity
		err = decoder.Decode(&e)
		return &e, err

	case ENTITY_OBSERVEDPROPERTIES:
		var e ObservedPropertyEntity
		err = decoder.Decode(&e)
		return &e, err
	}
	return nil, errors.New("Unknown Entity Type")
}

func ThrowHttpCreated(msg string, w http.ResponseWriter) {
	http.Error(w, msg, http.StatusCreated)
}

func ThrowHttpBadRequest(msg string, w http.ResponseWriter) {
	http.Error(w, msg, http.StatusBadRequest)
}

func ThrowHttpInternalServerError(msg string, w http.ResponseWriter) {
	http.Error(w, msg, http.StatusInternalServerError)
}

func ThrowHttpMethodNotAllowed(msg string, w http.ResponseWriter) {
	http.Error(w, msg, http.StatusMethodNotAllowed)
}

func ThrowNotAcceptable(msg string, w http.ResponseWriter) {
	http.Error(w, msg, http.StatusNotAcceptable)
}

func GenerateEntityId() string {
	return uuid.NewV4().String()
}
