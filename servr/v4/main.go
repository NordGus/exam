// 7) Añadir al servidor anterior una API con la URL "/api/v1/reset" que
// borre todas las sumas y sus resultados almacenados en base de
// datos. Como respuesta esta API no devuelve contenido ninguno.

package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbname = "./servr.db"
	driver = "sqlite3"
)

func init() {
	err := createSumTableIfDoesntExist()
	if err != nil {
		log.Println(err)
		os.Exit(7)
	}
}

func main() {
	addr := fmt.Sprintf(":%v", 8080)

	r := http.NewServeMux()

	r.HandleFunc("/api/v1/hello", HelloChameleon)
	r.HandleFunc("/api/v1/sum", Sum)
	r.HandleFunc("/api/v1/sumdb", SumDB)
	r.HandleFunc("/api/v1/reset", Reset)

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

// HelloChameleon escucha a la URL "/api/v1/hello" y retorna "Hello Chamalleon" como respuesta.
func HelloChameleon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello Chameleon"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Opps!, algo salió mal. Error: %v", err), http.StatusInternalServerError)
		return
	}
}

// Sum escucha a la URL "/api/v1/sum" esperando dos numeros enteros, "a" y "b", como entrada
// y devuelve la suma de los mismos como salida.
func Sum(w http.ResponseWriter, r *http.Request) {
	a, b, err := retrieveNumbers(r.URL.Query())
	if err != nil {
		http.Error(w, fmt.Sprint("Opps algo salió mal. Error: ", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	result := strconv.Itoa(a + b)
	_, err = w.Write([]byte(result))
	if err != nil {
		http.Error(w, fmt.Sprintf("Opps!, algo salió mal. Error: %v", err), http.StatusInternalServerError)
		return
	}
}

// SumDB escucha a la URL "/api/v1/sumdb" esperando dos numeros enteros, "a" y "b", como entrada,
// calcula la suma de los mismos, guarda ambos números y el resultado de la suma en la base de datos
// y retorna el número de resultados en la base de datos.
func SumDB(w http.ResponseWriter, r *http.Request) {
	a, b, err := retrieveNumbers(r.URL.Query())
	if err != nil {
		http.Error(w, fmt.Sprint("Opps algo salió mal. Error: ", err), http.StatusInternalServerError)
		return
	}
	total, err := registerSumInDB(a, b)
	if err != nil {
		http.Error(w, fmt.Sprint("Opps algo salió mal. Error: ", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	result := strconv.Itoa(total)
	_, err = w.Write([]byte(result))
	if err != nil {
		http.Error(w, fmt.Sprintf("Opps!, algo salió mal. Error: %v", err), http.StatusInternalServerError)
		return
	}
}

// Reset escucha a la URL "/api/v1/reset" borra todos los registros de sumas y sus resultados
// de la base de datos y no devuelve contenido.
func Reset(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(driver, dbname)
	if err != nil {
		http.Error(w, fmt.Sprint("Opps algo salió mal. Error: ", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	result, err := db.Exec("DELETE FROM sums;")
	if err != nil {
		http.Error(w, fmt.Sprint("Opps algo salió mal. Error: ", err), http.StatusInternalServerError)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprint("Opps algo salió mal. Error: ", err), http.StatusInternalServerError)
		return
	}
	log.Println("Número de operaciones borradas:", affected)
	w.WriteHeader(http.StatusNoContent)
}

// retrieveNumbers procesa los parametros de la url y devuelve los numeros a sumar o un error en caso de que algo falle
func retrieveNumbers(values map[string][]string) (int, int, error) {
	var numbers []int
	if len(values) != 2 {
		return 0, 0, errors.New("no se recibieron los dos números enteros en la URL")
	}
	for key, value := range values {
		number, err := strconv.Atoi(value[0])
		if err != nil {
			return 0, 0, fmt.Errorf("\"%v\" no es un valor válido para \"%v\"", value[0], key)
		}
		numbers = append(numbers, number)
	}
	return numbers[0], numbers[1], nil
}

// CreateSumTableIfDoesntExist es una función para que en el momento de la inicialización del programa
// que revisa si exista la base de datos y la tabla necesaria dentro de ella
func createSumTableIfDoesntExist() error {
	db, err := sql.Open(driver, dbname)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS sums (id INTEGER PRIMARY KEY AUTOINCREMENT, first_number INTEGER NOT NULL, second_number INTEGER NOT NULL, total INTEGER NOT NULL);`)
	if err != nil {
		return err
	}
	return nil
}

// registerSumInDB toma dos números enteros, "a" y "b", calcula la suma de los mismos,
// guarda ambos números y el resultado de la suma en la base de datos,
// calcula el número de resultados en la base de datos y lo retorna o un error en caso de que algo falle
func registerSumInDB(a, b int) (int, error) {
	db, err := sql.Open(driver, dbname)
	if err != nil {
		return 0, err
	}
	defer db.Close()
	insert, err := db.Prepare("INSERT INTO sums(first_number, second_number, total) VALUES (?, ?, ?);")
	if err != nil {
		return 0, err
	}
	defer insert.Close()
	_, err = insert.Exec(a, b, (a + b))
	if err != nil {
		return 0, err
	}
	count, err := db.Query("SELECT COUNT(total) FROM sums;")
	if err != nil {
		return 0, err
	}
	defer count.Close()
	var total int
	for count.Next() {
		err = count.Scan(&total)
		if err != nil {
			return 0, err
		}
		break
	}
	db.Close()
	insert.Close()
	count.Close()
	return total, nil
}
