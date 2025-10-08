package main 

import (
	"net/http"
	"log"
)

func main(){
	const port = "8080"
	const filepathRoot = "."

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	srv := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	log.Printf("serving on port :%s : %s", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}