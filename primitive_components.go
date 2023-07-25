package main

var Position = "position"

type PositionData struct {
	pos Vector
}

func (pos PositionData) Type() *string {
	return &Position
}

func (e Entity) Position(pos Vector) Entity {
    currentWorld.components = append(currentWorld.components, Component{e, PositionData{pos}})
    return e
}
