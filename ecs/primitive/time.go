package primitive

import (
	"gogame/ecs"
	"time"
)

var DeltaTime float32

func CreateTimeSystem() ecs.System {
	var before time.Time = time.Now()
	return func(_ *ecs.World) {
		now := time.Now()
        DeltaTime = subtractTime(now, before)
		before = now
	}
}

func subtractTime(a, b time.Time) float32 {
    aTime := float64(a.UnixNano()) / 1_000_000_000
    bTime := float64(b.UnixNano()) / 1_000_000_000
    return float32(aTime - bTime)
}
