package main

/**
 * Article 文章
 */
type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Articles []Article
