package main

import (
	"fmt"
)

// Shape interface defining a method to calculate area
type Shape interface {
	GetArea() float64
}

// Circle struct implementing the Shape interface
type Circle struct {
	radius float64
}

func (c *Circle) GetArea() float64 {
	return 3.14 * c.radius * c.radius
}

// Rectangle struct implementing the Shape interface
type Rectangle struct {
	width, height float64
}

func (r *Rectangle) GetArea() float64 {
	return r.width * r.height
}

// ShapeFactory defining a factory method
type ShapeFactory struct{}

// CreateShape is the factory method to create shapes
func (sf *ShapeFactory) CreateShape(shapeType string, dimensions ...float64) Shape {
	switch shapeType {
	case "circle":
		return &Circle{radius: dimensions[0]}
	case "rectangle":
		return &Rectangle{width: dimensions[0], height: dimensions[1]}
	default:
		return nil
	}
}

func main() {
	factory := &ShapeFactory{}
	// Creating a circle
	circle := factory.CreateShape("circle", 5)
	fmt.Printf("Circle Area: %.2f\n", circle.GetArea())

	// Creating a rectangle
	rectangle := factory.CreateShape("rectangle", 4, 6)
	fmt.Printf("Rectangle Area: %.2f\n", rectangle.GetArea())
}
