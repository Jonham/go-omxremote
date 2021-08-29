package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Page is the HTML page struct
type Page struct {
	Title     string
	Timestamp int64
}

// Index func that serves the HTML for the home page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p := &Page{Title: "go-OMX Remote", Timestamp: time.Now().Unix()}
	tmpl, err := FSString(false, "/dist/index.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	t, _ := template.New("index").Parse(tmpl)
	t.Execute(w, p)
}
