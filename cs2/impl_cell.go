package cs2

import (
	"sort"

	"github.com/golang/geo/s2"
	"github.com/suzuito/cgeo/entity"
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
) ([]entity.CellID, error) {
	rb := s2.NewRectBounder()
	rb.AddPoint(newS2PointFromLatLng(&southWest))
	rb.AddPoint(newS2PointFromLatLng(&northEast))
	/**
	 * 1
	 */
	// rc := s2.NewRegionCoverer()
	// rc.MaxCells = 10
	// rc.MinLevel = 2
	// rc.MaxLevel = 50
	// rc.LevelMod = 1
	// // ucell := rc.Covering(rb.RectBound())
	// ucell := rc.InteriorCovering(rb.RectBound())
	// cellIDs := []string{}
	// for _, cellID := range ucell {
	// 	cellIDs = append(cellIDs, cellID.ToToken())
	// }
	// if len(cellIDs) <= 1 {
	// 	return nil, xerrors.Errorf("Cannot create range '%s'", strings.Join(cellIDs, ","))
	// }
	// sort.Strings(cellIDs)
	/**
	 * 2
	 */
	// cellIDs := []string{}
	// for _, cellID := range rb.RectBound().CellUnionBound() {
	// 	cellIDs = append(cellIDs, cellID.ToToken())
	// }
	// sort.Strings(cellIDs)
	/**
	 * 3
	 */
	cellIDs := []string{}
	rect := rb.RectBound()
	centroid := rect.Centroid()
	for _, cellID := range s2.SimpleRegionCovering(rect, centroid, 15) {
		cellIDs = append(cellIDs, cellID.ToToken())
	}
	sort.Strings(cellIDs)

	ret := []entity.CellID{}
	for _, cellID := range cellIDs {
		ret = append(ret, entity.CellID(cellID))
	}
	return ret, nil
}
