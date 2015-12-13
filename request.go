package gossamer

import (
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
	req := &GossamerRequest{
		resourcePath: rp,
		queryOptions: queryOpts,
	}
	return req, nil
}

type GossamerRequest struct {
	resourcePath ResourcePath
	queryOptions QueryOptions
}

func (r *GossamerRequest) GetProtocol() ProtocolType     { return 0 }
func (r *GossamerRequest) GetQueryOptions() QueryOptions { return nil }

func (r *GossamerRequest) GetResourcePath() ResourcePath {
	return r.resourcePath
}
