package main

import "github.com/julienschmidt/httprouter"

func initRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/files.json", handleListVideo)

	router.POST("/start/:name", Start)
	router.POST("/player/:command", SendCommand)

	router.ServeFiles("/app/*filepath", FS(false))

	return router
}
