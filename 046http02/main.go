package main

import "net/http"

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/h", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World<hr>"))
	})

	http.ListenAndServe(":8080", nil)

}
