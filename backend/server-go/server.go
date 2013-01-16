package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

var UrlMap = make(map[string]string)

//Show the HTTP Header ,Just for Test
func showHTTPHeader(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "rootHandler: %s\n", r.URL.Path)
	fmt.Fprintf(w, "URL: %s\n", r.URL)
	fmt.Fprintf(w, "Method:%s\n", r.Method)
	fmt.Fprintf(w, "RquestURI:%s\n", r.RequestURI)
	fmt.Fprintf(w, "Proto:%s\n", r.Proto)
	fmt.Fprintf(w, "HOST:%s\n", r.Host)

}

func showUrlMap() {

	for k, v := range UrlMap {

		fmt.Println(k, v)
	}
}

func getParams(r *http.Request) (string, string, string, string) {

	params := r.URL.Query()
	uri := params.Get("url")
	user := params.Get(":user")
	topic := params.Get(":topic")
	key := user + "@" + topic
	return user, topic, key, uri
}

func getSlide(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	showHTTPHeader(w, r)
	user, topic, key, uri := getParams(r)
	url := UrlMap[key]
	fmt.Fprintf(w, "you are get user %s topic %s key %s uri %s,url %s", user, topic, key, uri, url)
	showUrlMap()
}

func postSlide(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	showHTTPHeader(w, r)
	user, topic, key, uri := getParams(r)
	url := uri
	UrlMap[key] = url
	fmt.Fprintf(w, "you are get user %s topic %s key %s url %s", user, topic, key, url)

}

func putSlide(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	user, topic, key, uri := getParams(r)
	url := uri
	UrlMap[key] = url
	fmt.Fprintf(w, "you are get user %s topic %s key %s url %s", user, topic, key, url)
}

func main() {

	mux := routes.New()
	mux.Get("/:user/:topic", getSlide)
	mux.Post("/:user/:topic", postSlide)
	mux.Put("/:user/:topic", putSlide)
	http.Handle("/", mux)
	http.ListenAndServe(":7777", nil)
}
