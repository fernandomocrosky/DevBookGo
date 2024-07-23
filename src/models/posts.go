package models

import (
	"errors"
	"strings"
	"time"
)

type Posts struct {
	ID        uint64    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	UserId    uint64    `json:"user_id,omitempty"`
	UserNick  string    `json:"user_nick,omitempty"`
	Likes     uint64    `json:"likes"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (post *Posts) Prepare() error {
	if err := post.validate(); err != nil {
		return err
	}
	post.format()

	return nil
}

func (post *Posts) validate() error {
	if post.Title == "" {
		return errors.New("title is required and cannot be empty")
	}

	if post.Content == "" {
		return errors.New("content is required and cannot be empty")
	}

	return nil
}

func (post *Posts) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
