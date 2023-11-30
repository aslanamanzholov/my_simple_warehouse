package product

import (
	"github.com/gin-gonic/gin"
	"my_simple_warehouse/pkg/common/models"
	"net/http"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func calculatePercentageIncrease(currentPrice, costPrice float64) float64 {
	if costPrice == 0 {
		return 0 // избегаем деления на ноль
	}

	percentageIncrease := ((currentPrice - costPrice) / costPrice) * 100
	return percentageIncrease
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/products")
	routes.POST("/", h.AddProduct)
	routes.GET("/", h.GetProducts)
	routes.GET("/:id", h.GetProduct)
	routes.PUT("/:id", h.UpdateProduct)
	routes.DELETE("/:id", h.DeleteProduct)
	routes.GET("/view", func(c *gin.Context) {
		var products []models.Product
		db.Find(&products)
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the product_list.html template
			"product_list.html",
			// Pass the data that the page uses
			gin.H{
				"title":   "Almagul Product List",
				"payload": products,
			},
		)

	})
}
