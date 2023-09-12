package functs

import (
	"math"
	"shapelib/types"
)

// EDistance returns Euclidean distance between two points/coordinates
func EDistance(pointOne, pointTwo types.Point2D) (dist float64) {
	deltaX := pointTwo.X - pointOne.X
	deltaY := pointOne.Y - pointOne.Y
	dist = math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
	return dist
}

// MinkowskiDistance returns minkowski distance of points in an N dimensional space,
// basically an approx of Euclidean and manhattan distance.
func MinkowskiDistance(PointA, PointB types.Point2D, p float64) (distMinkowski float64) {
	var deltaA = math.Pow(PointA.X-PointA.Y, p)
	var deltaB = math.Pow(PointB.X-PointB.Y, p)
	distMinkowski = math.Pow(deltaA+deltaB, 1/p)
	return distMinkowski
}
