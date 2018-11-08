package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var washingtonPosts = []byte(`
<sitemapindex>
<sitemap>
	 <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
</sitemap>
</sitemapindex>
`)

// <sitemap>
// 	 <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
// </sitemap>
// <sitemap>
// 	 <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
// </sitemap>

// AppStore struct that contains all data for the app
type AppStore struct {
	Title string
	News  map[string]NewsItem
}

// Sitemapindex struct to store site map locations
type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

// News struct to store all News details
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

// NewsItem struct to store details about a news
type NewsItem struct {
	Keyword  string
	Location string
}

func getCurrentDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	return exPath
}

func main() {
	// All the Routes are defiled below
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/news", newsHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	app := AppStore{Title: "Programming in Go"}

	// filePrefix := getCurrentDir()  // Not fetching the project directory
	templateFilePrefix := "/home/sudheesh/Development/go/go_basics/server/templates/"

	t, err := template.ParseFiles(templateFilePrefix + "home.html")
	err = t.Execute(w, app)
	fmt.Println("Error: ", err)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<div>
<h3>Welcome to Go Programming</h3>
</div>`)
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	var s Sitemapindex
	var n News
	// resp, err := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	// fmt.Println("\n\n\n\n\n", resp)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	bytes := washingtonPosts
	// bytes, err := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	newsItems := make(map[string]NewsItem)
	fmt.Println(s)
	for _, Location := range s.Locations {
		resp, err := http.Get(Location)
		if err != nil {
			log.Fatal(err)
		}
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		xml.Unmarshal(bytes, &n)

		for idx := range n.Keywords {
			newsItems[n.Titles[idx]] = NewsItem{n.Keywords[idx], n.Locations[idx]}
		}
	}
	for idx, data := range newsItems {
		fmt.Println("Key: ", idx)
		fmt.Println("Keyword", data.Keyword)
		fmt.Println("Location", data.Location)
	}

	templateFilePrefix := "/home/sudheesh/Development/go/go_basics/server/templates/"

	p := AppStore{Title: "Washington News", News: newsItems}
	t, err := template.ParseFiles(templateFilePrefix + "news.html")
	err = t.Execute(w, p)
	fmt.Println("Error: ", err)

}
