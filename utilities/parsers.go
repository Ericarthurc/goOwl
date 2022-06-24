package utilities

import (
	"os"
	"strings"
)

const BLOGFOLDER string = "blogPosts"
const BLOGEXTENSION string = ".md"

type Meta struct {
	fileName string
	title    string
	date     string
	tags     []string
	series   string
}

func NewMeta(fileName string) *Meta {
	raw, _ := os.ReadFile("./" + fileName + BLOGEXTENSION)

	rawMeta := strings.Split(strings.TrimSpace(strings.Split(string(raw), "---")[1]), "\n")

	var meta Meta
	meta.fileName = fileName

	for _, v := range rawMeta {
		switch strings.Split(v, ":")[0] {
		case "title":
			meta.title = strings.TrimSpace(strings.Split(v, ":")[1])
		case "date":
			meta.date = strings.TrimSpace(strings.Split(v, ":")[1])
		case "tags":
			meta.tags = strings.Split(strings.TrimSpace(strings.Split(v, ":")[1]), ",")
		case "series":
			meta.series = strings.TrimSpace(strings.Split(v, ":")[1])
		}
	}

	return &meta
}
