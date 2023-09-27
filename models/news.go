package models

import "time"

type News struct {
	ID    int
	Title string
	Slug  string
	Body  string
	Date  time.Time
}

var (
	news   []*News
	nextID = 1
)

func GetNews() []*News {
	return news
}

func AddNewsItem(n News) (News, error) {
	n.ID = nextID
	nextID++

	news = append(news, &n)

	return n, nil
}
