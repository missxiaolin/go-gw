package admin

import (
	"github.com/gin-gonic/gin"
	"go-gw/web/controller"
	"go-gw/web/form"
)

type Category struct{ controller.Base }

func (t *Category) CategoryAdd(c *gin.Context)  {
	category := &form.CategoryForm{
		Name: "ceshi",
		PantId: 1,
		Status: 1,
	}
	t.Validator(c, category)
}