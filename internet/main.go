package main

import (
	"encoding/xml"
	"fmt"
)

var washingtonPosts = []byte(`
<sitemapindex>
<sitemap>
	 <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
</sitemap>
<sitemap>
	 <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
</sitemap>
<sitemap>
	 <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
</sitemap>
</sitemapindex>
`)

// SiteMapIndex Struct
// to store all locations from the post
type SiteMapIndex struct {
	Locations []Location `xml:"sitemap"`
}

// Location Struct
// to store a specific location
type Location struct {
	Loc string `xml:"loc"`
}

// Method of Struct location
// used to convert a struct type of Loc to a string
func (e Location) String() string {
	return fmt.Sprintf(e.Loc)
}

func main() {
	// resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	// bytes, _ := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()

	bytes := washingtonPosts

	stringBody := string(bytes)
	fmt.Println("Strings from Washington Post: ", stringBody)
	// fmt.Println("Bytes from Washington Post: ", bytes)

	var s SiteMapIndex
	fmt.Println("SiteMapIndex s: ", s)
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
	fmt.Println("SiteMapIndex s: ", s)
}
