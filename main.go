
package main

import (
	"fmt"
	"log"
	"encoding/json"

	"net/http"
	"github.com/gorilla/mux"
)

type Block struct {
	Pos           int;
	Data          BookCheckout; // `Data  is  an  instance  of  BookCheckout  struct`
	TimeStamp     string;
	Hash          string;
	PreviousHash  string;
}

type Blockchain struct {
	blocks []*Block

}

type BookCheckout struct {
	BookID        string;  `json:"book_id"`
	User          string;  `json:"user"`
	CheckoutDate  string;  `json:"checkout_date"`
	isGenesis     bool;    `json:"is_genesis"`
}

type Book struct {
	ID           string;	`json:"id"`
	Title        string;	`json:"title"`
	Author       string;	`json:"author`
	PublishDate  string;	`json:"publish_date"`
	ISBN         string;	`json:"isbn"`
}

// Variable used to store all the created blocks.
var Blockchain *Blockchain;

func newBook(w http.ResponseWriter, r *http.Request){
	var book Book;

	if err := json.NewDecoder(r.Body);

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


