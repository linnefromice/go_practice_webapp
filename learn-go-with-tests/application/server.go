package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) string
}

type PlayerServer struct {
	store PlayerStore
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}
	if name == "Floyd" {
		return "10"
	}
	return ""
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprintf(w, p.store.GetPlayerScore(player))
}

func main() {
	server := &PlayerServer{}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
