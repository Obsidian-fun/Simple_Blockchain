
package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

type book struct {
	id           int;
	title        string;
	author       string;
	ISBN         int;
	PublishDate  Time.Now();
}

func getBlockchain(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome to port %d!\n",3000);
}

func main() {
	fmt.Println("Hello World!");

	r := mux.NewRouter();
	r.HandleFunc("/", getBlockchain).Methods("GET");
	r.HandleFunc("/", writeBlock).Methods("POST");
	r.HandleFunc("/new", newBook).Methods("POST");

	http.ListenAndServe(":3000",r);
}


