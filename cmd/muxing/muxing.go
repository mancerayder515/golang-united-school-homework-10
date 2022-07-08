package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{n}", namesHandler)
	router.HandleFunc("/bad", badRequestHandler)
	router.HandleFunc("/data", bodyHandler).Methods("POST")
	router.HandleFunc("/headers", headersHandler)
	
	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func namesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, fmt.Sprintf("Hello, %s!", strings.Trim(vars["n"], "{}")))
}

func badRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func bodyHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		body = []byte(err.Error())
	}
	w.Write([]byte(fmt.Sprintf("I got message:\n%s", body)))
}

func headersHandler(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header["A"][0])
	if err != nil {
		fmt.Println(err)
	}
	b, _ := strconv.Atoi(r.Header["B"][0])
	w.Header().Add("a+b", fmt.Sprint(a+b))
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
