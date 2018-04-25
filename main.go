package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	startServer()
}

var logger = logrus.New()

var h = herodot.NewJSONWriter(logger)

func startServer() {
	router := httprouter.New()
	router.GET("/api", handleApiCall)


	c := cors.Default().Handler(router)

	if err := http.ListenAndServe(":3001", c); err != nil {
		logrus.WithError(err).Fatal("Unable to start HTTP server")
	}
}

type item struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func handleApiCall(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.Write(w, r, []item{
		{Title: "Hi whatever", URL: "https://github.com"},
		{Title: "LUX rocks", URL: "https://google.com"},
		{Title: "LUX rocks", URL: "https://google.com"},
		{Title: "LUX rocks", URL: "https://google.com"},
		{Title: "LUX rocks", URL: "https://google.com"},
		{Title: "LUX rocks", URL: "https://google.com"},
		{Title: "LUX rocks", URL: "https://google.com"},
		{Title: "LUX rocks", URL: "https://google.com"},
		{Title: "LUX rocks", URL: "https://google.com"},
	})
}
