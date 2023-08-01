package physics2d

import "gogame/fp"

type Transform struct {
	Position *Vector
	Rotation *float32
	Scale    *Vector
	//Transform *Matrix
}

type TransformParams struct {
	Position *Vector
	Rotation fp.Maybe[float32]
	Scale    fp.Maybe[Vector]
}

func NewTransform(params TransformParams) Transform {
	if fp.IsNone(params.Scale) {
		if fp.IsNone(params.Rotation) {
			rotation := float32(0)
			return Transform{
				Position: params.Position,
				Rotation: &rotation,
				Scale:    Vec(1, 1),
			}
		}

		return Transform{
			Position: params.Position,
			Rotation: params.Rotation,
			Scale:    Vec(1, 1),
		}
	}

	return Transform{
		Position: params.Position,
		Rotation: params.Rotation,
		Scale:    params.Scale,
	}
}
