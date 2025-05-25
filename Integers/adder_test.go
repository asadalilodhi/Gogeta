package integers

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdder(t *testing.T) {
	Convey("Given two integers", t, func() {
		sum_1 := add(2, 2)
		Convey("The result should be their sum", func() {
			So(sum_1, ShouldEqual, 4)
		})
	})
	Convey("Given two integers", t, func() {
		sum_2 := add(5, -6)
		Convey("The result should be their sum", func() {
			So(sum_2, ShouldEqual, -1)
		})
	})
}

func Exampleadd() {
	sum := add(1, 5)
	fmt.Println(sum)
	// Output: 6
}