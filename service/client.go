package service

import (
	geojson "github.com/paulmach/go.geojson"
	"github.com/suzuito/cgeo/entity"
)

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
	NewPolygonsFromGeojsonPolygon(
		gpolygon *geojson.Geometry,
	) (*[]entity.Polygon, error)
}
