package models

import (
	"time"

	"github.com/google/uuid"
)

type News struct {
	ID    string
	Title string
	Slug  string
	Body  string
	Date  time.Time
}

var (
	news   []*News
)

func GetNews() []*News {
	return news
}

func AddNewsItem(n News) (News, error) {
	n.ID = uuid.NewString()

	news = append(news, &n)

	return n, nil
}

func GetNewsById(id string) (News, error)  {
	return News{}, nil
}

func RemoveNewsById(id string) (error)  {
	return nil
}
