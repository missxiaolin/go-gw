package formatter

import "go-gw/web/model"

func NewBase(article model.Article) map[string]interface{} {
	return map[string]interface{}{
		"id": article.ID,
		"cid": article.Cid,
		"title": article.Title,
		"author": article.Author,
		"content": article.Content,
		"keywords": article.Keywords,
		"description": article.Description,
		"status": article.Status,
		"created_at": article.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at": article.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
