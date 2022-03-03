package handler

import (
	"net/http"
	"time"

	"github.com/GenkiHirano/go-debug/repository"
	"github.com/gin-gonic/gin"
)

type requestProduct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type responseProduct struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// CreateUserを実装してください。
// 引数、戻り値、関数の中身はお好きに変更可能です。
func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestProduct
		if err := c.Bind(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		}

		createdProduct, err := repository.Create(req.ID, req.Name)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		}

		res := responseProduct{
			ID:        createdProduct.ID,
			Name:      createdProduct.Name,
			CreatedAt: createdProduct.CreatedAt,
		}

		c.JSON(http.StatusCreated, res)
	}
}
