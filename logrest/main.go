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
type Matriz [3][3]int

// Subtraction es un metodo que retorna la resta de todo los elementos de una Matriz dada
func (m Matriz) Subtraction() int {
	var t int
	for _, row := range m {
		for _, value := range row {
			t -= value
		}
	}
	return t
}

// SubstractAndLog es un metodo que calcula la resta de todos los valores de una Matriz
// y guarda el resultado en un fichero de log en disco (filename) con el formato:
// <nombre_operación>,<fecha>,<hora>,<resultado>. Y devuelve un error en cado de que algo halla fallado
func (m Matriz) SubstractAndLog(filename string) error {
	r := m.Subtraction()
	lf, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer lf.Close()
	_, err = lf.Write([]byte(fmt.Sprintf("%s,%s,%v\n", "resta", time.Now().Format("2006-01-02,15:04"), r)))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	logger := log.New(os.Stdout, "", 0)
	m := generateRandomMatrix()
	err := m.SubstractAndLog("logrestfile.log")
	if err != nil {
		logger.Println(err)
		os.Exit(7)
	}
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
