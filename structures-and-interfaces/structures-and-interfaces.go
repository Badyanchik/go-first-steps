package main

import ("fmt"; "math")

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	area() float64
	perimetr() float64
}

type MultiShape struct {
	shapes []Shape
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, shape := range shapes {
		area += shape.area()
	}
	return area
}

func totalPerimetr(rectangles ...Shape) float64 {
	var perimetr float64
	for _, shape := range rectangles {
		perimetr += shape.perimetr()
	}
	return perimetr
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, shape := range m.shapes {
		area += shape.area()
	}
	return area
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimetr() float64 {
	return math.Pi * (c.r * 2)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimetr() float64 {
	return r.y1 + r.x1 + r.x2 + r.y2;
}

type Person struct {
	name string
}

type Android struct {
	Person
	model string
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is ", p.name)
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0 ,10, 10}
	r1 := Rectangle{10, 20 ,30, 40}
	p := Person{"Bohdan"}
	a := Android{
		Person: Person{"Samsung"},
		model:  "Galaxy S10",
	}
	fmt.Println(totalArea(&c, &r))
	fmt.Println(totalPerimetr(&c, &r1))
	p.talk()
	a.talk()
	a.Person.talk()
}