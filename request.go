package gossamer

import (
	"log"
	"net/url"
	"strings"
)

func CreateIncomingRequest(url *url.URL, t ProtocolType) (Request, error) {
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

	queryOpts, err := CreateQueryOptions(url.RawQuery)
	if err != nil {
		log.Println("Error creating option: ", err)
		return nil, err
	}
	req := &GossamerRequest{
		protocol:     t,
		resourcePath: rp,
		queryOptions: queryOpts,
	}
	return req, nil
}

type GossamerRequest struct {
	protocol     ProtocolType
	resourcePath ResourcePath
	queryOptions QueryOptions
}

func (r *GossamerRequest) GetProtocol() ProtocolType {
	return r.protocol
}

func (r *GossamerRequest) GetQueryOptions() QueryOptions {
	return r.queryOptions
}

func (r *GossamerRequest) GetResourcePath() ResourcePath {
	return r.resourcePath
}
