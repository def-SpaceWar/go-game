package main

type ComponentData interface {
    Type() string;
}

type Component struct {
	entity Entity
	data   ComponentData
}

func (c *Component) Entity() Entity {
    return c.entity
}
