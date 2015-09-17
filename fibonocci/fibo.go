package main
import "fmt"

func main(){
	var seriesRange int
	fmt.Println("Enter the number of terms:")
	fmt.Scanf("%d", &seriesRange)
	fmt.Println("The series is:")
	for i:=0;i<seriesRange;i++{
		fmt.Println(FibSeries(i))
	}
}

func FibSeries(limit int) int{
	if (limit < 2){
		return limit
	}else{
		return FibSeries(limit - 1) + FibSeries(limit - 2) 
	}	
}
	