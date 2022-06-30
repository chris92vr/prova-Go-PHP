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
	getDogApi := "https://dog.ceo/api/breeds/image/random"
	dogApiResponse, _ := http.Get(getDogApi)
	dogApiBody, _ := ioutil.ReadAll(dogApiResponse.Body)
	dogApiJson := string(dogApiBody)
	var dogApi map[string]interface{}
	json.Unmarshal([]byte(dogApiJson), &dogApi)
	resp["dog"] = fmt.Sprint(dogApi["message"])
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error happened in JSON marshal. Err: %s", err)
	} else {
		w.Write(jsonResp)
	}
	fmt.Fprintf(w, resp["dog"], "<img src='"+resp["dog"]+"'>")

}
