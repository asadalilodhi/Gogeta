package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)	

func TestPerimeter(t *testing.T) {
	Convey("Given a rectangle with height and width", t, func() {
		rectangle := Rectangle{10.0, 10.0}
		result := perimeter(rectangle)

		Convey("It should return the correct perimeter", func() {
			So(result, ShouldEqual, 40.0)
		})
	})

	
}

func TestArea(t *testing.T) {
	
	Convey("Given a shape", t, func() {
		
		areaTests := []struct {
			shape Shape
			want float64
		}{
			{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
    		{shape: Circle{Radius: 10}, want: 314.1592653589793},
        	{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
		}

		for _, tt := range areaTests {
			tt := tt
			Convey(fmt.Sprintf("should return the area of %T", tt.shape), func() {
				So(tt.shape.area(), ShouldAlmostEqual, tt.want)
			})
		}

	})

}