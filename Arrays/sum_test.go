package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSum(t *testing.T) {
	Convey("Given an array of integers", t, func() {
		numbers := []int{1, 2, 3, 4, 5}
		result := sum(numbers)
		Convey("The sum of the array of integers should be ", func() {
			So(result, ShouldEqual, 15)
		})
	})

	Convey("Given an array of integers", t, func() {
		numbers := []int{1, 2, 3}
		result := sum(numbers)
		Convey("The sum of the array of integers should be ", func() {
			So(result, ShouldEqual, 6)
		})
	})

	Convey("Given two array of integers", t, func() {
		array1 := []int{1, 2}
		array2 := []int{0, 9}
		result := sumAll(array1, array2)
		Convey("The sum of the two arrays of integers should be ", func() {
			So(result, ShouldResemble, []int{3, 9})
		})
	})

	Convey("Given two array of integers", t, func() {
		array1 := []int{1, 2}
		array2 := []int{0, 9}
		result := sumAlltails(array1, array2)
		Convey("The sum of the tails of two arrays of integers should be ", func() {
			So(result, ShouldResemble, []int{2, 9})
		})
	})

	Convey("Given two array of integers", t, func() {
		array1 := []int{}
		array2 := []int{0, 4, 5}
		result := sumAlltails(array1, array2)
		Convey("The sum of the tails of two arrays of integers should be ", func() {
			So(result, ShouldResemble, []int{0, 9})
		})
	})
}