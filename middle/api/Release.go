package api

import "time"

type Release struct {
	ID              int       `json:"id"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	Name            string    `json:"name"`
	Body            string    `json:"body"`
	URL             string    `json:"url"`
	HTMLURL         string    `json:"html_url"`
	TarballURL      string    `json:"tarball_url"`
	ZipballURL      string    `json:"zipball_url"`
	Draft           bool      `json:"draft"`
	Prerelease      bool      `json:"prerelease"`
	CreatedAt       time.Time `json:"created_at"`
	PublishedAt     time.Time `json:"published_at"`
	Author          Author    `json:"author"`
	Assets          []Assets  `json:"assets"`
}
type Author struct {
	ID                int       `json:"id"`
	Login             string    `json:"login"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	AvatarURL         string    `json:"avatar_url"`
	Language          string    `json:"language"`
	IsAdmin           bool      `json:"is_admin"`
	LastLogin         time.Time `json:"last_login"`
	Created           time.Time `json:"created"`
	Restricted        bool      `json:"restricted"`
	Active            bool      `json:"active"`
	ProhibitLogin     bool      `json:"prohibit_login"`
	Location          string    `json:"location"`
	Website           string    `json:"website"`
	Description       string    `json:"description"`
	Visibility        string    `json:"visibility"`
	FollowersCount    int       `json:"followers_count"`
	FollowingCount    int       `json:"following_count"`
	StarredReposCount int       `json:"starred_repos_count"`
	Username          string    `json:"username"`
}
type Assets struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Size               int       `json:"size"`
	DownloadCount      int       `json:"download_count"`
	CreatedAt          time.Time `json:"created_at"`
	UUID               string    `json:"uuid"`
	BrowserDownloadURL string    `json:"browser_download_url"`
}
