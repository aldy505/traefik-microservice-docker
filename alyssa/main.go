package main

import "net/http"

func main() {
	http.HandleFunc("/alyssa", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Alyssa!"))
	})

	http.HandleFunc("/pouring-alyssa", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("We are pouring on Alyssa. (NB. It's a flower)"))
	})

	http.ListenAndServe(":8080", nil)
}