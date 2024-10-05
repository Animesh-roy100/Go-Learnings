package main

import "fmt"

// Defining a struct 'Engine' that will be embedded into another struct.
type Engine struct {
	HorsePower int
}

func (e Engine) Start() {
	fmt.Println("Engine started with horsepower:", e.HorsePower)
}

// Car struct that embeds the Engine struct (composition)
type Car struct {
	Engine    // Embedded struct (Engine)
	BrandName string
}

func main() {
	// Creating an instance of Car and setting fields
	myCar := Car{
		Engine:    Engine{HorsePower: 300},
		BrandName: "Tesla",
	}

	// Accessing the Engine's fields and methods directly
	fmt.Println("Car brand:", myCar.BrandName)
	myCar.Start() // Reusing Engine's Start method
}
