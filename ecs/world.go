package ecs

type World struct {
	entities  Entities
	resources Resources
}

func (world *World) addResource(resource Resource) {
	world.resources.add(resource)
}

func (world *World) getResource(typeId string) Resource {
	return world.resources.get(typeId)
}

func (world *World) removeResource(typeId string) {
	world.resources.remove(typeId)
}

/// RESOURCES

type Resource interface {
	typeId() string
}

type Resources struct {
	data map[string]Resource
}

func (resources *Resources) add(resource Resource) {
	resources.data[resource.typeId()] = resource
}

func (resources *Resources) get(typeId string) Resource {
	return resources.data[typeId]
}

func (resources *Resources) remove(typeId string) {
	delete(resources.data, typeId)
}
