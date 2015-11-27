package gossamer
import "time"

type EncodingType string
const (
	ENCODINGTYPE_PDF		EncodingType = "application/pdf"
	ENCODINGTYPE_SENSORML	EncodingType = "http://www.opengis.net/doc/IS/SensorML/2.0"
	ENCODINGTYPE_GEO		EncodingType = "application/vnd.geo+json"
)

type ObservationType string

/*
{
	"@iot.id": 1,
	"@iot.selfLink": "http://example.org/v1.0/Things(1)",
	"Locations@iot.navigationLink": "Things(1)/Locations",
	"Datastreams@iot.navigationLink": "Things(1)/Datastreams",
	"HistoricalLocations@iot.navigationLink": "Things(1)/HistoricalLocations",
	"description": "This thing is an oven.",
	"properties": {
		"owner": "John Doe",
		"color": "Silver"
	}
}
*/
type Thing struct {
	description 		string
	properties			map[string]string
	// location 			*Location
	// historicalLocations	[]*HistoricalLocation
	// datastreams			[]*Datastream
}

/*
{
"@iot.id": 1,
"@iot.selfLink": "http://example.org/v1.0/Locations(1)", "Things@iot.navigationLink": "Locations(1)/Things", "HistoricalLocations@iot.navigationLink": "Locations(1)/HistoricalLocations", "encodingType": "application/vnd.geo+json",
"location": {
"type": "Point",
"coordinates": [-114.06,51.05] }
}
 */
type Location struct {
	description 		string
	encodingType 		EncodingType
	// location 		??
	// things 				[]*Thing
	// HistoricalLocation	[]*HistoricalLocation
}

/*
{
	"@iot.id": 1,
	"@iot.selfLink": "http://example.org/v1.0/Locations(1)",
	 "Things@iot.navigationLink": "Locations(1)/Things",
	 "HistoricalLocations@iot.navigationLink": "Locations(1)/HistoricalLocations",
	 "encodingType": "application/vnd.geo+json",
	"location": {
		"type": "Point",
		"coordinates": [-114.06,51.05]
	}
}
 */
type HistoricalLocation struct {
	time 		time.Time
	// thing 		*Thing
	// locations 	[]*Location
}

/*
{
	"@iot.id": 1,
	"@iot.selfLink": "http://example.org/v1.0/Datastreams(1)",
	"Thing@iot.navigationLink": "HistoricalLocations(1)/Thing",
	"Sensor@iot.navigationLink": "Datastreams(1)/Sensor",
	"ObservedProperty@iot.navigationLink": "Datastreams(1)/ObservedProperty",
	"Observations@iot.navigationLink": "Datastreams(1)/Observations",
	"description": "This is a datastream measuring the temperature in an oven.",
	"unitOfMeasurement": {
		"name": "degree Celsius",
		"symbol": "°C",
		"definition": "http://unitsofmeasure.org/ucum.html#para-30"
	},
	"observationType": "http://www.opengis.net/def/observationType/OGC- OM/2.0/OM_Measurement",
	"observedArea": {
		"type": "Polygon",
		"coordinates": [[[100,0],[101,0],[101,1],[100,1],[100,0]]]
	},
	"phenomenonTime": "2014-03-01T13:00:00Z/2015-05-11T15:30:00Z",
	"resultTime": "2014-03-01T13:00:00Z/2015-05-11T15:30:00Z"
}
 */

type UCUM struct {

}

type Datastream struct {
	description 		string
	unitOfMeasurement	*UCUM
	observationType 	ObservationType
	// observedArea		??
	// phenomenonTime 	??
	// resultTime 			??
	// thing 	*Thing
	// sensor 	?
	// ObservedProperty
	// Observation
}

/*
{
"@iot.id": 1,
"@iot.selfLink": "http://example.org/v1.0/Sensors(1)", "Datastreams@iot.navigationLink": "Sensors(1)/Datastreams", "description": "TMP36 - Analog Temperature sensor", "encodingType": "application/pdf",
"metadata": "http://example.org/TMP35_36_37.pdf"
}
 */
type Sensor struct {
	description 	string
	encodingType 	EncodingType
	metadata 		string
	datastream 		[]*Datastream
}

/*
{
	"@iot.id": 1,
	"@iot.selfLink": "http://example.org/v1.0/ObservedProperties(1)",
	"Datastreams@iot.navigationLink": "ObservedProperties(1)/Datastreams",
	"description": "The dewpoint temperature is the temperature to which the air must be cooled, at constant pressure, for dew to form. As the grass and other objects near the ground cool to the dewpoint, some of the water vapor in the atmosphere condenses into liquid water on the objects.",
	"name": "DewPoint Temperature",
	"definition": "http://dbpedia.org/page/Dew_point"
}
 */
type ObservedProperty struct {
	name 			string
	definition 		string
	description		string
	// Datastream
}

/*
{
	"@iot.id": 1,
	"@iot.selfLink": "http://example.org/v1.0/Observations(1)",
	"FeatureOfInterest@iot.navigationLink": "Observations(1)/FeatureOfInterest",
	"Datastream@iot.navigationLink":"Observations(1)/Datastream",
	"phenomenonTime": "2014-12-31T11:59:59.00+08:00",
	"resultTime": "2014-12-31T11:59:59.00+08:00",
	"result": 70.4
}￼￼￼
 */
type Observation struct {

}

/*
{
	"@iot.id": 1,
	"@iot.selfLink": "http://example.org/v1.0/FeaturesOfInterest(1)",
	"Observations@iot.navigationLink": "FeaturesOfInterest(1)/Observations",
	"description": "This is a weather station.",
	"encodingType": "application/vnd.geo+json",
	"feature": {
		"type": "Point",
		"coordinates": [-114.06,51.05] }
	}
 */
type FeatureOfInterest struct {

}