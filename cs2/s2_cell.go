package cs2

import (
	"github.com/golang/geo/s2"
	"github.com/suzuito/cgeo/entity"
)

func newCellAtLevel(
	position *entity.LatLng,
	lv int,
) *s2.Cell {
	cell := s2.CellFromLatLng(s2.LatLngFromDegrees(position.Lat, position.Lng))
	for {
		if cell.Level() == lv {
			break
		}
		parentCellID := cell.ID().Parent(cell.Level() - 1)
		cell = s2.CellFromCellID(parentCellID)
	}
	return &cell
}
