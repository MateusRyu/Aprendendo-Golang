package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

type figura interface {
	area() float64
}

func info(f figura) float64 {
	return f.area()
}

func main() {
	quadrado1 := quadrado{lado: 5}
	circulo1 := circulo{3}

	fmt.Println("Área do quadrado:", info(quadrado1))
	fmt.Println("Área do círculo:", info(circulo1))
}
