
package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"github.com/gorilla/mux"
)

type Blockchain struct {
	blocks []*Block

}

type BookCheckout struct {
	user          string;
	checkoutDate  time;
	isGenesis     bool;
}


type Book struct {
	id           int;
	title        string;
	author       string;
	ISBN         int;
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

	log.Fatal(http.ListenAndServe(":3000",r));
}


