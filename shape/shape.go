package main
import ("fmt";"math")
type shape interface{
	Area() float64
	Perimeter() float64
}

type circle struct{
	radius float64
}

type rectangle struct{
	x1, y1, x2, y2 float64
}

func (r rectangle) Area() float64{
	length := Distance(r.x1, r.y1, r.x1, r.y2)
	width := Distance(r.x1, r.y1, r.x2, r.y1)
	return length * width
}

func (r rectangle) Perimeter() float64{
	length := Distance(r.x1, r.y1, r.x1, r.y2)
	width := Distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (length + width)
}

func Distance(x1, y1 ,x2 , y2 float64) float64{
	return math.Sqrt(math.Pow(y2-y1, 2) + math.Pow(x2-x1, 2))
}

func (c circle) Area() float64{
	return math.Pi * c.radius * c.radius
}

func (c circle) Perimeter() float64{
	return 2 * math.Pi * c.radius
}

func main(){
	r := rectangle{0,0,10,10}
	fmt.Println(r.Area())
        fmt.Println(r.Perimeter())
	c := circle{3}
	fmt.Println(c.Area())
        fmt.Println(c.Perimeter())
}