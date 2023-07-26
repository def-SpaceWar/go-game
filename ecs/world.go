package ecs

type World struct {
	entityCount       Entity
	destroyedEntities []Entity
	Systems           []System
	onEntityCreate    func(Entity)
	onEntityDestroy   func(Entity)
}

func (w *World) CreateEntity() Entity {
    var result Entity

	if len(w.destroyedEntities) > 0 {
		result = w.destroyedEntities[0]
		w.destroyedEntities = w.destroyedEntities[1:len(w.destroyedEntities)]
	} else {
		result = w.entityCount
		w.entityCount++
	}

	w.onEntityCreate(result)
	return result
}

func (w *World) DestroyEntity(e Entity) {
    w.destroyedEntities = append(w.destroyedEntities, e)
    w.onEntityDestroy(e)
}

func indexOf[T comparable](element T, data []T) int {
   for k, v := range data {
       if element == v {
           return k
       }
   }
   return -1
}

func (w *World) Entities() []Entity {
    entities := []Entity{}
    for e := Entity(0); e < Entity(w.entityCount); e++ {
        if indexOf(e, w.destroyedEntities) == -1 {
            entities = append(entities, e)
        }
    }
    return entities
}

func (w *World) RunSystems() {
	for i := 0; i < len(w.Systems); i++ {
		w.Systems[i](w)
	}
}

func CreateWorld(systems []System, onEntityCreate func(Entity), onEntityDestroy func(Entity)) World {
	return World{
		entityCount:     0,
		Systems:         systems,
		onEntityCreate:  onEntityCreate,
		onEntityDestroy: onEntityDestroy,
	}
}
