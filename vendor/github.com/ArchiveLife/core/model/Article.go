package model

import "time"

// Article
type Article struct {
	ID            ID
	Title         *string
	Author        *Author
	CoAuthors     []*Author
	PublishDate   *time.Time
	Content       *string
	Tags          []string
	Type          string
	Medias        []*Media
	References    []*Reference
	ExtAttributes map[string]interface{}
}
