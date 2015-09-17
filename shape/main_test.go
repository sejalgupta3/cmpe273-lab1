package main
import "testing"

type testShape struct{
	object shape
	area, perimeter float64
}

var testShapeArray = []testShape{
	{circle{2}, 12.566370614359172, 12.566370614359172},
	{circle{1}, 3.141592653589793, 6.283185307179586},
	{rectangle{0, 0, 10, 10}, 100, 40},
	{rectangle{0, 0, 1, 2}, 2, 6},
}

func TestArea(t *testing.T){
	for _, test:= range testShapeArray{
		i := test.object.Area()
		if i != test.area{
			t.Error(
				"For", test.object,
				"expected", test.area,
				"got", i,
				)
		}
	}
}

func TestPerimeter(t *testing.T){
	for _, test:= range testShapeArray{
		i := test.object.Perimeter()
		if i != test.perimeter{
			t.Error(
				"For", test.object,
				"expected", test.perimeter,
				"got", i,
				)
		}
	}
}

