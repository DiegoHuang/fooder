package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	query = "all day roasting company"
)

func buildAPI(search string) string {
	u, _ := url.Parse("https://emma.pixnet.cc/blog/articles/search")
	q := u.Query()
	q.Add("format", "json")
	q.Add("key[]", search)
	u.RawQuery = q.Encode()
	log.Print(u.String())
	return u.String()
}

func main() {
	url := buildAPI(query)
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
