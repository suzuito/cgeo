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
	) *entity.Cell
}
