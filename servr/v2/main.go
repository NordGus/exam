// 5) Añadir al servidor anterior una API con la URL "/api/v1/sum" que
// reciba dos números enteros como entrada y devuelva su suma
// como salida

package main

import (
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

	r.HandleFunc("/api/v1/hello", HelloChameleon)
	r.HandleFunc("/api/v1/sum", Sum)

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
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello Chameleon"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Opps!, algo salió mal. Error: %v", err), http.StatusInternalServerError)
		return
	}
}

// Sum Escucha a la URL "/api/v1/sum" esperando dos numeros enteros (a y b) como entrada y devuelve la suma de los mismos como salida
func Sum(w http.ResponseWriter, r *http.Request) {

	ua, ok := r.URL.Query()["a"]
	ub, ok := r.URL.Query()["b"]
	if !ok {
		http.Error(w, "Opps!, algo salió mal. Error: No se recibieron los dos números enteros \"a\" y \"b\"", http.StatusInternalServerError)
		log.Println("No se recibieron los dos números enteros \"a\" y \"b\"")
		return
	}
	a, err := strconv.Atoi(ua[0])
	if err != nil {
		http.Error(w, fmt.Sprintf("Opps!, algo salió mal, Error: Valor de \"%v\" no válido para \"a\".", ua[0]), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	b, err := strconv.Atoi(ub[0])
	if err != nil {
		http.Error(w, fmt.Sprintf("Opps!, algo salió mal, Error: Valor de \"%v\" no válido para \"a\".", ua[0]), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Respondiendo con el resultado de la suma de los dos numeros
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	result := strconv.Itoa(a + b)
	_, err = w.Write([]byte(result))
	if err != nil {
		http.Error(w, fmt.Sprintf("Opps!, algo salió mal. Error: %v", err), http.StatusInternalServerError)
		return
	}
}
