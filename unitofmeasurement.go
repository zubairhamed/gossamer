package gossamer

const (
	UNIT_VOLTAGE = "Voltage"
	UNIT_CELSIUS = "Celsius"
)

type UnitOfMeasurement interface {
	GetName() string
	GetSymbol() string
	GetDefinition() string
}

type GossamerUnitOfMeasurement struct {
	name       string
	symbol     string
	definition string
}

func (u *GossamerUnitOfMeasurement) GetName() string {
	return u.name
}

func (u *GossamerUnitOfMeasurement) GetSymbol() string {
	return u.symbol
}

func (u *GossamerUnitOfMeasurement) GetDefinition() string {
	return u.definition
}

func NewUnitOfMeasurementBySymbol(symbol string) UnitOfMeasurement {
	switch symbol {
	case UNIT_VOLTAGE:
		return &GossamerUnitOfMeasurement{
			name:       "Voltage",
			symbol:     "V",
			definition: "http://dbpedia.org/page/Voltage",
		}

	case UNIT_CELSIUS:
		return &GossamerUnitOfMeasurement{
			name:       "Celsius",
			symbol:     "°C",
			definition: "http://dbpedia.org/page/Celsius",
		}
	}
	return nil
}
