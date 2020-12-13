package cs2

import "github.com/golang/geo/s1"

func minChordAngle(a, b s1.ChordAngle) s1.ChordAngle {
	if a < b {
		return a
	}
	return b
}
