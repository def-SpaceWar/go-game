package ecs

type World struct {
	entityCount EntityID
	Entities    []Entity
	systems     []System
}

func CreateWorld() World {
	return World{
		entityCount: 0,
		Entities:    []Entity{},
		systems:     []System{},
	}
}

func (world *World) CreateEntity() *Entity {
    entity := Entity{
        Id: world.entityCount,
        Parent: nil,
        Components: []Component{},
    }

    world.entityCount++
    return &entity
}

func (world *World) CreateChildEntity(parent *Entity) *Entity {
    entity := Entity{
        Id: world.entityCount,
        Parent: parent,
        Components: []Component{},
    }

    world.entityCount++
    return &entity
}
