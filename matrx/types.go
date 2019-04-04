// 2) Programa en Go que implemente un tipo "Matriz", que represente
// una matriz de 3 x 3, y las siguientes operaciones sobre ella:
// "transpuesta" de la matriz, y "suma" de todos sus valores.

package main

// Matriz es un tipo alias que represanta una matriz de 3x3 en memoria
type Matriz [3][3]float64

// Transposed es un metodo que retorna la Matriz transpuesta de una Matriz dada
func (m Matriz) Transposed() Matriz {
	t := Matriz{}
	for i, row := range m {
		for j, value := range row {
			t[j][i] = value
		}
	}
	return t
}

// Sum es un metodo que retorna la suma de todo los elementos de una Matriz dada
func (m Matriz) Sum() float64 {
	var t float64
	for _, row := range m {
		for _, value := range row {
			t += value
		}
	}
	return t
}
