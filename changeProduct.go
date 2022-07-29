package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"products/model"
	prostruct "products/model"
)

func AddProduct(c echo.Context) error {
	//开始接收接口参数
	product := new(prostruct.Product)
	err := c.Bind(product)
	if err != nil {
		return err
	}

	fmt.Printf("Here are the products %v", product)

	// 将数据插入数据库
	err = model.DB.Exec(`insert into products(id,name,description,price) values(?,?,?,?)`,
		product.Id,
		product.Name,
		product.Description,
		product.Price,
	).Error
	if err != nil {
		fmt.Println("SQL 执行错误", err)
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"input": product,
		"msg":   "数据插入成功",
	})
}


func DeleteProduct(c echo.Context) error {
	result := model.DB.Exec(`delete from products WHERE id = ?`,
		c.QueryParam("id"),
	)
	err := result.Error
	if err != nil {
		fmt.Println("SQL 执行错误", err)
		return err
	}
	//判断删除是否成功
	if result.RowsAffected == 1 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "数据删除成功",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "数据删除失败",
	})
}


func ChangePrice(c echo.Context) error {
	product := new(prostruct.Product)
	err := c.Bind(product)
	if err != nil {
		return err
	}
	err = model.DB.Raw(`update products set price = ? WHERE id = ?`,
		product.Price,
		product.Id,
	).Scan(&product).Error
	if err != nil {
		fmt.Println("SQL 执行错误", err)
		return err
	}
	return c.JSON(http.StatusOK, product)
}