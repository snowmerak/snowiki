package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"

	embed "github.com/13rac1/goldmark-embed"
	hashtag "github.com/abhinav/goldmark-hashtag"
	mermaid "github.com/abhinav/goldmark-mermaid"
	toc "github.com/abhinav/goldmark-toc"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

type parsed struct {
	name     string
	data     []byte
	hashtags []string
	err      error
}

func startParser(input <-chan string, output chan<- parsed) {
	for name := range input {
		data, err := os.ReadFile(name)
		if err != nil {
			output <- parsed{err: err}
			continue
		}
		buf := bytes.NewBuffer(nil)
		mdParser := goldmark.New(
			goldmark.WithParserOptions(parser.WithAutoHeadingID()),
			goldmark.WithExtensions(
				extension.GFM,
				&hashtag.Extender{},
				&mermaid.Extender{},
				embed.New(),
				&toc.Extender{},
				highlighting.NewHighlighting(
					highlighting.WithStyle("monokai"),
					highlighting.WithFormatOptions(
						html.WithLineNumbers(true),
					),
				),
			),
		)
		if err := mdParser.Convert(data, buf); err != nil {
			output <- parsed{err: err}
			continue
		}
		file, err := os.Open(name)
		if err != nil {
			output <- parsed{err: err}
			continue
		}
		scanner := bufio.NewReader(file)
		var hashtags []string
		for {
			line, err := scanner.ReadString('\n')
			if err != nil && err != io.EOF {
				output <- parsed{err: err}
			}
			contents := strings.Split(line, " ")
			for _, content := range contents {
				if strings.HasPrefix(content, "#") {
					content = strings.TrimLeft(content, "#")
					if content != "" {
						hashtags = append(hashtags, content)
					}
				}
			}
			if err == io.EOF {
				break
			}
		}
		output <- parsed{data: buf.Bytes(), hashtags: hashtags, name: name}
	}
	close(output)
}
