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
		if casted, ok := component.(T); ok {
			return &casted
		}
	}

	return nil
}

func GetComponentInChildren[T Component](e *Entity) *T {
	for _, child := range e.Children {
		for _, component := range child.Components {
			if casted, ok := component.(T); ok {
				return &casted
			}
		}
	}

	return nil
}
