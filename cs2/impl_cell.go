package cs2

import (
	"github.com/golang/geo/s2"
	"github.com/suzuito/cgeo/entity"
)

func (i *Impl) NewCellFromLatLng(
	position *entity.LatLng,
) *entity.Cell {
	point := newS2PointFromLatLng(position)
	cell := s2.CellFromPoint(point)
	loop := s2.LoopFromCell(cell)
	ret := entity.Cell{
		ID:      entity.CellID(cell.ID().ToToken()),
		Polygon: *newPolygonFromS2Loop(loop),
	}
	return &ret
}
