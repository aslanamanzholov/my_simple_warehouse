package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"my_simple_warehouse/pkg/common/models"
)

func (h handler) GetProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &product)
}
