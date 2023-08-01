package render2d

import (
	"gogame/fp"
	"gogame/physics2d"

	"github.com/veandco/go-sdl2/sdl"
)

type Polygon struct {
	Points []*physics2d.Vector
	Color  *sdl.Color
	zIndex *int32
}

type PolygonParams struct {
	Points []*physics2d.Vector
	Color  fp.Maybe[sdl.Color]
	ZIndex fp.Maybe[int32]
}

func NewPolygon(params PolygonParams) Polygon {
	var points []*physics2d.Vector
	if params.Points == nil {
		points = []*physics2d.Vector{}
	} else {
		points = params.Points
	}

	var color *sdl.Color
	if fp.IsNone(params.Color) {
		color = RGB(255, 0, 255)
	} else {
		color = params.Color
	}

	var zIndex int32
	if fp.IsNone(params.ZIndex) {
		zIndex = 0
	} else {
		zIndex = *params.ZIndex
	}

	return Polygon{
		Points: points,
		Color:  color,
		zIndex: &zIndex,
	}
}

func (r Polygon) ZIndex() int32 {
	return *r.zIndex
}

func (r Polygon) SetZIndex(index int32) {
	*r.zIndex = index
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
			newVec.Rotate(*transform.Rotation)
			newVec.Add(transform.Position)
		}
		vertices = append(vertices, sdl.Vertex{
			Position: sdl.FPoint(newVec),
			Color:    *p.Color,
		})
	}
	return vertices
}
