package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

func getParams(r *http.Request) (string, string) {

	params := r.URL.Query()
	user := params.Get(":user")
	topic := params.Get(":topic")
	return user, topic
}

func genKey(user string, topic string) (key string) {
	key := user + "@" + topic
	return key
}

func getSlide(w http.ResponseWriter, r *http.Request) {
	user, topic := getParams(r)
	key := genKey(user, topic)
	fmt.Fprintf(w, "you are get user %s topic %s key %s", user, topic, key)
}

func postSlide(w http.ResponseWriter, r *http.Request) {
	user, topic := getParams(r)
	key := genKey(user, topic)
	fmt.Fprintf(w, "you are get user %s topic %s", user, topic)
}

func putSlide(w http.ResponseWriter, r *http.Request) {
	user, topic := getParams(r)
	key := genKey(user, topic)
	fmt.Fprintf(w, "you are get user %s topic %s", user, topic)
}

func main() {
	mux := routes.New()
	mux.Get("/:user/:topic/", getSlide)
	mux.Post("/:user/:topic/", postSlide)
	mux.Put("/:user/:topic/", putSlide)
	http.Handle("/", mux)
	http.ListenAndServe(":7788", nil)
}
