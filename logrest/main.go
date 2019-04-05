// 3) A partir del tipo Matriz anterior, guardar la resta de los valores de
// la matriz en un fichero de logging en disco con el siguiente
// formato: <nombre_operación>,<fecha>,<hora>,<resultado>. Por
// ejemplo, "resta,2018-02-14,18:20,5"

package main

import (
	"log"
	"os"
	"time"
)

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

func main() {
	lf, err := os.OpenFile("logrestfile.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("A ocurrido un error al intentar crear el fichero de log")
	}
	defer lf.Close()
	log.SetOutput(lf)

	m := Matriz{
		{-5.0, 4.6, -0.5},
		{7.6, -9.3, 1.2},
		{-8.3, -4.1, 0.6},
	}
	r := m.Subtraction()
	log.Printf("%s,%s,%v\n", "resta", time.Now().Format("2006-01-02,15:04"), r)
}
