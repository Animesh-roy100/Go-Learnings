package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Measurable interface {
	Perimeter() float64
}

// embeds two other interfaces
// interface composition
type Geometry interface {
	Shape
	Measurable
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// func calculateArea(s Shape) float64 {
// 	return s.Area()
// }

func describeShape(g Geometry) {
	fmt.Println("Area: ", g.Area())
	fmt.Println("Perimeter: ", g.Perimeter())
}

func describeValue(t interface{}) {
	fmt.Printf("Type: %T Value: %v\n", t, t)
}

func main() {
	rect := Rectangle{width: 3, height: 4}
	// circle := Circle{radius: 5}

	// fmt.Println("Rectangle Area: ", calculateArea(rect))
	// fmt.Println("Circle Area: ", calculateArea(circle))

	describeShape(rect)

	// interface type -----------------------------------
	mysteryBox := interface{}("10")
	describeValue(mysteryBox)

	retrievedInt, ok := mysteryBox.(int)
	if ok {
		fmt.Println("Retrieved int:", retrievedInt)
	} else {
		fmt.Println("mysteryBox is not an int")
	}

	// Errors and Interfaces --------------------------------
	result, err := performCalculation(10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

type CalculationError struct {
	msg string
}

func (ce CalculationError) Error() string {
	return ce.msg
}

func performCalculation(val float64) (float64, error) {
	if val < 0 {
		return 0, CalculationError{
			msg: "Invalid input",
		}
	}

	return math.Sqrt(val), nil
}
