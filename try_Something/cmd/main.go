// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("welcome to index page"))
// 	w.WriteHeader(http.StatusFound)
// }

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("welcome to about page"))
// }

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Merhabalar efenim!\n"))
// 		fmt.Fprintf(w, "Merhaba  merhaba\n")
// 		w.Write([]byte("hey\n")) //---> byte'I vermek zorunluymuş düz string diye yazamıyorsun
// 	})

// 	http.HandleFunc("/index", indexHandler)

// 	http.HandleFunc("/about", aboutHandler)

// 	http.ListenAndServe(":2000", nil)
// }

package main

import (
	"log"
	"net/http"

	. "test/handlers"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server starting...")

	r := mux.NewRouter()

	r.HandleFunc("/api/products", GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/products/{id}", GetProductHandler).Methods("GET")
	r.HandleFunc("/api/products", PostProductHandler).Methods("POST")
	r.HandleFunc("/api/products/{id}", PutProductHandler).Methods("PUT")
	r.HandleFunc("/api/products/{id}", DeleteProductHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":2000",
		Handler: r,
	}

	server.ListenAndServe()
	log.Println("Server ending...")
}
