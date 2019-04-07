package main

import(
  "log"
  "net/http"
  "encoding/json"
  "math/rand"
  "strconv"
  "github.com/gorilla/mux"
)

// Book struct model
type Book struct {
  ID    string   `json:"id"`
  Isbn  string   `json:"isbn"`
  Title string   `json:"title"`
  Author *Author `json:"author"`
}

// Author struct
type Author struct{
  Firstname string `json:"firstname"`
  Lastname  string `json:"lastname"`
}

// Inittialize a books var as a slice Book struct
var books []Book

// Get all getBooks
func getBooks(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(books)

}

// Get getBook
func getBook(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)  // get params
  // loop through books and find the id
  for _, item := range books {
    if item.ID == params["id"]{
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Book{})
}

// Get create Book
func createBook(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  var book Book
  _= json.NewDecoder(r.Body).Decode(&book)
  // mock id
  book.ID = strconv.Itoa(rand.Intn(1000000))
  books = append(books, book)
  json.NewEncoder(w).Encode(book)
}

// Get update Book
func updateBook(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)  // get params
  // loop through books and find the id
  for index, item := range books {
    if item.ID == params["id"]{
      books = append(books[:index], books[index+1:]...)
      var book Book
      _= json.NewDecoder(r.Body).Decode(&book)
      // mock id
      book.ID = strconv.Itoa(rand.Intn(1000000))
      books = append(books, book)
      json.NewEncoder(w).Encode(book)
      return
    }
  }
  json.NewEncoder(w).Encode(books)
}

// Get delete Book
func deleteBook(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)  // get params
  // loop through books and find the id
  for index, item := range books {
    if item.ID == params["id"]{
      books = append(books[:index], books[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(books)
}

func main()  {

  // init router
  r := mux.NewRouter()

  // Mock Data
  books = append(books, Book{ID: "1", Isbn: "232321", Title: "My Book one", Author: &Author{
    Firstname: "Mano",
    Lastname: "Chandran",
    }})
    books = append(books, Book{ID: "2", Isbn: "232333", Title: "My Book two", Author: &Author{
      Firstname: "Mano",
      Lastname: "Chandran",
      }})

  // Route handlers / endpoints
  r.HandleFunc("/api/books", getBooks).Methods("GET")
  r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
  r.HandleFunc("/api/books", createBook).Methods("POST")
  r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
  r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

  // run the server
  log.Fatal(http.ListenAndServe(":8000", r))
}
