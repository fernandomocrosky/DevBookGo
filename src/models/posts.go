package models

import "time"

type Posts struct {
	ID        uint64    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	UserId    uint64    `json:"user_id,omitempty"`
	UserNick  string    `json:"user_nick,omitempty"`
	Likes     uint64    `json:"likes"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
