package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
)

//go:embed THIRD_PARTY_LICENSES.md
var file_ThirdPartyLicenses []byte

func main() {
	wikiName := os.Getenv("WIKI_NAME")
	siteDomain := os.Getenv("SITE_DOMAIN")
	lang := os.Getenv("LANG")

	dirs, err := os.ReadDir(filepath.Join(".", "src"))
	if err != nil {
		log.Fatal(err)
	}
	parseds := make([]parsed, 0, len(dirs))
	for _, dir := range dirs {
		if dir.IsDir() {
			continue
		}
		if strings.HasSuffix(dir.Name(), ".md") {
			parsed, err := parse(filepath.Join(".", "src", dir.Name()))
			if err != nil {
				log.Fatal(err)
			}
			parseds = append(parseds, parsed)
		}
	}
	publicPath := filepath.Join(".", "public")
	if err := os.RemoveAll(publicPath); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(publicPath, 0755); err != nil {
		log.Fatal(err)
	}

	// Initialize minifier
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)

	var buf bytes.Buffer

	for _, p := range parseds {
		name := strings.TrimSuffix(filepath.Base(p.name), ".md") + ".html"
		fmt.Printf("Create %s\n", name)
		for _, tag := range p.hashtags {
			tagMap[tag] = append(tagMap[tag], name)
		}
		file, err := os.Create(filepath.Join(publicPath, name))
		if err != nil {
			log.Fatal(err)
		}
		// Minify HTML
		buf.Reset()
		title := strings.TrimSuffix(name, ".html")
		if title == "index" && wikiName != "" {
			title = wikiName
		}
		if _, err := fmt.Fprintf(&buf, template, lang, title, p.data); err != nil {
			log.Fatal(err)
		}
		err = m.Minify("text/html", file, &buf)
		if err != nil {
			log.Fatal(err)
		}
		file.WriteString("\n<!-- Third Party Licenses : /THIRD_PARTY_LICENSES.md -->")
		file.Close()
		siteMap = append(siteMap, name)
	}
	{
		fmt.Println("create water.css")
		file, err := os.Create(filepath.Join(publicPath, "water.css"))
		if err != nil {
			log.Fatal(err)
		}
		if _, err := file.Write([]byte(watercss)); err != nil {
			log.Fatal(err)
		}
		file.Close()
	}
	{
		fmt.Println("create tags folders")
		if err := os.MkdirAll(filepath.Join(".", "public", "tags"), 0755); err != nil {
			log.Fatal(err)
		}
		fmt.Println("create tags.html")
		makeTagsPage(lang)
	}
	{
		fmt.Println("create each tag pages")
		makeEachTagPage(lang)
	}
	{
		// Create THIRD_PARTY_LICENSES.md
		file, err := os.Create(filepath.Join(publicPath, "THIRD_PARTY_LICENSES.md"))
		if err != nil {
			log.Fatal(err)
		}
		if _, err := file.Write(file_ThirdPartyLicenses); err != nil {
			log.Fatal(err)
		}
		file.Close()
	}
	{
		fmt.Println("create robots.txt file")
		file, err := os.Create(filepath.Join(publicPath, "robots.txt"))
		if err != nil {
			log.Fatal(err)
		}
		if _, err := file.Write([]byte(robotstxt)); err != nil {
			log.Fatal(err)
		}
		file.Close()
	}
	{
		fmt.Println("create sitemap.txt file")
		makeSiteMap(siteDomain)
	}
	fmt.Println("Done")
}
