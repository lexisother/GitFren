package middle

import "time"

type Repository struct {
	ID                        int             `json:"id"`
	Owner                     Owner           `json:"owner"`
	Name                      string          `json:"name"`
	FullName                  string          `json:"full_name"`
	Description               string          `json:"description"`
	Empty                     bool            `json:"empty"`
	Private                   bool            `json:"private"`
	Fork                      bool            `json:"fork"`
	Template                  bool            `json:"template"`
	Parent                    interface{}     `json:"parent"`
	Mirror                    bool            `json:"mirror"`
	Size                      int             `json:"size"`
	HTMLURL                   string          `json:"html_url"`
	SSHURL                    string          `json:"ssh_url"`
	CloneURL                  string          `json:"clone_url"`
	OriginalURL               string          `json:"original_url"`
	Website                   string          `json:"website"`
	StarsCount                int             `json:"stars_count"`
	ForksCount                int             `json:"forks_count"`
	WatchersCount             int             `json:"watchers_count"`
	OpenIssuesCount           int             `json:"open_issues_count"`
	OpenPrCounter             int             `json:"open_pr_counter"`
	ReleaseCounter            int             `json:"release_counter"`
	DefaultBranch             string          `json:"default_branch"`
	Archived                  bool            `json:"archived"`
	CreatedAt                 time.Time       `json:"created_at"`
	UpdatedAt                 time.Time       `json:"updated_at"`
	Permissions               Permissions     `json:"permissions"`
	HasIssues                 bool            `json:"has_issues"`
	InternalTracker           InternalTracker `json:"internal_tracker"`
	HasWiki                   bool            `json:"has_wiki"`
	HasPullRequests           bool            `json:"has_pull_requests"`
	HasProjects               bool            `json:"has_projects"`
	IgnoreWhitespaceConflicts bool            `json:"ignore_whitespace_conflicts"`
	AllowMergeCommits         bool            `json:"allow_merge_commits"`
	AllowRebase               bool            `json:"allow_rebase"`
	AllowRebaseExplicit       bool            `json:"allow_rebase_explicit"`
	AllowSquashMerge          bool            `json:"allow_squash_merge"`
	DefaultMergeStyle         string          `json:"default_merge_style"`
	AvatarURL                 string          `json:"avatar_url"`
	Internal                  bool            `json:"internal"`
	MirrorInterval            string          `json:"mirror_interval"`
}
type Owner struct {
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
type Permissions struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}
type InternalTracker struct {
	EnableTimeTracker                bool `json:"enable_time_tracker"`
	AllowOnlyContributorsToTrackTime bool `json:"allow_only_contributors_to_track_time"`
	EnableIssueDependencies          bool `json:"enable_issue_dependencies"`
}
