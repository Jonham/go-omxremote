package main

import "net/http"

func addCorsHeader(writer http.ResponseWriter) {
	headers := writer.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
}
