package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHello(t *testing.T) {
	Convey("Given a name and no language", t, func() {
		result := hello("Chris", "")
		Convey("It should return 'Hello, Chris'", func() {
			So(result, ShouldEqual, "Hello, Chris")
		})
	})

	Convey("Given an empty name", t, func() {
		result := hello("", "")
		Convey("It should default to 'Hello, World'", func() {
			So(result, ShouldEqual, "Hello, World")
		})
	})
}
