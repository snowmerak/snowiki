package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var siteMap = make([]string, 0)

func makeSiteMap(domain string) {
	file, err := os.Create(filepath.Join(".", "public", "sitemap.txt"))
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range siteMap {
		if _, err := fmt.Fprintf(file, "%s/%s\n", domain, name); err != nil {
			log.Fatal(err)
		}
	}
	file.Close()
}
