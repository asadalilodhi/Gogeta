package iteration

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)


func TestRepeat(t *testing.T) {
	Convey("Given a character and a repeat amount", t, func() {
		result := repeat("a", 6)
		Convey("The character should repeat the given number of times", func() {
			So(result, ShouldEqual, "aaaaaa")
		})
	})
	Convey("Given a character and a repeat amount", t, func() {
		result := repeat("b", 8)
		Convey("The character should repeat the given number of times", func() {
			So(result, ShouldEqual, "bbbbbbbb")
		})
	})
	Convey("Given a character and a repeat amount", t, func() {
		result := repeat("x", -6)
		Convey("The character should repeat the given number of times", func() {
			So(result, ShouldEqual, "")
		})
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a",5)
	}
}

func ExampleRepeat() {
	repeated := repeat("a",6)
	fmt.Println(repeated)
	// Output: aaaaaa
}