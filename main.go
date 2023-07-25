package main

import "fmt"

var currentWorld World

func main() {
	currentWorld = World{
		entityCount: 0,
		components:  []Component{},
		systems:     []System{},
	}

	fmt.Println(currentWorld.GenerateEntity())
	fmt.Println(currentWorld.GenerateEntity())
	fmt.Println(currentWorld.GenerateEntity())

	for {
		currentWorld.RunSystems()
	}
}
