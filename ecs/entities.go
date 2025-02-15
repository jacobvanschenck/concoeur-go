package ecs

type Component interface {
	typeId() TypeId
}

type Entities struct {
	entityMap   []int32
	bitMasks    map[TypeId]int32
	components  map[TypeId][]Component
	insertIndex int
}

func (entities *Entities) registerComponent(typeId TypeId) {
	entities.components[typeId] = []Component{}
	entities.bitMasks[typeId] = 1 << len(entities.bitMasks)
}

func (entities *Entities) createEntity() *Entities {
	var index int
	found := false
	for i, mask := range entities.entityMap {
		if mask == 0 {
			index = i
			found = true
			break
		}
	}

	if index == 0 {
		for i, componentList := range entities.components {
			entities.components[i] = append(componentList, nil)
		}
		if !found {
			entities.entityMap = append(entities.entityMap, 0)
		}
		entities.insertIndex = len(entities.entityMap) - 1
	}

	return entities
}

func (entities *Entities) withComponent(component Component) *Entities {
	typeId := component.typeId()
	index := entities.insertIndex

	_, exists := entities.components[typeId]
	if !exists {
		panic(typeId + " is not registered")
	}

	entities.components[typeId][index] = component

	bitMask := entities.bitMasks[typeId]
	entities.entityMap[index] |= bitMask

	return entities
}

func (entities *Entities) getBitmask(typeId TypeId) int32 {
	return entities.bitMasks[typeId]
}
