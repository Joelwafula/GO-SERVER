package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset= utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome channel</h1>")

}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Tyype", "text/plain; charset = utf-8")

	fmt.Fprint(w, "<h1>contact </h1><p>to get in touch email me at  <a href=\"mailto:joel@gmail.com\" </a> </p>")

}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path == "/" {
// 		handlerFunc(w, r)

// 	} else if r.URL.Path == "/contacts" {
// 		contactHandler(w, r)
// 	}
// }

//like the above commented code, you can also implement it using the switch case, see bellow

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		handlerFunc(w, r)
// 	case "/contacts":
// 		contactHandler(w, r)

// 	default:
// 		// w.WriteHeader(http.StatusNotFound)
// 		// fmt.Println(w, "page not found")
// 		http.Error(w, "page not found", http.StatusNotFound)

// 	}
// }

type Router struct{}

// main benefit of using ths typeof struct, is that our handler function can have multiple access eg to DB
func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handlerFunc(w, r)
	case "/contacts":
		contactHandler(w, r)

	default:
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Println(w, "page not found")
		http.Error(w, "page not found", http.StatusNotFound)

	}

}

// theabove function is used to print default path

func main() {

	var router Router
	fmt.Println("Starting server on port 3000")
	http.ListenAndServe("3000", router)

	// ///we can also use the bellow commented code
	// // mux := http.NewServeMux()
	// // mux.HandleFunc("/", hanlerFunc)

	// //http.HandleFunc("/", pathHandler)
	// //this function prints the default path from theabove function

	// http.HandleFunc("/", handlerFunc)
	// http.HandleFunc("/contacts", contactHandler)
	// fmt.Println("Starting server on port 3000")
	// err := http.ListenAndServe(":3000", nil)

	// if err != nil {
	// 	panic(err)
	// }
	// //http.ListenAndServe(":3000", mux)
}
