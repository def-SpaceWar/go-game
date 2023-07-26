package physics

type Vector struct {
	X float32
	Y float32
}

func Vec(x, y float32) *Vector {
	return &Vector{X: x, Y: y}
}

func (v Vector) Clone() *Vector {
    n := v
    return &n
}

func (v *Vector) Add(o *Vector) *Vector {
    v.X += o.X
    v.Y += o.Y
    return v
}

func (v *Vector) Scale(scalar float32) *Vector {
    v.X *= scalar
    v.Y *= scalar
    return v
}
