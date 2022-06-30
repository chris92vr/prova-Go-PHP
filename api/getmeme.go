package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"html/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	getMemeApi := "https://dog.ceo/api/breeds/image/random"
	getMemeResponse, err := http.Get(getMemeApi)
	getMemeBody, err := ioutil.ReadAll(getMemeResponse.Body)
	getMemeJson := string(getMemeBody)
	fmt.Println("GetMeme API Json String:", getMemeJson)
	var getMeme map[string]interface{}
	json.Unmarshal([]byte(getMemeBody), &getMeme)
	resp["message"] = fmt.Sprint(getMeme["message"])
	templ, err := template.ParseFiles("layout.html")
	data := struct {
		Message string
	}{
		Message: fmt.Sprint(getMeme["message"]),
	}
	if err != nil {
		panic(err)
	} else {
		templ.Execute(w, data)
	}

}

func cat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	getMemeApi := "https://api.thecatapi.com/v1/images/search"
	getMemeResponse, err := http.Get(getMemeApi)
	getMemeBody, err := ioutil.ReadAll(getMemeResponse.Body)
	getMemeJson := string(getMemeBody)
	fmt.Println("GetMeme API Json String:", getMemeJson)
	var getMeme map[string]interface{}
	json.Unmarshal([]byte(getMemeBody), &getMeme)
	resp["message"] = fmt.Sprint(getMeme["message"])
	templ, err := template.ParseFiles("layout1.html")
	data := struct {
		Message string
	}{
		Message: fmt.Sprint(getMeme["message"]),
	}
	if err != nil {
		panic(err)
	} else {
		templ.Execute(w, data)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		index(w, r)
	} else if r.URL.Path == "/dog" {
		dog(w, r)
	} else if r.URL.Path == "/cat" {
		cat(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
