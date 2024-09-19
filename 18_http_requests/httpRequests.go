package main

import (
	"fmt"
	"io"
	"net/http"
)

const Url = "https://www.google.com"

func main() {
	fmt.Println("All About Http Requests")
	response, err := http.Get(Url)
	CheckNilError(err)
	fmt.Printf("Type of response is %T\n", response) // Type of response is *http.Response
	defer response.Body.Close()	// Callers responsibility to close the connection

	byte_data, err := io.ReadAll(response.Body)
	CheckNilError(err)
	content := string(byte_data)
	fmt.Println("Content is: ", content)
}

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}



// import (
// 	"net/http"
// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/chi/v5/middleware"
// )

// func main() {
// 	r := chi.NewRouter()
// 	r.Use(middleware.Logger)
// 	r.Use(middleware.Recoverer)
// 	r.Get(Url, func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("root."))
// 	})
// 	http.ListenAndServe(":3333", r)
// }