// Module untuk entri utama aplikasi
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MHervian/golang-tasking-webapp/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Membuat route
	route := mux.NewRouter()

	// Daftarkan semua route dari module routes
	routes.RegisterTasksRoutes(route)

	// Memulai server
	fmt.Println("Server started at localhost:9010")
	http.Handle("/", route)

	// Tampilkan jika terjadi error
	log.Fatal(http.ListenAndServe("localhost:9010", route))
}
