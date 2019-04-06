// 2) Programa en Go que implemente un tipo "Matriz", que represente
// una matriz de 3 x 3, y las siguientes operaciones sobre ella:
// "transpuesta" de la matriz, y "suma" de todos sus valores.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Matriz es un tipo alias que represanta una matriz de 3x3 en memoria
type Matriz [3][3]int

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
func (m Matriz) Sum() int {
	var t int
	for _, row := range m {
		for _, value := range row {
			t += value
		}
	}
	return t
}

func main() {
	m := generateRandomMatrix()

	t := m.Transposed()
	sm := m.Sum()

	fmt.Println("Matriz:", m)
	fmt.Println("Transpuesta de la Matriz:", t)
	fmt.Println("Suma de todos los valores de la matriz:", sm)
}

func generateRandomMatrix() Matriz {
	var m Matriz
	min := -10
	max := 10
	for i, row := range m {
		for j := range row {
			rand.Seed(time.Now().UnixNano())
			m[i][j] = rand.Intn(max-min) + min
		}
	}
	return m
}
