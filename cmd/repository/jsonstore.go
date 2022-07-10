package repository

import (
	"net/url"

	"github.com/sreeks87/tolink/cmd/domain"
)

type jsonRepo struct {
	Path string
}

func NewJSONFileRepository(path string) domain.Repository {
	return &jsonRepo{
		Path: path,
	}
}

func (j *jsonRepo) Get(key string) (*domain.Bookmark, error) {
	return nil, nil
}

func (j *jsonRepo) Add(string, url.URL, string) error    { return nil }
func (j *jsonRepo) Update(string, url.URL, string) error { return nil }
func (j *jsonRepo) Delete(string) error                  { return nil }
