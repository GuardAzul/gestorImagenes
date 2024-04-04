package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo")
}

func main() {
	// Manejador para la ruta "/"
	http.HandleFunc("/", handler)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en el puerto 5000...")
	http.ListenAndServe(":5000", nil)
}
