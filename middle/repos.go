package middle

import (
	"regexp"

	"github.com/anaskhan96/soup"
)

func GetRepoNames() []string {
	resp, _ := soup.Get("https://git.nova-vps.ml/explore/repos")
	final := []string{}

	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "class", "repository").FindAll("div", "class", "item")
	for _, link := range links {
		names := link.FindAll("a", "class", "name")
		for _, name := range names {
			text := name.Text()
			m1 := regexp.MustCompile(`\s`)
			text = m1.ReplaceAllString(text, "")
			final = append(final, text)
		}
	}

	return final
}