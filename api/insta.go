package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func insta(w http.ResponseWriter, r *http.Request) {

	url := "https://instagram29.p.rapidapi.com/search/avtistkid"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "instagram29.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "251669fb4bmshce01bc256be2467p164627jsnefa56c023a18")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(w, res)
	fmt.Println(string(body))

}
