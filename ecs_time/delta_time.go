package ecs_time

import (
	"gogame/ecs"
	"time"
	"github.com/orcaman/concurrent-map/v2"
)

func SubtractTimes(now, before int64) float32 {
	return float32(now-before) / 1_000_000_000
}

var DeltaTimes = cmap.New[float32]()

func CreateTimedSystem(name string, systems ...ecs.System) ecs.System {
	before := time.Now().UnixNano()
	return func(w *ecs.World) (ecs.WorldState, *ecs.World) {
		now := time.Now().UnixNano()
		for _, system := range systems {
			DeltaTimes.Set(name, SubtractTimes(now, before))
            worldState, nextWorld := system(w)
            if worldState != ecs.CONTINUE {
                return worldState, nextWorld
            }
		}
		before = now
        return ecs.CONTINUE, nil
	}
}
