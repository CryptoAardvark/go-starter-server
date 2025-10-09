package main 

import (
	"net/http"
	"log"
)

type server struct{
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, from server!"))
}

func main(){
	s := &server{addr: ":5000"}
	if err := http.ListenAndServe(a.addr, s); err != nil {
		log.Fatal(err)
	}
}
