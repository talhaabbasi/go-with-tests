package structs

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {

		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func BenchmarkPerimeter(b *testing.B) {
	rectangle := Rectangle{10.0, 10.0}
	for i := 0; i < b.N; i++ {
		Perimeter(rectangle)
	}
}

func ExamplePerimeter() {
	rectangle := Rectangle{10.0, 10.0}
	repeats := Perimeter(rectangle)
	fmt.Println(repeats)
	// Output: 40
}

func BenchmarkRectangleArea(b *testing.B) {
	rectangle := Rectangle{10.0, 10.0}
	for i := 0; i < b.N; i++ {
		rectangle.Area()
	}
}

func ExampleRectangleArea() {
	rectangle := Rectangle{10.0, 10.0}
	repeats := rectangle.Area()
	fmt.Println(repeats)
	// Output: 100
}

func BenchmarkCircleArea(b *testing.B) {
	circle := Circle{10.0}
	for i := 0; i < b.N; i++ {
		circle.Area()
	}
}

func ExampleCircleArea() {
	circle := Circle{10.0}
	repeats := circle.Area()
	fmt.Println(repeats)
	// Output: 314.1592653589793
}

func BenchmarkTriangleArea(b *testing.B) {
	triangle := Triangle{10.0, 10.0}
	for i := 0; i < b.N; i++ {
		triangle.Area()
	}
}

func ExampleTriangleArea() {
	triangle := Triangle{10.0, 10.0}
	repeats := triangle.Area()
	fmt.Println(repeats)
	// Output: 50
}
