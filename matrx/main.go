// 2) Programa en Go que implemente un tipo "Matriz", que represente
// una matriz de 3 x 3, y las siguientes operaciones sobre ella:
// "transpuesta" de la matriz, y "suma" de todos sus valores.

package main

import "fmt"

func main() {
	m := Matriz{
		{1.0, 2.0, 3.0},
		{4.0, 5.0, 6.0},
		{7.0, 8.0, 9.0},
	}

	t := m.Transposed()
	sm := m.Sum()

	fmt.Println("Matriz:", m)
	fmt.Println("Transpuesta de la Matriz:", t)
	fmt.Println("Suma de todos los valores de la matriz:", sm)
}
