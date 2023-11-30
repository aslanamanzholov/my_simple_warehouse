package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"my_simple_warehouse/pkg/common/db"
	"my_simple_warehouse/pkg/product"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	product.RegisterRoutes(r, h)
	// здесь регистрируются дополнительные маршруты

	r.LoadHTMLGlob("templates/*")
	r.Run(port)
}
