
package main

import (
	"fmt"
	"log"
	"encoding/json"	// To parse POST requests.
	"crypto/md5"		// For Book ID
	"io"
	"time"
	"encoding/hex"
	"crypto/sha256"
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
	BookID        string;  //`json:"book_id"`
	User          string;  //`json:"user"`
	CheckoutDate  string;  //`json:"checkout_date"`
	isGenesis     bool;    //`json:"is_genesis"`
}

type Book struct {
	ID           string;	//`json:"id"`
	Title        string;	//`json:"title"`
	Author       string;	//`json:"author`
	PublishDate  string;	//`json:"publish_date"`
	ISBN         string;	//`json:"isbn"`
}

// Variable used to store all the created blocks.
var blockchain *Blockchain;

func newBook(w http.ResponseWriter, r *http.Request){
	var book Book;

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		log.Printf("New book not generated: %v",err);
		w.Write([]byte("could not create new book"));
		return;
	}

	h := md5.New();
	io.WriteString(h, book.ISBN + book.PublishDate);   // Create a hash of ISBN and date published
	book.ID = fmt.Sprintf("%x", h.Sum(nil));

	resp, err := json.MarshalIndent(book, "", " ");
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError);
		log.Printf("FAILED: JSON Marshall");
	}
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


