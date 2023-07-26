package primitive

import (
	"gogame/ecs"
	"gogame/physics"
)

var Position = ecs.Component[physics.Vector]{}
var Velocity = ecs.Component[physics.Vector]{}
var Rotation = ecs.Component[float32]{}
