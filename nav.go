package gossamer

type DefaultNavigation struct {
	property 		string
	propertyValue 	string
	items 			[]NavigationItem
}

func (n *DefaultNavigation) GetItems() []NavigationItem {
	return n.items
}

type DefaultNavigationItem struct {
	entityType 		EntityType
	entityId 		string
	queryOptions	QueryOptions
}

func (n *DefaultNavigationItem) GetQueryOptions() QueryOptions {
	return n.queryOptions
}

func (n *DefaultNavigationItem) GetEntity() EntityType {
	return n.entityType
}

func (n *DefaultNavigationItem) GetId() string {
	return n.entityId
}


