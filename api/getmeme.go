package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	getDogAPI := "https://dog.ceo/api/breeds/image/random"
	dogResponse, err := http.Get(getDogAPI)
	dogBody, err := ioutil.ReadAll(dogResponse.Body)
	dogJson := string(dogBody)
	fmt.Println("Dog API Json String:", dogJson)
	var dog map[string]interface{}
	json.Unmarshal([]byte(dogJson), &dog)
	resp["dog"] = fmt.Sprint(dog["message"])
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error happened in JSON marshal. Err: %s", err)
	} else {
		w.Write(jsonResp)
	}

}
