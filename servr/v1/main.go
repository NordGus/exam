// 4) Programa en Go que ejecute un servidor web y sirva una API con
// la URL "/api/v1/hello" que no reciba ningún parámetro de entrada
// y devuelva "Hello Chameleon" como respuesta.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := fmt.Sprintf(":%v", 8080)

	r := http.NewServeMux()

	r.HandleFunc("/api/v1/hello", RequestLogger(HelloChameleon))

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

// RequestLogger logs incoming request
func RequestLogger(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		log.Printf("%v | [%v] %v - %v \n", r.Proto, r.RemoteAddr, r.Method, r.RequestURI)
		f(w, r)
		log.Printf("[%v] request processed in %s\n", r.RemoteAddr, time.Now().Sub(started))
	}
}
