package main

import (
	"go-web/config"
	admincontroller "go-web/controllers/admincontroller"
	laborcontroller "go-web/controllers/laborcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1. Route Admin
	http.HandleFunc("/admin/", admincontroller.Index)
	http.HandleFunc("/admin/lab", laborcontroller.Index)
	http.HandleFunc("/admin/lab/add", laborcontroller.Add)
	http.HandleFunc("/admin/lab/edit", laborcontroller.Edit)
	http.HandleFunc("/admin/lab/delete", laborcontroller.Delete)


	log.Println("server running on port 8080")
	http.ListenAndServe("localhost:8080", nil)
}