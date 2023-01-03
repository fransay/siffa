package types

// Polygon type polygon
type Polygon []Point2D

// polygon methods

// Area returns the area of the polygon
// algorithm: shoelace
// Optimise for resource and time complexities
func (p *Polygon) Area() (area float64) {
	// fetch Point2D points
	return area

}

// Centroid returns the centroid of the polygon
// algorithms: none yet
func (p *Polygon) Centroid() (cent float64) {
	return cent
}

// Perimeter returns total distance around the polygon
// traverse the total length of compositing strings
func (p *Polygon) Perimeter() (perim float64) {
	return perim
}

// ShortestLineSegment returns the shortest line string
// use ranking algorithm to sort out distances
func (p *Polygon) ShortestLineSegment() (perim float64) {
	return perim
}

// IsClosed returns a boolean value whether polygon is closed or not
func (p *Polygon) IsClosed() (isClosedStatus float64) {
	return isClosedStatus
}

// IsOpened returns a boolean value whether polygon is opened or not
func (p *Polygon) IsOpened() (isOpenedStatus float64) {
	return isOpenedStatus
}

// NumberOfLineSegments returns the number of compositing line segments
func (p *Polygon) NumberOfLineSegments() float64 {
	return 0.0
}

// NumberOfNodes number of nodes in a polygon
func (p *Polygon) NumberOfNodes() int {
	var numbReturn int = 1
	for index, _ := range *p {
		numbReturn += index
	}
	return numbReturn
}
