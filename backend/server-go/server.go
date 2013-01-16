package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

func getParams(r *http.Request) (string, string, string) {

	params := r.URL.Query()
	user := params.Get(":user")
	topic := params.Get(":topic")
	key := user + "@" + topic
	return user, topic, key
}

func getSlide(w http.ResponseWriter, r *http.Request) {
	user, topic, key := getParams(r)
	fmt.Fprintf(w, "you are get user %s topic %s key %s", user, topic, key)
}

func postSlide(w http.ResponseWriter, r *http.Request) {
	user, topic, key := getParams(r)
	fmt.Fprintf(w, "you are get user %s topic %s key %s", user, topic, key)
}

func putSlide(w http.ResponseWriter, r *http.Request) {
	user, topic, key := getParams(r)
	fmt.Fprintf(w, "you are get user %s topic %s key %s", user, topic, key)
}

func main() {
	mux := routes.New()
	mux.Get("/:user/:topic/", getSlide)
	mux.Post("/:user/:topic/", postSlide)
	mux.Put("/:user/:topic/", putSlide)
	http.Handle("/", mux)
	http.ListenAndServe(":7788", nil)
}
