package main 

import (
	"net/http"
	"log"
)

type server struct{
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index pages"))
			return
		case "/users":
			w.Write([]byte("users pages"))
			return
		default:
			w.Write([]byte("404 page"))
		}
	default:
		w.Write([]byte("404 page"))
		return
	}
}

func main(){
	s := &server{addr: ":5000"}
	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal(err)
	}
}
