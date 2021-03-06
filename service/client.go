package service

import "github.com/suzuito/cgeo/entity"

// Client ...
type Client interface {
	NewCirclePolygon(
		center *entity.LatLng,
		onChord *entity.LatLng,
		n int,
		maxRadius float64,
	) *entity.Polygon
	NewCentroidFromPolygons(
		polygons *[]entity.Polygon,
	) *entity.LatLng
	NewCellFromLatLng(
		position *entity.LatLng,
		level int,
	) *entity.Cell
	NewRangeCellIDs(
		southWest entity.LatLng,
		northEast entity.LatLng,
	) (entity.CellID, entity.CellID, error)
}
