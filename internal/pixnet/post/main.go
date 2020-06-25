package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	user   = "kenwhitney"
	postID = "46665129"
)

func buildAPI(user string, postID string) string {
	u := fmt.Sprintf("https://emma.pixnet.cc/blog/articles/%s?user=%s", postID, user)
	log.Print(u)
	return u
}

func main() {
	url := buildAPI(user, postID)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	article, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", article)
}
