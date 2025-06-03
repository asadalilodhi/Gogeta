package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSearch(t *testing.T) {
	Convey("Search function", t, func() {
		dictionary := Dictionary{"test": "this is just a test"}
		
		Convey("should return the correct value for a valid key", func() {
			result, err := dictionary.search("test")
			want := "this is just a test"
			So(err, ShouldBeNil)
			So(result, ShouldEqual, want)
		})

		Convey("should return 'word not in dictionary' for a non-existent key", func() {
			_, err := dictionary.search("unkown")
			want := "could not find the word you were looking for"
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, want)
		})
	
	})

}