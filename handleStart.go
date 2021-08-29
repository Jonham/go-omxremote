package main

import (
	"encoding/base64"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strings"
)

// Start playback http handler
func Start(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Printf("Start %s", ps.ByName("name"))

	addCorsHeader(w)
	filename, _ := base64.URLEncoding.DecodeString(ps.ByName("name"))
	stringFilename := string(filename[:])
	omxOptions := append(strings.Split(omx, " "), stringFilename)

	err := p.Start(omxOptions)
	if err != nil {
		p.Playing = false
		http.Error(w, err.Error(), 500)
		return
	}

	log.Printf("Playing media file: %s\n", stringFilename)
	w.WriteHeader(http.StatusOK)
}
