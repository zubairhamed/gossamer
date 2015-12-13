package gossamer

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

func IsQueryOption(s string) bool {
	if strings.HasPrefix(s, "$expand") || strings.HasPrefix(s, "$select") || strings.HasPrefix(s, "$orderby") ||
		strings.HasPrefix(s, "$top") || strings.HasPrefix(s, "$skip") || strings.HasPrefix(s, "$count") ||
		strings.HasPrefix(s, "$filter") {
		return true
	}
	return false
}

func DiscoverQueryOptionType(s string) QueryOptionType {
	switch {
	case strings.HasPrefix(s, string(QUERYOPT_EXPAND)):
		return QUERYOPT_EXPAND

	case strings.HasPrefix(s, string(QUERYOPT_SELECT)):
		return QUERYOPT_SELECT

	case strings.HasPrefix(s, string(QUERYOPT_ORDERBY)):
		return QUERYOPT_ORDERBY

	case strings.HasPrefix(s, string(QUERYOPT_TOP)):
		return QUERYOPT_TOP

	case strings.HasPrefix(s, string(QUERYOPT_SKIP)):
		return QUERYOPT_SKIP

	case strings.HasPrefix(s, string(QUERYOPT_COUNT)):
		return QUERYOPT_COUNT

	case strings.HasPrefix(s, string(QUERYOPT_FILTER)):
		return QUERYOPT_FILTER

	default:
		return QUERYOPT_UNKNOWN
	}
}

func CreateQueryOptions(q string) (QueryOptions, error) {
	opts := &GossamerQueryOption{}
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
func CreateExpandOption(s string) (ExpandOption, error) {
	splitValues := strings.Split(s, ",")

	return &GossamerExpandOption{
		values: splitValues,
	}, nil
}

func CreateTopOption(s string) (TopOption, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}

	return &GossamerTopOption{
		value: i,
	}, nil
}

func CreateCountOption(s string) (CountOption, error) {
	i, err := strconv.ParseBool(s)
	if err != nil {
		return nil, err
	}

	return &GossamerCountOption{
		value: i,
	}, nil
}

func CreateFilterOption(s string) (FilterOption, error) {
	return &GossamerFilterOption{}, nil
}

func CreateSkipOption(s string) (SkipOption, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}

	return &GossamerSkipOption{
		value: i,
	}, nil
}

func CreateOrderByOption(s string) (OrderByOption, error) {
	vals := strings.Split(s, ",")
	orderByVals := []OrderByOptionValue{}

	for _, v := range vals {
		orderByVals = append(orderByVals, GossamerOrderByOptionValue{
			s: v,
		})
	}

	return &GossamerOrderByOption{
		values: orderByVals,
	}, nil
}

func CreateSelectOption(s string) (SelectOption, error) {
	splitValues := strings.Split(s, ",")

	return &GossamerSelectOption{
		values: splitValues,
	}, nil
}

type GossamerExpandOption struct {
	values []string
}

func (o *GossamerExpandOption) GetType() QueryOptionType {
	return QUERYOPT_EXPAND
}

func (o *GossamerExpandOption) GetValue() []string {
	return o.values
}

type GossamerSelectOption struct {
	values []string
}

func (o *GossamerSelectOption) GetType() QueryOptionType {
	return QUERYOPT_SELECT
}

func (o *GossamerSelectOption) GetValue() []string {
	return o.values
}

type GossamerOrderByOption struct {
	values []OrderByOptionValue
}

func (o *GossamerOrderByOption) GetType() QueryOptionType {
	return QUERYOPT_ORDERBY
}

func (o *GossamerOrderByOption) GetValue() []OrderByOptionValue {
	return o.values
}

func (o *GossamerOrderByOption) GetSortProperties() []string {
	vals := []string{}
	for _, v := range o.values {
		vals = append(vals, v.GetSortProperty())
	}
	return vals
}

type GossamerOrderByOptionValue struct {
	s 	string
}

func (o GossamerOrderByOptionValue) GetSortProperty() string {
	return o.s
}



type GossamerTopOption struct {
	value int
}

func (o *GossamerTopOption) GetValue() int {
	return o.value
}

func (o *GossamerTopOption) GetType() QueryOptionType {
	return QUERYOPT_TOP
}

type GossamerSkipOption struct {
	value int
}

func (o *GossamerSkipOption) GetType() QueryOptionType {
	return QUERYOPT_SKIP
}

func (o *GossamerSkipOption) GetValue() int {
	return o.value
}

type GossamerCountOption struct {
	value bool
}

func (o *GossamerCountOption) GetType() QueryOptionType {
	return QUERYOPT_COUNT
}

func (o *GossamerCountOption) GetValue() bool {
	return o.value
}

type GossamerFilterOption struct {
}

func (o *GossamerFilterOption) GetType() QueryOptionType {
	return QUERYOPT_FILTER
}

type GossamerQueryOption struct {
	expandOption  QueryOption
	selectOption  QueryOption
	orderByOption QueryOption
	topOption     QueryOption
	skipOption    QueryOption
	countOption   QueryOption
	filterOption  QueryOption
}

func (o *GossamerQueryOption) Set(optType QueryOptionType, value QueryOption) {
	switch optType {
	case QUERYOPT_EXPAND:
		o.expandOption = value

	case QUERYOPT_COUNT:
		o.countOption = value

	case QUERYOPT_FILTER:
		o.filterOption = value

	case QUERYOPT_TOP:
		o.topOption = value

	case QUERYOPT_SKIP:
		o.skipOption = value

	case QUERYOPT_ORDERBY:
		o.orderByOption = value

	case QUERYOPT_SELECT:
		o.selectOption = value

	case QUERYOPT_UNKNOWN:
		log.Println("Attempting to set unknown Query Option")
		return
	}
}

func (o *GossamerQueryOption) ExpandSet() bool {
	return o.expandOption != nil
}

func (o *GossamerQueryOption) SelectSet() bool {
	return o.selectOption != nil
}

func (o *GossamerQueryOption) OrderBySet() bool {
	return o.orderByOption != nil
}

func (o *GossamerQueryOption) TopSet() bool {
	return o.topOption != nil
}

func (o *GossamerQueryOption) SkipSet() bool {
	return o.skipOption != nil
}

func (o *GossamerQueryOption) CountSet() bool {
	return o.countOption != nil
}

func (o *GossamerQueryOption) FilterSet() bool {
	return o.filterOption != nil
}

func (o *GossamerQueryOption) GetExpandOption() ExpandOption {
	return o.expandOption.(ExpandOption)
}
func (o *GossamerQueryOption) GetSelectOption() SelectOption {
	return o.selectOption.(SelectOption)
}

func (o *GossamerQueryOption) GetOrderByOption() OrderByOption {
	return o.orderByOption.(OrderByOption)
}

func (o *GossamerQueryOption) GetTopOption() TopOption {
	return o.topOption.(TopOption)
}

func (o *GossamerQueryOption) GetSkipOption() SkipOption {
	return o.skipOption.(SkipOption)
}

func (o *GossamerQueryOption) GetCountOption() CountOption {
	return o.countOption.(CountOption)
}

func (o *GossamerQueryOption) GetFilterOption() FilterOption {
	return o.filterOption.(FilterOption)
}
