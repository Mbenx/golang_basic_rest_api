package main

import (
	"golang_basic_rest_api/routes"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/employees", routes.GetEmployee)
	http.HandleFunc("/post-employee", routes.PostEmployee)
	http.HandleFunc("/put-employee", routes.PutEmployee)
	http.HandleFunc("/delete-employee", routes.DeleteEmployee)

	http.HandleFunc("/departments", routes.GetDepartment)
	http.HandleFunc("/post-department", routes.PostDepartment)
	http.HandleFunc("/put-department", routes.PutDepartment)
	http.HandleFunc("/delete-department", routes.DeleteDepartment)

	log.Printf("Server Run In Localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("testing kamu berada di home"))
}
