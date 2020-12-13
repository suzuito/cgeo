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
