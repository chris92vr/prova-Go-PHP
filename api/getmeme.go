package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"html/template"
)

func GetMemeTemplate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/html")
	resp := make(map[string]string)
	getMemeApi := "https://dog.ceo/api/breeds/image/random"
	getMemeResponse, err := http.Get(getMemeApi)
	getMemeBody, err := ioutil.ReadAll(getMemeResponse.Body)
	getMemeJson := string(getMemeBody)
	fmt.Println("GetMeme API Json String:", getMemeJson)
	var getMeme map[string]interface{}
	json.Unmarshal([]byte(getMemeBody), &getMeme)
	resp["message"] = fmt.Sprint(getMeme["message"])
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println("Error happened in template. Err: %s", err)
	} else {
		tmpl.Execute(w, resp)
	}
	return
}

func GetMeme(w http.ResponseWriter, r *http.Request) {
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
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error happened in JSON marshal. Err: %s", err)
	} else {
		w.Write(jsonResp)
	}
	return
}
