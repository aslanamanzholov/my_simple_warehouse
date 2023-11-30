package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"my_simple_warehouse/pkg/common/models"
)

type AddProductRequestBody struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Amount       int    `json:"amount"`
	CurrentPrice int    `json:"current_price"`
	CostPrice    int    `json:"cost_price"`
}

func (h handler) AddProduct(c *gin.Context) {
	body := AddProductRequestBody{}

	// получаем тело запроса
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var product models.Product

	product.Name = body.Name
	product.Description = body.Description
	product.CurrentPrice = body.CurrentPrice
	product.CostPrice = body.CostPrice
	product.Amount = body.Amount

	if result := h.DB.Create(&product); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &product)
}
