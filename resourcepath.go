package gossamer

type SensorThingsResourcePathItem struct {
	entityType   EntityType
	entityId     string
	queryOptions QueryOptions
}

func (n *SensorThingsResourcePathItem) GetQueryOptions() QueryOptions {
	return n.queryOptions
}

func (n *SensorThingsResourcePathItem) GetEntity() EntityType {
	return n.entityType
}

func (n *SensorThingsResourcePathItem) GetId() string {
	return n.entityId
}

type SensorThingsResourcePath struct {
	currIndex     int
	items         []ResourcePathItem
	property      string
	propertyValue string
}

func (r *SensorThingsResourcePath) Next() ResourcePathItem {
	r.currIndex++
	return r.At(r.currIndex)
}

func (r *SensorThingsResourcePath) Prev() ResourcePathItem {
	r.currIndex--
	return r.At(r.currIndex)
}

func (r *SensorThingsResourcePath) Current() ResourcePathItem {
	return r.At(r.currIndex)
}

func (r *SensorThingsResourcePath) Containing() ResourcePathItem {
	l := len(r.items)
	if l > 1 {
		return r.At(l - 2)
	}
	return nil
}

func (r *SensorThingsResourcePath) First() ResourcePathItem {
	r.currIndex = 0
	return r.At(r.currIndex)
}

func (r *SensorThingsResourcePath) Last() ResourcePathItem {
	r.currIndex = r.Size() - 1
	return r.At(r.currIndex)
}

func (r *SensorThingsResourcePath) HasNext() bool {
	if r.IsLast() {
		return false
	}
	return true
}

func (r *SensorThingsResourcePath) IsLast() bool {
	sz := r.Size()
	if r.CurrentIndex() == sz-1 {
		return true
	}
	return false
}

func (r *SensorThingsResourcePath) IsFirst() bool {
	if r.CurrentIndex() == 0 {
		return true
	}
	return false
}

func (r *SensorThingsResourcePath) CurrentIndex() int {
	return r.currIndex
}

func (r *SensorThingsResourcePath) Size() int {
	return len(r.items)
}

func (r *SensorThingsResourcePath) Add(rsrc ResourcePathItem) {
	r.items = append(r.items, rsrc)
}

func (r *SensorThingsResourcePath) At(idx int) ResourcePathItem {
	sz := r.Size() - 1
	if idx > sz || idx < 0 {
		return nil
	}
	return r.items[idx]
}

func (r *SensorThingsResourcePath) All() []ResourcePathItem {
	return r.items
}
