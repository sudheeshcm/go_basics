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
	"sync"
)

// <sitemap>
// 	<loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
// </sitemap>
var washingtonPosts = []byte(`
<sitemapindex>
	<sitemap>
		<loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
	</sitemap>
	</sitemapindex>
	`)

// <sitemap>
// 	<loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
// </sitemap>

var wg sync.WaitGroup

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
	templateFilePrefix := "/home/sudheesh/Development/go/go_basics/project/templates/"

	t, err := template.ParseFiles(templateFilePrefix + "home.html")
	err = t.Execute(w, app)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<div>
<h3>Welcome to Go Programming</h3>
</div>`)
}

func newsRoutine(c chan News, Location string) {
	defer wg.Done()
	var n News
	resp, err := http.Get(Location)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()
	c <- n
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	var s Sitemapindex
	queue := make(chan News, 30)
	// resp, err := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	// fmt.Println("\n\n\n\n\n", resp)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// bytes, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()

	bytes := washingtonPosts
	xml.Unmarshal(bytes, &s)
	newsItems := make(map[string]NewsItem)
	fmt.Println(s)

	for _, Location := range s.Locations {
		wg.Add(1)
		go newsRoutine(queue, Location)
	}

	wg.Wait()
	close(queue)

	for elem := range queue {
		for idx := range elem.Keywords {
			newsItems[elem.Titles[idx]] = NewsItem{elem.Keywords[idx], elem.Locations[idx]}
		}
	}

	// To print the NewsItems in the console
	// for idx, data := range newsItems {
	// 	fmt.Println("Key: ", idx)
	// 	fmt.Println("Keyword", data.Keyword)
	// 	fmt.Println("Location", data.Location)
	// }

	templateFilePrefix := "/home/sudheesh/Development/go/go_basics/project/templates/"
	p := AppStore{Title: "Washington News", News: newsItems}
	t, err := template.ParseFiles(templateFilePrefix + "news.html")
	err = t.Execute(w, p)
	fmt.Println("Error: ", err)

}
