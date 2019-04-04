// 1) Programa en Go que lea de consola una lista de palabras
// introducidas por un usuario y escriba en el directorio "/tmp" un
// fichero con las palabras impares del usuario en base 64

package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

const (
	filePath = "/tmp/listr64.txt"
)

var words []string

func init() {
	// Ignorando el primer elemento del slice que regresa os.Args,
	// debido a que es el nombre del programa
	words = os.Args[1:]
}

func main() {
	// Retornando temprano en caso de que no hallan palabras que escribir al fichero
	if len(words) == 0 {
		log.Fatalln("No hay palabras que escribir al archivo")
	}

	// Creando fichero en el directorio /tmp en el cual se van a escribir las palabras
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("A ocurrido un error al intentar crear el fichero <%s>. Error: %v", filePath, err)
	}
	defer file.Close()

	encoder := base64.NewEncoder(base64.StdEncoding, file)
	defer encoder.Close()

	// Procesando el listado de palabras recibidas
	for i, w := range words {
		// Las palabras impares para el usuario tienen posiciones pares en slice
		if i%2 == 0 {
			_, err := encoder.Write([]byte(fmt.Sprint(w, "\n")))
			if err != nil {
				log.Fatalf("A ocurrido un error al intentar escribir la palabra %s al fichero %s. Error: %v", w, file.Name(), err)
			}
		}
	}
}
