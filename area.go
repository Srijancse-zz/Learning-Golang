//program to print sum of all numbers below 50 completely divisible by 3 or 5

package main

import(
        "fmt"
	"os"
	"strconv"
)

type Rectangle struct {
        length float64
	width float64
}
type circle struct {
        radius float64
}
func (r Rectangle) Area() float64{
        return r.length * r.width
}
func (r Rectangle) Perimeter() float64{
        return 2 * (r.length + r.width)
}
func (c circle) Area() float64{
	return 3.14 * r.radius * r.radius
}
func (c circle) Perimeter() float64{
        return 2 * 3.14 * r.radius
}

func main() {
 	shape := os.Args[1]
	if shape == "circle" {
		r := os.Args[2]
		radius, _ := strconv.Atoi(r)
		circle := Circle{float64(radius)}
		area := circle.Area()
		perimeter := circle.Perimeter()
		fmt.Println("Area:", area)
		fmt.Println("Perimeter:", perimeter)

       }else {
		w := os.Args[2]
		h := os.Args[2]
		width, _ := strcov.Atoi(w)
		height, _ := strcov.Atoi(h)
		rectangle := Rectangle{float64(width,height)}
		area := Rectangle.Area()
		perimeter := Rectangle.Perimeter()
		fmt.Println("Area:", area)
		fmt.Println("Perimeter:", perimeter)
	}
}
