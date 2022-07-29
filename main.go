package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"products/controller"
	"products/model"
)


func main() {
	var err error
	dsn := "postgresql://postgres:ShenNiu.001@192.168.1.118:5432/qh_test?connect_timeout=5&sslmode=disable&TimeZone=Asia/Shanghai"
	model.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(fmt.Sprintf("数据库连接错误: %s", err))
	}

	e := echo.New() //开启web服务
	e.GET("/", controller.Home)
	e.POST("/addProduct", controller.AddProduct)
	e.GET("/findProduct", controller.FindProduct)
	e.DELETE("/deleteProduct", controller.DeleteProduct)
	e.GET("/findAll", controller.FindAll)
	e.GET("/getHighPrice", controller.GetHighPrice)
	e.PUT("/changePrice", controller.ChangePrice)
	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}




