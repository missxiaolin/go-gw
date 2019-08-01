package formatter

import "go-gw/web/model"

func CategoryBase(category model.Category) map[string]interface{} {
	return map[string]interface{}{
		"id": category.ID,
		"name": category.Name,
		"pid": category.Pid,
		"status": category.Status,
		"created_at": category.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at": category.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
