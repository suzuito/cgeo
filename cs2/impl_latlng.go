package cs2

import (
	"github.com/golang/geo/s2"
	"github.com/suzuito/cgeo/entity"
)

func newLatLngFromS2LatLng(ll *s2.LatLng) *entity.LatLng {
	return &entity.LatLng{
		Lat: ll.Lat.Degrees(),
		Lng: ll.Lng.Degrees(),
	}
}
