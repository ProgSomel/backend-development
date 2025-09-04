package main

import "fmt"

type Rectangle struct {
	length float64
	width float64
}

func(r Rectangle) Area() float64 {
	return r.length * r.width
}

func main(){
	rect := Rectangle{length: 10, width: 0}
	area := rect.Area()
	fmt.Println("Area of Rectangle with width 9 and length 10 is: ", area)
}