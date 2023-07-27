package physics2d

type Vector struct {
	X float32
	Y float32
}

func (v *Vector) Add(o *Vector) *Vector {
    v.X += o.X
    v.Y += o.Y
    return v
}

func (v *Vector) Subtract(o *Vector) *Vector {
    v.X -= o.X
    v.Y -= o.Y
    return v
}

func (v *Vector) Scale(n float32) *Vector {
    v.X *= n
    v.Y *= n
    return v
}
