package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

func dog(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	responseData := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	fmt.Fprintln(w, "DOG DOG DOG")
	fmt.Fprintf(w, "Method: %s \n", responseData.Method)
	fmt.Fprintf(w, "URL: %s \n", responseData.URL)
	fmt.Fprintf(w, "Submissions: %s \n", responseData.Submissions)
	fmt.Fprintf(w, "Header: %s \n", responseData.Header)
	fmt.Fprintf(w, "Host: %s \n", responseData.Host)
	fmt.Fprintf(w, "ContentLength: %d \n", responseData.ContentLength)
}

func cat(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	responseData := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	fmt.Fprintln(w, "CAT CAT CAT")
	fmt.Fprintf(w, "Method: %s \n", responseData.Method)
	fmt.Fprintf(w, "URL: %s \n", responseData.URL)
	fmt.Fprintf(w, "Submissions: %s \n", responseData.Submissions)
	fmt.Fprintf(w, "Header: %s \n", responseData.Header)
	fmt.Fprintf(w, "Host: %s \n", responseData.Host)
	fmt.Fprintf(w, "ContentLength: %d \n", responseData.ContentLength)
}

func hello(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello %s!", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/dog", dog)
	router.GET("/cat", cat)
	router.GET("/hello/:name", hello)

	http.ListenAndServe(":8080", router)
}
