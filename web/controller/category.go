package controller

import (
	"github.com/gin-gonic/gin"
	"go-gw/web/form"
)

type Category struct{ Base }

func (t *Category) CategoryAdd(c *gin.Context)  {
	category := &form.CategoryForm{
		Name: "ceshi",
		PantId: 1,
		Status: 1,
	}
	t.Validator(c, category)
}