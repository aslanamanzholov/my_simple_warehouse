package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"my_simple_warehouse/pkg/common/models"
)

type UpdateProductRequestBody struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Amount       int    `json:"amount"`
	CurrentPrice int    `json:"current_price"`
	CostPrice    int    `json:"cost_price"`
}

func (h handler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	body := UpdateProductRequestBody{}

	// получаем тело запроса
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	product.Name = body.Name
	product.Description = body.Description
	product.Amount = body.Amount
	product.CurrentPrice = body.CurrentPrice
	product.CostPrice = body.CostPrice

	h.DB.Save(&product)

	c.JSON(http.StatusOK, &product)
}
