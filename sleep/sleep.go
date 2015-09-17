package main

import("fmt";"time")

func SleepFunc(interval int) {
    <-time.After(time.Second * time.Duration(interval))
}

func main(){
	for k:=0;k<10;k++{
		fmt.Println(k)
		SleepFunc(2)
	}
}

