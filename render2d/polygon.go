package render2d

import (
	"gogame/physics2d"

	"github.com/veandco/go-sdl2/sdl"
)

type Polygon struct {
	Points []*physics2d.Vector
	Color  *sdl.Color
	zIndex int32
}

func (r Polygon) ZIndex() int32 {
	return r.zIndex
}

func Rect(x, y, w, h float32) []*physics2d.Vector {
	return []*physics2d.Vector{
		physics2d.Vec(x+(w/2), y+(h/2)),
		physics2d.Vec(x-(w/2), y+(h/2)),
		physics2d.Vec(x-(w/2), y-(h/2)),
		physics2d.Vec(x+(w/2), y-(h/2)),
		physics2d.Vec(x-(w/2), y-(h/2)),
		physics2d.Vec(x+(w/2), y+(h/2)),
	}
}

func (p *Polygon) toVertices(transforms []*physics2d.Transform) []sdl.Vertex {
	vertices := []sdl.Vertex{}
	for _, vec := range p.Points {
		newVec := vec.Clone()
		for _, transform := range transforms {
			newVec.ScaleVec(transform.Scale)
			newVec.Add(transform.Position)
			//newVec.Rotate(transform.Rotation)
		}
		vertices = append(vertices, sdl.Vertex{
			Position: sdl.FPoint(newVec),
			Color:    *p.Color,
		})
	}
	return vertices
}
