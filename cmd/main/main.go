package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/taninchot0919/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	server := "localhost:9000"
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("The server is running at :" + server)
	log.Fatal(http.ListenAndServe(server, r))
}
