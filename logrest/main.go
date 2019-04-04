// 3) A partir del tipo Matriz anterior, guardar la resta de los valores de
// la matriz en un fichero de logging en disco con el siguiente
// formato: <nombre_operación>,<fecha>,<hora>,<resultado>. Por
// ejemplo, "resta,2018-02-14,18:20,5"

package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func init() {
	lf, err := os.Create(fmt.Sprintf("%s.log", time.Now().Format("20060102150405")))
	if err != nil {
		log.Fatalln("A ocurrido un error al intentar crear el fichero de log")
	}
	log.SetOutput(lf)
}

func main() {
	m := Matriz{
		{-5.0, 4.6, -0.5},
		{7.6, -9.3, 1.2},
		{-8.3, -4.1, 0.6},
	}
	r := m.Subtraction()
	log.Printf("%s,%s,%v\n", "resta", time.Now().Format("2006-01-02,15:04"), r)
}
