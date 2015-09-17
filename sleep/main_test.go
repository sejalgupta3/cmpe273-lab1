package main

import "testing"
import "time"

type testSleep struct{
	time int
	waiteTime int
}

var testSleepArray = []testSleep{
	{1,1},
	{10,10},
	{7,7},
}

func TestSleepFunc(t *testing.T){
	for _, test:= range testSleepArray{
		tBefore:= time.Now()
		SleepFunc(test.time)
		tAfter:= time.Now()
		tDiff:=tAfter.Sub(tBefore)
		i := int(tDiff.Nanoseconds())
		i = i/1000000000
		if i != test.waiteTime{
			t.Error(
				"For", test.time,
				"expected", test.waiteTime,
				"got", i,
				)
		}
	}
}