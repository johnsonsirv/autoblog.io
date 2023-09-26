package main

import (
	"fmt"
	"time"

	"github.com/johnsosirv/auto-blog/models"
)

func main() {
	news := models.News{
		ID:    1,
		Title: "Dollar to Naira",
		Slug:  "dollar-to-naira",
		Body:  "Naira crashes to N1000 to $1 dollar",
		Date:  time.Now(),
	}

	fmt.Println(news)
}
