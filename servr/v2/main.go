// 5) Añadir al servidor anterior una API con la URL "/api/v1/sum" que
// reciba dos números enteros como entrada y devuelva su suma
// como salida

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	addr := fmt.Sprintf(":%v", 8080)

	r := http.NewServeMux()

	r.HandleFunc("/api/v1/hello", LogIncomingRequest(HelloChameleon))
	r.HandleFunc("/api/v1/sum", LogIncomingRequest(Sum))

	s := http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := s.ListenAndServe()

	if err != nil {
		log.Println(err)
		os.Exit(7)
	}
}

// HelloChameleon sirve a la URL "/api/v1/hello" y retorna "Hello Chamalleon" como respuesta
func HelloChameleon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello Chameleon"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Opps!, algo salió mal. Error: %v", err), http.StatusInternalServerError)
		return
	}
}

// Sum Escucha a la URL "/api/v1/sum" esperando dos numeros enteros (a y b) como entrada y devuelve la suma de los mismos como salida
func Sum(w http.ResponseWriter, r *http.Request) {
	a, b, err := retrieveNumbers(r.URL.Query())
	if err != nil {
		http.Error(w, fmt.Sprint("Opps algo salió mal. Error: ", err), http.StatusInternalServerError)
		return
	}

	// Respondiendo con el resultado de la suma de los dos numeros
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	result := strconv.Itoa(a + b)
	_, err = w.Write([]byte(result))
	if err != nil {
		http.Error(w, fmt.Sprintf("Opps!, algo salió mal. Error: %v", err), http.StatusInternalServerError)
		return
	}
}

// LogIncomingRequest logs incoming request
func LogIncomingRequest(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v | [%v] %v - %v \n", r.Proto, r.RemoteAddr, r.Method, r.RequestURI)
		f(w, r)
	}
}

// retrieveNumbers procesa los parametros de la url y devuelve los numeros a sumar o un error en caso de que algo salga mal
func retrieveNumbers(values map[string][]string) (int, int, error) {
	var numbers []int
	if len(values) != 2 {
		err := errors.New("no se recibieron los dos números enteros en la URL")
		log.Println(err)
		return 0, 0, err
	}
	for key, value := range values {
		number, err := strconv.Atoi(value[0])
		if err != nil {
			err := fmt.Errorf("\"%v\" no es un valor válido para \"%v\"", value[0], key)
			log.Println(err)
			return 0, 0, err
		}
		numbers = append(numbers, number)
	}
	return numbers[0], numbers[1], nil
}
