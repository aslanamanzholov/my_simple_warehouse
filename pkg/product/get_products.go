package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"my_simple_warehouse/pkg/common/models"
)

func (h handler) GetProducts(c *gin.Context) {
	var products []models.Product

	if result := h.DB.Find(&products); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &products)
}
