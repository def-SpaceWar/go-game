package main

import "fmt"

var currentWorld World

func main() {
	currentWorld = World{
		entityCount: 0,
		components:  []Component{},
		systems:     []System{},
	}

	myEntity := currentWorld.GenerateEntity().
		Position(Vector{100, 100})

	component, err := myEntity.GetComponent(&Position)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(component)
	}

	for {
		currentWorld.RunSystems()
	}
}
