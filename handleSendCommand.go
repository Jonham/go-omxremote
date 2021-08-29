package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// SendCommand is the HTTP handler for sending a command to the player
func SendCommand(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	addCorsHeader(w)

	err := p.SendCommand(ps.ByName("command"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
