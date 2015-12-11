package gossamer

import (
	"log"
	"net/url"
	"strings"
)

func CreateRequest(url *url.URL) (Request, error) {
	rp := &SensorThingsResourcePath{
		currIndex: -1,
		items:     []ResourcePathItem{},
	}

	path := url.Path
	pathSplit := strings.Split(path, "/")[2:]
	items := []ResourcePathItem{}
	pathSplitItems := len(pathSplit)

	for idx, val := range pathSplit {
		if IsEntity(val) {
			navItem := &SensorThingsResourcePathItem{}
			entityType := DiscoverEntityType(val)
			navItem.entityType = entityType

			br1Index := strings.Index(val, "(")
			br2Index := strings.Index(val, ")")

			if br1Index != -1 && br2Index != -1 {
				parenthesisValue := val[br1Index+1 : br2Index]

				// Query Option
				if strings.HasPrefix(parenthesisValue, "$") {
					navItem.queryOptions, _ = CreateQueryOptions(parenthesisValue)
				} else {
					navItem.entityId = parenthesisValue
				}
			}
			items = append(items, navItem)
		} else {
			if strings.HasPrefix(val, "$") && idx == pathSplitItems-1 {
				rp.property = val
			} else {
				if idx == pathSplitItems-1 || idx == pathSplitItems-2 {
					rp.propertyValue = val
				} else {
					return nil, ERR_INVALID_ENTITY
				}
			}
		}
	}

	rp.items = items

	queryOpts, _ := CreateQueryOptions(url.RawQuery)
	req := &DefaultRequest{
		resourcePath: rp,
		queryOptions: queryOpts,
	}
	return req, nil
}

type DefaultRequest struct {
	resourcePath ResourcePath
	queryOptions QueryOptions
}

func (r *DefaultRequest) GetProtocol() ProtocolType     { return 0 }
func (r *DefaultRequest) GetQueryOptions() QueryOptions { return nil }

func (r *DefaultRequest) GetResourcePath() ResourcePath {
	return r.resourcePath
}

type DefaultQueryOption struct {
	expandOption  QueryOption
	selectOption  QueryOption
	orderByOption QueryOption
	topOption     QueryOption
	skipOption    QueryOption
	countOption   QueryOption
	filterOption  QueryOption
}

func (o *DefaultQueryOption) Set(optType QueryOptionType, value QueryOption) {
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

func (o *DefaultQueryOption) ExpandSet() bool {
	return o.expandOption != nil
}

func (o *DefaultQueryOption) SelectSet() bool {
	return o.selectOption != nil
}

func (o *DefaultQueryOption) OrderBySet() bool {
	return o.orderByOption != nil
}

func (o *DefaultQueryOption) TopSet() bool {
	return o.topOption != nil
}

func (o *DefaultQueryOption) SkipSet() bool {
	return o.skipOption != nil
}

func (o *DefaultQueryOption) CountSet() bool {
	return o.countOption != nil
}

func (o *DefaultQueryOption) FilterSet() bool {
	return o.filterOption != nil
}

func (o *DefaultQueryOption) GetExpandOption() ExpandOption {
	return o.expandOption.(ExpandOption)
}
func (o *DefaultQueryOption) GetSelectOption() SelectOption {
	return o.selectOption.(SelectOption)
}

func (o *DefaultQueryOption) GetOrderByOption() OrderByOption {
	return o.orderByOption.(OrderByOption)
}

func (o *DefaultQueryOption) GetTopOption() TopOption {
	return o.topOption.(TopOption)
}

func (o *DefaultQueryOption) GetSkipOption() SkipOption {
	return o.skipOption.(SkipOption)
}

func (o *DefaultQueryOption) GetCountOption() CountOption {
	return o.countOption.(CountOption)
}

func (o *DefaultQueryOption) GetFilterOption() FilterOption {
	return o.filterOption.(FilterOption)
}
