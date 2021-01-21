package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var (
	detectTheHTMLContent = regexp.MustCompile(`(<img src=")+((https?)://([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?)" +(class="img-responsive") +(alt="+([\w\d\s\t\v,ñ\.\,@?^=%&/~#])+)`)
	detectTheTitle       = regexp.MustCompile(`alt="+([\w\d\s\t\v,ñ\.\,@?^=%&/~#]+)`)
	detectTheImageURL    = regexp.MustCompile(`(https?://([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?)`)
)

type bodyMeme struct {
	Title    string `json:"title"`
	ImageURL string `json:"imageURL"`
}
type apiMeme struct {
	Memes []bodyMeme `json:"memes"`
}

func getMemes(apiChan chan apiMeme, errChan chan error) {
	api := apiMeme{}
	res, err := http.Get("https://www.memedroid.com/memes/latest")
	if err != nil {
		log.Println(err.Error())
		apiChan <- api
		errChan <- nil
	}
	defer res.Body.Close()
	htmlInfo, _ := ioutil.ReadAll(res.Body)

	for _, i := range detectTheHTMLContent.FindAllString(string(htmlInfo), -1) {
		api.Memes = append(api.Memes, bodyMeme{
			Title:    strings.Replace(detectTheTitle.FindString(i), `alt="`, "", -1),
			ImageURL: detectTheImageURL.FindString(i),
		})
	}
	apiChan <- api
	errChan <- nil

}
func sendMemes(w http.ResponseWriter, r *http.Request) {
	api, errChan := make(chan apiMeme), make(chan error)
	go getMemes(api, errChan)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(<-api)
}
func main() {
	http.HandleFunc("/", sendMemes)
	http.ListenAndServe(":8080", nil)

}
