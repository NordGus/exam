// 3) A partir del tipo Matriz anterior, guardar la resta de los valores de
// la matriz en un fichero de logging en disco con el siguiente
// formato: <nombre_operación>,<fecha>,<hora>,<resultado>. Por
// ejemplo, "resta,2018-02-14,18:20,5"

package main

import (
	"fmt"
	"log"
	"math/rand"
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
	logger := log.New(os.Stdout, "logrest ", log.LstdFlags|log.Lshortfile)
	lf, err := os.OpenFile("logrestfile.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Println(err)
		os.Exit(7)
	}
	defer lf.Close()

	m := generateRandomMatrix()
	r := m.Subtraction()

	err = writeToLogFile(r, lf)
	if err != nil {
		logger.Println(err)
		os.Exit(7)
	}
}

func generateRandomMatrix() Matriz {
	var m Matriz
	for i, row := range m {
		for j := range row {
			rand.Seed(time.Now().UnixNano())
			mlt := rand.Intn(10)
			if mlt >= 5 {
				m[i][j] = rand.Float64() * float64(mlt)
			} else {
				m[i][j] = -rand.Float64() * float64(mlt)
			}
		}
	}
	return m
}

func writeToLogFile(result float64, logfile *os.File) error {
	_, err := logfile.Write([]byte(fmt.Sprintf("%s,%s,%v\n", "resta", time.Now().Format("2006-01-02,15:04"), result)))
	return err
}
