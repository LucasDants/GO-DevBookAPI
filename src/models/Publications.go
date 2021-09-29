package models

import (
	"errors"
	"strings"
	"time"
)

type Publications struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorID,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (publication *Publications) Prepare() error {
	if err := publication.validate(); err != nil {
		return err
	}

	publication.format()
	return nil
}

func (publication *Publications) validate() error {
	if publication.Title == "" {
		return errors.New("título é obrigatório e não pode estar em branco")
	}
	if publication.Content == "" {
		return errors.New("conteúdo é obrigatório e não pode estar em branco")
	}
	return nil
}

func (publication *Publications) format() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}
