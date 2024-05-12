
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
	user         
	checkoutDate 
	isGenesis    
}


type Book struct {
	ID           string;	`json:"id:"`
	Title        string;	`json:"title:"`
	Author       string;	`json:"author:`
	PublishDate  string;	`json:"publish_date"`
	ISBN         string;	`json:"isbn:"`
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


