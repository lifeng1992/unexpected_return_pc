package main

import (
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/sqlite"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())

	if err := r.Run(":8086"); err != nil {
		panic(err)
	}
}
