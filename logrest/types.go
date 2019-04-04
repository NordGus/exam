package main

// Matriz es un tipo alias que represanta una matriz de 3x3 en memoria
type Matriz [3][3]float64

// Subtraction es un metodo que retorna la resta de todo los elementos de una Matriz dada
func (m Matriz) Subtraction() float64 {
	var t float64
	for _, row := range m {
		for _, value := range row {
			t -= value
		}
	}
	return t
}
