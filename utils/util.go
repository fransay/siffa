package utils

import "math"

// EqualSlice compare two slice of types float64
func EqualSlice(sliceOne, sliceTwo []float64) (equal bool) {
	if len(sliceOne) != len(sliceTwo) {
		equal = false
		return
	}

	for index, _ := range sliceOne {
		if sliceOne[index] != sliceTwo[index] {
			equal = false
			return
		}
	}
	equal = true
	return equal
}

// AbsDiff absolute difference between two quantities of type float64
func AbsDiff(quantOne, quantTwo float64) (abs float64) {
	abs = math.Abs(quantTwo - quantOne)
	return abs
}

// Rad2Deg convert radians to degree
func Rad2Deg(radians float64) (deg float64) {
	deg = radians * (180 / math.Pi)
	return deg
}

// Deg2Rad convert degrees to radians
func Deg2Rad(deg float64) (rad float64) {
	rad = deg * (math.Pi / 180)
	return rad
}
