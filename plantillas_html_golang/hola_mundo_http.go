// https://www.youtube.com/watch?v=5t2jPn5PUMc
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("pagina.html")

	if err != nil {
		fmt.Fprint(w, "Página no encontrada")
	} else {
		template.Execute(w, nil)
	}
	fmt.Fprintf(w, "Hola Mundo")
}

func main() {
	// Manejador para la ruta "/"
	http.HandleFunc("/", handler)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en el puerto 5000...")
	http.ListenAndServe(":5000", nil)
}
