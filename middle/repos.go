package middle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"git.nova-vps.ml/lexisother/GitFren/middle/api"
	"github.com/anaskhan96/soup"
)

func GetUserNames() []string {
	resp, _ := soup.Get("https://git.nova-vps.ml/explore/users")
	final := []string{}

	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "class", "user").FindAll("div", "class", "item")
	for _, link := range links {
		names := link.FindAll("span", "class", "header")
		for _, name := range names {
			text := name.Find("a").Text()
			m1 := regexp.MustCompile(`\s`)
			text = m1.ReplaceAllString(text, "")
			final = append(final, text)
		}
	}

	return final
}

func GetUserRepos(url string, user string) []api.Repository {
	var repos []api.Repository

	resp, err := http.Get(fmt.Sprintf("%s/users/%s/repos", url, user))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	json.Unmarshal([]byte(body), &repos)

	return repos
}

func GetRepoData(url string, owner string, repoName string) api.Repository {
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s", url, owner, repoName))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var repo api.Repository
	json.Unmarshal([]byte(body), &repo)

	return repo
}
