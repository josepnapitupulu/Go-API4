package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/josepnapitupulu/Project_pasti/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterJemaatsRoutes(r)
	http.Handle("/", r)
	http.Handle("/jemaatBaru", r)

	fmt.Print("Starting Server localhost:7070")
	log.Fatal(http.ListenAndServe("localhost:7070", r))
}