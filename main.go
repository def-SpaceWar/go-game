package main

import (
	"fmt"
	"gogame/ecs"
	"gogame/physics"
)

type position struct {
	physics.Vector
}
type velocity struct {
	physics.Vector
}

func main() {
	world := ecs.CreateWorld()
	entity := world.CreateEntity()
	entity.AddComponents(
		position{physics.Vector{5, 0}},
		velocity{physics.Vector{0, 2}},
	)

	fmt.Println(ecs.GetComponent[position](entity))
	fmt.Println(ecs.GetComponent[velocity](entity))
}
