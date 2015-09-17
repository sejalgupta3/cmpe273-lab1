package main
import "testing"

type testVar struct{
	input, result int
}

var tests = []testVar{
	{0, 0},
	{4, 3},
	{3, 2},
	{2, 1},
	{6, 8},
	{1, 1},
	{9, 34},
}

func TestFibSeries(t *testing.T){
	for _, test:= range tests{
		i := FibSeries(test.input)
		if i != test.result{
			t.Error(
				"For", test.input,
				"expected", test.result,
				"got", i,
				)
		}else{
			t.Log("success",test)
		}
	}
}
