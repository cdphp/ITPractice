package model

// Article 文章信息
type Article struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Articles array
type Articles []Article
