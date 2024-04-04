// https://www.youtube.com/watch?v=5t2jPn5PUMc
package main

import (
	b64 "encoding/base64"
	"fmt"
	"html/template"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"strings"
)

var (
	puerto   string
	tema     string
	hostname string
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

type DatosPagina struct {
	Theme    string
	Images   []ImagenBase64
	HostName string
}

type ImagenBase64 struct {
	Encoding template.URL
	Nombre   string
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		fmt.Fprint(w, "Página no encontrada")
	} else {
		// Obtiene el nombre del primer argumento mandado a través de la línea de comandos
		carpeta := os.Args[1]
		directorio, err := os.Open(carpeta)
		check(err)
		defer directorio.Close()
		nombres, err := directorio.Readdirnames(0)
		check(err)

		var archivos []string
		for _, nombre := range nombres {

			if strings.HasSuffix(nombre, ".jpg") || strings.HasSuffix(nombre, ".png") ||
				strings.HasSuffix(nombre, ".jpeg") {
				archivos = append(archivos, nombre)
			}
		}

		fmt.Println("Cantidad de archivos en la carpeta: ", len(archivos))

		var imagen_aleatoria = archivos[rand.IntN(len(archivos)-1)]

		fmt.Println(imagen_aleatoria)

		// Obtener el nombre del host del sistema operativo
		hostname, err := os.Hostname()
		check(err)

		fmt.Println("Nombre del host: ", hostname)

		var listaGenerada []ImagenBase64
		for i := 0; i < 4; i++ {
			var imagen_aleatoria = archivos[rand.IntN(len(archivos)-1)]

			f, err := os.ReadFile(carpeta + imagen_aleatoria)
			check(err)

			var src = "data:image/jpg;base64," + b64.StdEncoding.EncodeToString(f)

			image := ImagenBase64{
				Encoding: template.URL(src),
				Nombre:   imagen_aleatoria,
			}

			listaGenerada = append(listaGenerada, image)
		}

		data := DatosPagina{
			Theme:    tema,
			Images:   listaGenerada,
			HostName: hostname,
		}

		// Ejecutar la plantilla y escribir el resultado en la respuesta HTTP
		err = tmpl.Execute(w, data)
	}
}

func main() {
	fmt.Print("Ingrese el puerto: ")
	fmt.Scan(&puerto)
	fmt.Print("Ingrese el tema: ")
	fmt.Scan(&tema)

	// Manejador para la ruta "/"
	http.HandleFunc("/", handler)

	// Iniciar el servidor en el puerto especificado
	fmt.Println("Servidor escuchando en el puerto 5000...")
	log.Fatal(http.ListenAndServe(":"+puerto, nil))
}
