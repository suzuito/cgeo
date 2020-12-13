package cs2

import (
	"github.com/golang/geo/s2"
	"github.com/suzuito/cgeo/entity"
)

func newS2PointFromLatLng(ll *entity.LatLng) s2.Point {
	return s2.PointFromLatLng(
		s2.LatLngFromDegrees(
			ll.Lat,
			ll.Lng,
		),
	)
}

func newLatLngFromS2Point(p s2.Point) *entity.LatLng {
	s2ll := s2.LatLngFromPoint(p)
	return &entity.LatLng{
		Lat: s2ll.Lat.Degrees(),
		Lng: s2ll.Lng.Degrees(),
	}
}
