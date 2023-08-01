package ecs

var CurrentWorld *World
var Running = true

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

func (world *World) CreateEntity(components ...Component) *Entity {
	entity := Entity{
		Id:         world.entityCount,
		Parent:     nil,
		Components: components,
	}

	world.Entities = append(world.Entities, entity)
	world.entityCount++
	return &entity
}

func (world *World) CreateChildEntity(parent *Entity, components ...Component) *Entity {
	entity := world.CreateEntity(components...)
	entity.Parent = parent
	return entity
}

func (world *World) AddSystems(systems ...System) {
	var newSystem System = func(world *World) {
		for _, system := range systems {
			system(world)
		}
	}

	world.systems = append(world.systems, newSystem)
}

func (world *World) RunSystems() {
	for _, system := range world.systems {
		system(world)
		//go system(world)
	}
}

type pair[A, B any] struct {
	First  A
	Second B
}

func FindComponents[T Component](w *World) chan pair[*Entity, *T] {
	c := make(chan pair[*Entity, *T])
	go func() {
		for _, entity := range w.Entities {
			for _, component := range entity.Components {
				casted, ok := component.(T)
				if ok {
					c <- pair[*Entity, *T]{&entity, &casted}
				}
			}
		}
		close(c)
	}()
	return c
}

func FindComponentsSlice[T Component](w *World) []pair[*Entity, *T] {
	components := []pair[*Entity, *T]{}
	for _, entity := range w.Entities {
		for _, component := range entity.Components {
			casted, ok := component.(T)
			if ok {
				components = append(components, pair[*Entity, *T]{&entity, &casted})
			}
		}
	}
	return components
}
