package main

import "math"

func main() {

}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A := nums1
	B := nums2
	var Aleft, Aright float64
	var Bleft, Bright float64
	total := len(A) + len(B)
	half := (total + 1) / 2
	if len(A) > len(B) {
		A, B = B, A
	}
	left := 0
	right := len(A) - 1
	for {
		midA := (left + right) >> 1
		midB := half - midA - 2
		if midA >= 0 {
			Aleft = float64(A[midA])
		} else {
			Aleft = math.Inf(-1)
		}
		if (midA + 1) < len(A) {
			Aright = float64(A[midA+1])
		} else {
			Aright = math.Inf(1)
		}

		if midB >= 0 {
			Bleft = float64(B[midB])
		} else {
			Bleft = math.Inf(-1)
		}

		if (midB + 1) < len(B) {
			Bright = float64(B[midB+1])
		} else {
			Bright = math.Inf(1)
		}
		if Aleft <= Bright && Bleft <= Aright {
			// odd
			if total%2 == 1 {
				return max(Aleft, Bleft)
			}
			// even
			return (max(Aleft, Bleft) + min(Aright, Bright)) / 2
		} else if Aleft > Bright {
			right = midA - 1
		} else {
			left = midA + 1
		}
	}
}
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b

}
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
