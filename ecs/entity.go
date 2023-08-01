package ecs

type EntityID uint32
type Entity struct {
	Id         EntityID
	Parent     *Entity
	Children   []*Entity
	Components []Component
}

func (e *Entity) AddComponents(c ...Component) {
	e.Components = append(e.Components, c...)
}

func (e *Entity) AddChildren(c ...*Entity) {
	e.Children = append(e.Children, c...)
}

func GetComponent[T Component](e *Entity) *T {
	for _, component := range e.Components {
		casted, ok := component.(T)
		if ok {
			return &casted
		}
	}
	return nil
}

func GetComponents[T Component](e *Entity) chan *T {
	c := make(chan *T)
	go func() {
		for _, component := range e.Components {
			casted, ok := component.(T)
			if ok {
				c <- &casted
			}
		}
		close(c)
	}()
	return c
}

func GetComponentsSlice[T Component](e *Entity) []*T {
	components := []*T{}
	for _, component := range e.Components {
		casted, ok := component.(T)
		if ok {
			components = append(components, &casted)
		}
	}
	return components
}

func GetComponentInChildren[T Component](e *Entity) *T {
	for _, child := range e.Children {
        component := GetComponent[T](child)
        if component != nil {
            return component
        }
        component = GetComponentInChildren[T](child)
        if component != nil {
            return component
        }
	}
    return nil
}
