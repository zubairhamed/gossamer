package gossamer
import (
	"net/url"
	"strings"
	"log"
)

func IsQueryOption(s string) bool {
	if 	strings.HasPrefix(s, "$expand") ||
		strings.HasPrefix(s, "$select") ||
		strings.HasPrefix(s, "$orderby") ||
		strings.HasPrefix(s, "$top") ||
		strings.HasPrefix(s, "$skip") ||
		strings.HasPrefix(s, "$count") ||
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

func CreateQueryOptions(q string) QueryOptions {
	opts := &DefaultQueryOption{}
	if q == "" {
		return opts
	}
	optsSplit := strings.Split(q, "&")
	for _, val := range optsSplit {
		if IsQueryOption(val) {
			optionType := DiscoverQueryOptionType(val)

			opts.Set(optionType, "")
		} else {
			log.Println("ERROR: Not a query option - ", val)
		}
	}

	return opts
}

func CreateRequest(url *url.URL) (Request, error) {
	nav := &DefaultNavigation{}

	path := url.Path
	pathSplit := strings.Split(path, "/")[2:]
	navItems := []NavigationItem{}
	pathSplitItems := len(pathSplit)

	for idx, val := range pathSplit {
		if IsEntity(val) {
			navItem := &DefaultNavigationItem{}
			entityType := DiscoverEntityType(val)
			navItem.entityType = entityType

			br1Index := strings.Index(val, "(")
			br2Index := strings.Index(val, ")")

			if br1Index != -1 && br2Index != -1 {
				parenthesisValue := val[br1Index+1: br2Index]

				// Query Option
				if strings.HasPrefix(parenthesisValue, "$") {
					navItem.queryOptions = CreateQueryOptions(parenthesisValue)
				} else {
					navItem.entityId = parenthesisValue
				}
			}
			navItems = append(navItems, navItem)
		} else {

			if strings.HasPrefix(val, "$") && idx == pathSplitItems -1 {
				nav.property = val
			} else {
				if idx == pathSplitItems -1 || idx == pathSplitItems -2 {
					nav.propertyValue = val
				} else {
					return nil, ERR_INVALID_ENTITY
				}
			}
		}
	}

	nav.items = navItems

	req := &DefaultRequest{
		navigation: nav,
		queryOptions: CreateQueryOptions(url.RawQuery),
	}
	return req, nil
}

type DefaultRequest struct {
	navigation 		Navigation
	queryOptions 	QueryOptions
}

func (r *DefaultRequest) GetProtocol() ProtocolType { return 0 }
func (r *DefaultRequest) GetQueryOptions() QueryOptions { return nil }
func (r *DefaultRequest) GetNavigation() Navigation {
	return r.navigation
}

type DefaultQueryOption struct {
	expandValue 	string
	selectValue		string
	orderByValue	string
	topValue		string
	skipValue		string
	countValue 		string
	filterValue		string
}

func (o *DefaultQueryOption) Set(optType QueryOptionType, value string)  {

}

func (o *DefaultQueryOption) ExpandSet() bool {
	return o.expandValue != ""
}

func (o *DefaultQueryOption) SelectSet() bool {
	return o.selectValue != ""
}

func (o *DefaultQueryOption) OrderBySet() bool {
	return o.orderByValue != ""
}

func (o *DefaultQueryOption) TopSet() bool {
	return o.topValue != ""
}

func (o *DefaultQueryOption) SkipSet() bool {
	return o.skipValue != ""
}

func (o *DefaultQueryOption) CountSet() bool {
	return o.countValue != ""
}

func (o *DefaultQueryOption) FilterSet() bool {
	return o.filterValue != ""
}