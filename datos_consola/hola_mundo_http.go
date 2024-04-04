// https://www.youtube.com/watch?v=5t2jPn5PUMc
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type DatosPagina struct {
	Theme    string
	ImageDir string
	HostName string
}

func handler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("index.html")

	if err != nil {
		fmt.Fprint(w, "Página no encontrada")
	} else {

		// Obtener el nombre del host del sistema operativo
		hostname, error := os.Hostname()

		if error != nil {
			fmt.Println("Error al obtener el nombre del host: ", error)
			return
		}

		data := DatosPagina{
			Theme:    tema,
			ImageDir: carpeta,
			HostName: hostname,
		}

		// Ejecutar la plantilla y escribir el resultado en la respuesta HTTP
		err = template.Execute(w, data)
	}
}

var (
	puerto   string
	carpeta  string
	tema     string
	hostname string
)

func main() {
	fmt.Print("Ingrese el puerto: ")
	fmt.Scan(&puerto)
	fmt.Print("Ingrese la carpeta: ")
	fmt.Scan(&carpeta)
	fmt.Print("Ingrese el tema: ")
	fmt.Scan(&tema)

	// Manejador para la ruta "/"
	http.HandleFunc("/", handler)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en el puerto 5000...")
	http.ListenAndServe(":"+puerto, nil)
}
