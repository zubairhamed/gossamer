package gossamer
import (
	"strings"
	"log"
	"errors"
	"strconv"
)

func IsQueryOption(s string) bool {
	if 	strings.HasPrefix(s, "$expand") || strings.HasPrefix(s, "$select") ||  strings.HasPrefix(s, "$orderby") ||
		strings.HasPrefix(s, "$top") || strings.HasPrefix(s, "$skip") || strings.HasPrefix(s, "$count") ||
		strings.HasPrefix(s, "$filter") {
		return true
	}
	return false
}

func DiscoverQueryOptionType(s string) QueryOptionType {
	switch {
	case strings.HasPrefix(s, "$expand"):
		return QUERYOPT_EXPAND

	case strings.HasPrefix(s, "$select"):
		return QUERYOPT_SELECT

	case strings.HasPrefix(s, "$orderby"):
		return QUERYOPT_ORDERBY

	case strings.HasPrefix(s, "$top"):
		return QUERYOPT_TOP

	case strings.HasPrefix(s, "$skip"):
		return QUERYOPT_SKIP

	case strings.HasPrefix(s, "$count"):
		return QUERYOPT_COUNT

	case strings.HasPrefix(s, "$filter"):
		return QUERYOPT_FILTER

	default:
		return QUERYOPT_UNKNOWN
	}
}

func CreateQueryOptions(q string) (QueryOptions, error) {
	opts := &DefaultQueryOption{}
	if q == "" {
		return opts, ERR_QUERYOPTION_BLANK
	}
	optsSplit := strings.Split(q, "&")
	for _, optItem := range optsSplit {
		firstEq := strings.Index(optItem, "=")
		opt := optItem[0:firstEq]
		if IsQueryOption(opt) {
			optionType := DiscoverQueryOptionType(opt)
			optValue, err := CreateQueryOption(optionType, optItem[firstEq+1:])
			if err != nil {
				return nil, ERR_QUERYOPTION_INVALID_VALUE
			}
			opts.Set(optionType, optValue)
		} else {
			log.Println("ERROR: Not a query option - ", opt)
		}
	}
	return opts, nil
}

func CreateQueryOption(o QueryOptionType, s string) (QueryOption, error) {
	switch o {
	case QUERYOPT_EXPAND:
		return CreateExpandOption(s)

	case QUERYOPT_TOP:
		return CreateTopOption(s)

	case QUERYOPT_COUNT:
		return CreateCountOption(s)

	case QUERYOPT_FILTER:
		return CreateFilterOption(s)

	case QUERYOPT_SKIP:
		return CreateSkipOption(s)

	case QUERYOPT_ORDERBY:
		return CreateOrderByOption(s)

	case QUERYOPT_SELECT:
		return CreateSelectOption(s)
	}
	return nil, errors.New("Unknown Option types not allowed")
}

// Creates an Expand QueryOption given a string value
func CreateExpandOption (s string) (ExpandOption, error) {
	splitValues := strings.Split(s, ",")

	return &DefaultExpandOption{
		values: splitValues,
	}, nil
}

func CreateTopOption (s string) (TopOption, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}

	return &DefaultTopOption{
		value: i,
	}, nil
}

func CreateCountOption (s string) (CountOption, error) {
	i, err := strconv.ParseBool(s)
	if err != nil {
		return nil, err
	}

	return &DefaultCountOption{
		value: i,
	}, nil
}

func CreateFilterOption (s string) (FilterOption, error) {
	return &DefaultFilterOption{}, nil
}

func CreateSkipOption(s string) (SkipOption, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}

	return &DefaultSkipOption{
		value: i,
	}, nil
}

func CreateOrderByOption(s string) (OrderByOption, error) {
	return &DefaultOrderByOption{}, nil
}

func CreateSelectOption(s string) (SelectOption, error) {
	splitValues := strings.Split(s, ",")

	return &DefaultSelectOption{
		values: splitValues,
	}, nil
}

type DefaultExpandOption struct {
	values 	[]string
}

func (o *DefaultExpandOption) GetType() QueryOptionType {
	return QUERYOPT_EXPAND
}

func (o *DefaultExpandOption) GetValue() []string {
	return o.values
}

type DefaultSelectOption struct {
	values []string
}

func (o *DefaultSelectOption) GetType() QueryOptionType {
	return QUERYOPT_SELECT
}

func (o *DefaultSelectOption) GetValue() []string {
	return o.values
}

type DefaultOrderByOption struct {
	values 	[]OrderByOptionValue
}

func (o *DefaultOrderByOption) GetType() QueryOptionType {
	return QUERYOPT_ORDERBY
}

func (o *DefaultOrderByOption) GetValue() []OrderByOptionValue {
	return o.values
}

type DefaultOrderByOptionValue struct {

}

type DefaultTopOption struct {
	value 	int
}

func (o *DefaultTopOption) GetValue() int {
	return o.value
}

func (o *DefaultTopOption) GetType() QueryOptionType {
	return QUERYOPT_TOP
}

type DefaultSkipOption struct {
	value 	int
}

func (o *DefaultSkipOption) GetType() QueryOptionType {
	return QUERYOPT_SKIP
}

func (o *DefaultSkipOption) GetValue() int {
	return o.value
}

type DefaultCountOption struct {
	value 	bool
}

func (o *DefaultCountOption) GetType() QueryOptionType {
	return QUERYOPT_COUNT
}

func (o *DefaultCountOption) GetValue() bool {
	return o.value
}

type DefaultFilterOption struct {

}

func (o *DefaultFilterOption) GetType() QueryOptionType {
	return QUERYOPT_FILTER
}


