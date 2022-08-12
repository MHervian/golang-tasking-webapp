// Module for routes in http response
package routes

import (
	"github.com/MHervian/golang-tasking-webapp/pkg/controllers"
	"github.com/gorilla/mux"
)

// Fungsi anonim untuk registrasi routes
var RegisterTasksRoutes = func(router *mux.Router) {
	// Methods menampilkan halaman halaman
	router.HandleFunc("/", controllers.Index).Methods("GET")
	router.HandleFunc("/form", controllers.Form).Methods("GET")
	router.HandleFunc("/task/id", controllers.DisplayTaskById).Methods("GET")
	router.HandleFunc("/task/edit/id", controllers.EditTaskById).Methods("GET")

	// Methods CRUD data task
}
