package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	input := make(chan string)
	output := make(chan parsed)
	go startParser(input, output)
	dirs, err := os.ReadDir(filepath.Join(".", "src"))
	if err != nil {
		log.Fatal(err)
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			continue
		}
		if strings.HasSuffix(dir.Name(), ".md") {
			input <- filepath.Join(".", "src", dir.Name())
		}
	}
	close(input)
	publicPath := filepath.Join(".", "public")
	if err := os.MkdirAll(publicPath, 0755); err != nil {
		log.Fatal(err)
	}
	for p := range output {
		if p.err != nil {
			log.Fatal(p.err)
		}
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
