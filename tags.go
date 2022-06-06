package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var tagMap = make(map[string][]string)

func makeTagsPage() {
	buttons := make([]string, 0, len(tagMap))
	tags := make([]string, 0, len(tagMap))
	for tag := range tagMap {
		tags = append(tags, tag)
	}
	sort.Strings(tags)
	for _, tag := range tags {
		buttons = append(buttons, fmt.Sprintf(`<button onclick="location.href='./tags/%s.html'">%s</button>`, tag, tag))
	}
	file, err := os.Create(filepath.Join(".", "public", "tags.html"))
	if err != nil {
		log.Fatal(err)
	}
	sb := strings.Builder{}
	sb.WriteString(`<h1>Tags</h1><br/><div style="display: flex; align-items: center; align-self: center; flex-direction: row; flex-wrap: wrap;">`)
	for _, button := range buttons {
		sb.WriteString(button)
	}
	sb.WriteString(`</div>`)
	if _, err := fmt.Fprintf(file, template, sb.String()); err != nil {
		log.Fatal(err)
	}
}

func makeEachTagPage() {
	replacedTemplate := strings.Replace(template, "./water.css", "../water.css", -1)
	replacedTemplate = strings.Replace(replacedTemplate, "./tags.html", "../tags.html", -1)
	for tag, names := range tagMap {
		file, err := os.Create(filepath.Join(".", "public", fmt.Sprintf("tags/%s.html", tag)))
		if err != nil {
			log.Fatal(err)
		}
		sb := strings.Builder{}
		sb.WriteString(`<h1>`)
		sb.WriteString(tag)
		sb.WriteString(`</h1><br/><div style="display: flex; align-items: center; align-self: center; flex-direction: row; flex-wrap: wrap;">`)
		for _, name := range names {
			sb.WriteString(fmt.Sprintf(`<button onclick="location.href='../%s'">%s</button>`, name, strings.TrimSuffix(name, ".html")))
		}
		sb.WriteString(`</div>`)
		if _, err := fmt.Fprintf(file, replacedTemplate, sb.String()); err != nil {
			log.Fatal(err)
		}
	}
}
