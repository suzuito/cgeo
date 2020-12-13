package cs2

import (
	"github.com/golang/geo/s1"
	"github.com/golang/geo/s2"
	"github.com/suzuito/cgeo/entity"
)

// NewCirclePolygon creates Polygon instance approximated to Circle
func (i *Impl) NewCirclePolygon(
	center *entity.LatLng,
	onChord *entity.LatLng,
	n int,
	maxRadius float64,
) *entity.Polygon {
	c := newS2PointFromLatLng(center)
	m := newS2PointFromLatLng(onChord)
	a := s2.ChordAngleBetweenPoints(c, m)
	a = minChordAngle(a, s1.ChordAngle(maxRadius))
	l := s2.RegularLoop(c, a.Angle(), n)
	return newPolygonFromS2Loop(l)
}

// NewCentroidFromPolygons ...
func (i *Impl) NewCentroidFromPolygons(
	polygons *[]entity.Polygon,
) *entity.LatLng {
	points := []s2.Point{}
	for _, polygon := range *polygons {
		for _, ll := range polygon.LatLngs {
			points = append(
				points,
				newS2PointFromLatLng(&ll),
			)
		}
	}
	l := s2.LoopFromPoints(points)
	return newLatLngFromS2Point(l.Centroid())
}

func newPolygonFromS2Loop(loop *s2.Loop) *entity.Polygon {
	ret := entity.Polygon{}
	for _, v := range loop.Vertices() {
		ll := s2.LatLngFromPoint(v)
		ret.LatLngs = append(ret.LatLngs, *newLatLngFromS2LatLng(&ll))
	}
	return &ret
}
