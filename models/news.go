package models

import "time"

type News struct {
	ID    int
	Title string
	Slug  string
	Body  string
	Date  time.Time
}

var news []*News
