package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
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
	if err := os.MkdirAll(publicPath, 0755); err != nil {
		log.Fatal(err)
	}
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
		if _, err := fmt.Fprintf(file, template, p.data); err != nil {
			log.Fatal(err)
		}
		file.Close()
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
		makeTagsPage()
	}
	{
		fmt.Println("create each tag pages")
		makeEachTagPage()
	}
	fmt.Println("Done")
}
