package ecs

type (
	queryIndexes    []int
	queryComponents [][]Component
)

type Query struct {
	queryMap int32
	entities *Entities
	typeIds  []TypeId
}

func (query *Query) new(entities *Entities) Query {
	return Query{
		entities: entities,
		queryMap: 0,
	}
}

func (query *Query) withComponent(typeId TypeId) *Query {
	bitMask := query.entities.getBitmask(typeId)
	query.queryMap |= bitMask
	query.typeIds = append(query.typeIds, typeId)
	return query
}

func (query *Query) run() (queryIndexes, queryComponents) {
	var filteredEntities []int
	for index, entityMap := range query.entities.entityMap {
		if entityMap&query.queryMap == query.queryMap {
			filteredEntities = append(filteredEntities, index)
		}
	}

	var result queryComponents

	for _, type_id := range query.typeIds {
		entityComponents := query.entities.components[type_id]
		var componentsToKeep []Component
		for _, filteredIndex := range filteredEntities {
			componentsToKeep = append(componentsToKeep, entityComponents[filteredIndex])
		}
		result = append(result, componentsToKeep)
	}

	return filteredEntities, result
}
