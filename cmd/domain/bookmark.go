package domain

import "net/url"

type Bookmark struct {
	Key         string  `json:"key"`
	URL         url.URL `json:"url"`
	Description string  `json:"description"`
}

type Service interface {
	GetURL(string) (url.URL, error)
	GetAll(string) (*Bookmark, error)
	Add(string, url.URL, string) error
	Update(string, url.URL, string) error
	Delete(string) error
}
