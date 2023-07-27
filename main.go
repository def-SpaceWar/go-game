package main

import (
	"fmt"
	"gogame/ecs"
	"gogame/physics2d"
)

type position struct {
	physics2d.Vector
}
type velocity struct {
	physics2d.Vector
}

func main() {
	world := ecs.CreateWorld()
	entity := world.CreateEntity()
	entity.AddComponents(
		position{physics2d.Vector{5, 0}},
		velocity{physics2d.Vector{0, 2}},
	)

	fmt.Println(ecs.GetComponent[position](entity))
	fmt.Println(ecs.GetComponent[velocity](entity))
}
