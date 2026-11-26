package main

import (
	"fmt"
	"math"
)

// 1. Interface Figura
type Figura interface {
	Area() float64
	Perimetro() float64
	Nombre() string
}

// 2. Struct Rectangulo
type Rectangulo struct {
	Largo float64
	Ancho float64
}

func (r Rectangulo) Nombre() string {
	return "Rectángulo"
}

func (r Rectangulo) Area() float64 {
	return r.Largo * r.Ancho
}

func (r Rectangulo) Perimetro() float64 {
	return 2 * (r.Largo + r.Ancho)
}

// 3. Struct Circulo
type Circulo struct {
	Radio float64
}

func (c Circulo) Nombre() string {
	return "Círculo"
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

// 4. Struct Triangulo (lados a, b, c)
type Triangulo struct {
	A float64
	B float64
	C float64
}

func (t Triangulo) Nombre() string {
	return "Triángulo"
}

func (t Triangulo) Perimetro() float64 {
	return t.A + t.B + t.C
}

// Fórmula de Herón: s = (a + b + c)/2 -> sqrt(s*(s-a)*(s-b)*(s-c))
func (t Triangulo) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

// 5. Función ImprimirFigura que recibe la interfaz Figura
func ImprimirFigura(f Figura) {
	fmt.Printf("Figura: %s | Área: %.2f | Perímetro: %.2f\n", f.Nombre(), f.Area(), f.Perimetro())
}

// 6. Función main() creando 3 figuras y llamando a ImprimirFigura()
func main() {
	rect := Rectangulo{Largo: 3, Ancho: 4}
	circ := Circulo{Radio: 5}
	tri := Triangulo{A: 3, B: 4, C: 5}

	figuras := []Figura{rect, circ, tri}

	for _, fig := range figuras {
		ImprimirFigura(fig)
	}
}
