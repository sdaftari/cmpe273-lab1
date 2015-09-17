package main

import "fmt"
import shape "Lab1/interface/InterfaceLibrary"

func main() {
	var x,y,r,x1,y1,x2,y2,t float64
	fmt.Println("Enter x co-ordinate of center of a circle : ")
	fmt.Scanf("%f", &x)
	fmt.Scanln(&t)
	fmt.Println("Enter y co-ordinate of center of a circle : ")
	fmt.Scanf("%f", &y)
	fmt.Scanln(&t)
	fmt.Println("Enter radius of a circle : ")
	fmt.Scanf("%f", &r)
	fmt.Scanln(&t)
	c := shape.Circle{x,y,r}

	fmt.Println("Enter first x co-ordinate of rectangle : ")
	fmt.Scanf("%f", &x1)
	fmt.Scanln(&t)
	fmt.Println("Enter first y co-ordinate of rectangle : ")
	fmt.Scanf("%f", &y1)
	fmt.Scanln(&t)
	fmt.Println("Enter second x co-ordinate of rectangle : ")
	fmt.Scanf("%f", &x2)
	fmt.Scanln(&t)
	fmt.Println("Enter second y co-ordinate of rectangle : ")
	fmt.Scanf("%f", &y2)
	fmt.Scanln(&t)
	rect := shape.Rectangle{x1,y1,x2,y2}

	fmt.Println("Total area is: ", shape.TotalArea(&c, &rect))
	fmt.Println("Total perimeter of circle is : ", shape.TotalPerimeter(&c))
	fmt.Println("Total perimeter of rectangle is : ", shape.TotalPerimeter(&rect))
	fmt.Println("Total perimeter is : ", shape.TotalPerimeter(&c, &rect))
}