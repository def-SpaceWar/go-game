package main

import (
	"fmt"
	"gogame/ecs"
)

type position struct {
    X float32
    Y float32
}

type velocity struct {
    X float32
    Y float32
}

func main() {
    entity := ecs.Entity{
        Components: []ecs.Component{
            velocity{},
        },
    }

    pos := *ecs.GetComponent[position](entity)
    fmt.Println(pos) // INTENTIONAL ERROR
}
