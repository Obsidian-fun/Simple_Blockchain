
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
	BookID        string	`json:"book_id"`
	User          string	`json:"user"`
	CheckoutDate  string	`json:"checkout_date"`
	isGenesis       bool	`json:"is_genesis"`
}

type Book struct {
	ID           string		`json:"id"`
	Title        string		`json:"title"`
	Author       string		`json:"author`
	PublishDate  string		`json:"publish_date"`
	ISBN         string		`json:"isbn"`
}

// Variable used to store all the created blocks.
var blockchain *Blockchain;

func (b *Block)generateHash(){
	bytes, _ := json.Marshal(b.Data);

	data := string(b.Pos) + b.TimeStamp + string(bytes) + b.PreviousHash;

	hash := sha256.New();
	hash.Write([]byte(data));
	b.Hash = hex.EncodeToString(hash.Sum(nil));

}


func CreateBlock(prevBlock *Block, checkoutItem BookCheckout) *Block {
	block := &Block{}; // Dereference * Block
	block.Pos = prevBlock.Pos + 1;
	block.TimeStamp = time.Now().String();
	block.PreviousHash = prevBlock.Hash
	block.generateHash();

	return block;
}

func (bc *Blockchain) AddBlock(data BookCheckout) {
	prevBlock := bc.blocks[len(bc.blocks)-1];

	block := CreateBlock(prevBlock, data);

	if validBlock(block, prevBlock) {
		bc.blocks = append(bc.blocks, block);
	}
}

func validBlock(block *Block, PrevBlock *Block) bool{

	// Check hash matches,
	if PrevBlock.Hash != block.PreviousHash {
		return false;
	}

	//Validate hash
	if !block.validateHash(block.Hash){
		return false;
	}

	// Check position is correct 
	if PrevBlock.Pos + 1 != block.Pos  {
		return false;
	}

	return true;
}


func writeBlock(w http.ResponseWriter, r *http.Request) {
	var checkoutItem BookCheckout;

	if err := json.NewDecoder(r.Body).Decode(&checkoutItem); err != nil {
		w.WriteHeader(http.StatusInternalServerError);
		log.Printf("Book Checkout not processed: %v",err);
		w.Write([]byte("Could not process transaction"));
		return;
	}

	blockchain.AddBlock(checkoutItem);

}

func newBook(w http.ResponseWriter, r *http.Request){
	var book Book;

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError);
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
		log.Printf("FAILED: JSON Marshal");
		w.Write([]byte("Could not save BookID"));
		return ;
	}

	w.WriteHeader(http.StatusOK);
	w.Write(resp);
}


func getBlockchain(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome to port %d!\n",3000);
}

func GenesisBlock() *Block {
	return CreateBlock(&Block{}, BookCheckout{isGenesis:true})

}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}};
}

func main() {

	blockchain = NewBlockchain();

	r := mux.NewRouter();
	r.HandleFunc("/", getBlockchain).Methods("GET");  // Get all blocks in the blockchain
	r.HandleFunc("/", writeBlock).Methods("POST");    // Write a transaction into a block
	r.HandleFunc("/new", newBook).Methods("POST");    // Create a new BookID

	log.Fatal(http.ListenAndServe(":3000",r));
}


