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
	words = os.Args[1:]
}

func main() {
	if len(words) == 0 {
		log.Println("No hay palabras que escribir al archivo")
		os.Exit(7)
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("A ocurrido un error al intentar crear el fichero <%s>. Error: %v\n", filePath, err)
		os.Exit(7)
	}
	defer file.Close()

	encoder := base64.NewEncoder(base64.StdEncoding, file)
	defer encoder.Close()

	for i, w := range words {
		// Las palabras impares para el usuario tienen posiciones pares en slice de argumentos
		if i%2 == 0 {
			input := fmt.Sprint(w, "\n")
			_, err := encoder.Write([]byte(input))
			if err != nil {
				log.Printf("A ocurrido un error al intentar escribir la palabra %s al fichero %s. Error: %v\n", w, file.Name(), err)
				os.Exit(7)
			}
		}
	}
}
