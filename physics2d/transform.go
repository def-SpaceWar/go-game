package physics2d

import "gogame/fp"

type Transform struct {
	Position *Vector
	Rotation *float32
	Scale    *Vector
	//Transform *Matrix
}

func NewTransform(position *Vector, rotation fp.Maybe[float32], scale fp.Maybe[Vector]) Transform {
	if fp.IsNone(scale) {
		if fp.IsNone(rotation) {
			rotation := float32(0)
			return Transform{
				Position: position,
				Rotation: &rotation,
				Scale:    Vec(1, 1),
			}
		}

		return Transform{
			Position: position,
			Rotation: rotation,
			Scale:    Vec(1, 1),
		}
	}

	return Transform{
		Position: position,
		Rotation: rotation,
		Scale:    scale,
	}
}
