package main

import "net/http"

func main() {
	// router := mux.NewRouter()
	// router.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
	// http.ListenAndServe(":8080", router)

	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":8080", nil)
}
