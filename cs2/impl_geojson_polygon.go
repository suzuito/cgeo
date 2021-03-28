package cs2

import (
	geojson "github.com/paulmach/go.geojson"
	"github.com/suzuito/cgeo/entity"
	"golang.org/x/xerrors"
)

// NewPolygonsFromGeojsonPolygon ...
func (i *Impl) NewPolygonsFromGeojsonPolygon(
	gpolygon *geojson.Geometry,
) (*[]*entity.Polygon, error) {
	return newPolygonsFromGeoJSONPolygon(gpolygon)
}

func newPolygonsFromGeoJSONPolygon(gpolygon *geojson.Geometry) (*[]*entity.Polygon, error) {
	ret := []*entity.Polygon{}
	if gpolygon.Type != geojson.GeometryPolygon {
		return nil, xerrors.Errorf("Geometry type is not Polygon '%s'", gpolygon.Type)
	}
	for _, polygon := range gpolygon.Polygon {
		for _, point := range polygon {
			p := entity.Polygon{
				LatLngs: []entity.LatLng{},
			}
			p.LatLngs = append(p.LatLngs, entity.LatLng{
				Lat: point[1],
				Lng: point[0],
			})
			ret = append(ret, &p)
		}
	}
	return &ret, nil
}
