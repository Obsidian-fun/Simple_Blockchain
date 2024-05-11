
package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func ServeHTTP (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome to port %d!\n",3000);
}


func main() {
	fmt.Println("Hello World!");

	r := mux.NewRouter();
	r.HandleFunc("/", ServeHTTP).Methods("GET");

	http.ListenAndServe(":3000",r);
}


