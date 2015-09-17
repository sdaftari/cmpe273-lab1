package shapeinterface

import "testing"

func TestPerimeterOfCircle(t *testing.T) {
	c := Circle{0,0,5}
	actualResult := TotalPerimeter(&c)
	if actualResult != 31.41592653589793 {
		t.Error("Actual Result - (%f), Expected Result - (%f)", actualResult, "31.41592653589793")
	}
}

func TestPerimeterOfRectangle(t *testing.T) {
	r := Rectangle{0,0,10,10}
	actualResult := TotalPerimeter(&r)
	if actualResult != 40 {
		t.Error("Actual Result - (%f), Expected Result - (%f)", actualResult, "40")
	}
}

func TestPerimeterOfCircleAndRectangle(t *testing.T) {
	c := Circle{0,0,5}
	r := Rectangle{0,0,10,10}
	actualResult := TotalPerimeter(&c, &r)
	if actualResult != 71.41592653589794 {
		t.Error("Actual Result - (%f), Expected Result - (%f)", actualResult, "71.41592653589793")
	}
}