package cs2

import (
	"fmt"
	"sort"
	"strings"

	"github.com/golang/geo/s2"
	"github.com/suzuito/cgeo/entity"
	"golang.org/x/xerrors"
)

func (i *Impl) NewCellFromLatLng(
	position *entity.LatLng,
	level int,
) *entity.Cell {
	cell := newCellAtLevel(position, level)
	loop := s2.LoopFromCell(*cell)
	ret := entity.Cell{
		ID:      entity.CellID(cell.ID().ToToken()),
		Polygon: *newPolygonFromS2Loop(loop),
	}
	return &ret
}

func (i *Impl) NewRangeCellIDs(
	southWest entity.LatLng,
	northEast entity.LatLng,
) (entity.CellID, entity.CellID, error) {
	rb := s2.NewRectBounder()
	rb.AddPoint(newS2PointFromLatLng(&southWest))
	rb.AddPoint(newS2PointFromLatLng(&northEast))
	rc := s2.NewRegionCoverer()
	rc.MaxCells = 5
	rc.MinLevel = 2
	rc.MaxLevel = 20
	rc.LevelMod = 1
	ucell := rc.Covering(rb.RectBound())
	cellIDs := []string{}
	for _, cellID := range ucell {
		cellIDs = append(cellIDs, cellID.ToToken())
	}
	if len(cellIDs) <= 1 {
		return entity.CellID(""), entity.CellID(""), xerrors.Errorf("Cannot create range '%s'", strings.Join(cellIDs, ","))
	}
	sort.Strings(cellIDs)
	fmt.Println(cellIDs)
	return entity.CellID(cellIDs[0]), entity.CellID(cellIDs[len(cellIDs)-1]), nil
}
