package model

import "strings"

type Repository struct {
	inputUrl string
	name     string
	path     string
	mode     Mode
}

func NewUrl(url string) *Repository {
	return &Repository{
		inputUrl: strings.Trim(url, " "),
	}
}

func (t *Repository) Url() string {
	return t.inputUrl
}

func (t *Repository) Valid() bool {
	if strings.HasPrefix(t.inputUrl, "https://") {
		t.mode = HTTPS
		return true
	}

	if strings.HasPrefix(t.inputUrl, "git@") {
		t.inputUrl = strings.Replace(t.inputUrl, "git@github.com:", "https://github.com/", -1)
		t.mode = HTTPS

		// TODO investigate why my ssh key isn't working :)

		return true
	}

	if strings.HasPrefix(t.inputUrl, "git@") {
		t.inputUrl = strings.Replace(t.inputUrl, "git@github.com:", "https://github.com/", -1)

		t.mode = SSH
		return true
	}

	return false
}

func (t *Repository) Name() string {
	if t.name == "" {
		items := strings.Split(t.inputUrl, "/")
		t.name = strings.Replace(items[len(items)-1], ".git", "", 1)
	}

	return t.name
}

func (t *Repository) Path() string {
	if t.path == "" {
		replacer := strings.NewReplacer(".git", "", ":", "/", "https://", "", "git@", "")
		t.path = replacer.Replace(t.inputUrl)
	}

	return t.path
}

func (t *Repository) Mode() Mode {
	return t.mode
}
