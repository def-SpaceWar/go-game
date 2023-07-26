package primitive

import "gogame/ecs"

func ForcesSystem(world *ecs.World) {
    for _, e := range world.Entities() {
        position := Position[e]
        velocity := Velocity[e]

        if position != nil && velocity != nil {
            position.Add(velocity.Clone().Scale(DeltaTime))
        }
    }
}
