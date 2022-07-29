package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"products/model"
)


func FindProduct(c echo.Context) error {
	//开始接收接口参数
	input := new(model.ProductInput)
	err := c.Bind(input)
	if err != nil {
		return err
	}

	product := new(model.Product)
	err = model.DB.Raw(`select id,name,description,price from products where id = ?`,
		input.Id,
	).Scan(&product).Error
	if err != nil {
		fmt.Println("SQL 执行错误", err)
		return err
	}
	return c.JSON(http.StatusOK, product)
}

func FindAll(c echo.Context) error {
	//var rows *sql.Rows
	var products []model.Product
	rows, err := model.DB.Raw(`select * from products`).Rows()
	for rows.Next() {
		var product model.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price)
		//TODO if errs
		if err != nil {
			return err
		}
		products = append(products, product)
	}
	if err != nil {
		fmt.Println("SQL 执行错误", err)
		return err
	}
	return c.JSON(http.StatusOK, products)
}

func GetHighPrice(c echo.Context) error {
	//var rows *sql.Rows
	var products []model.Product
	rows, err := model.DB.Raw(`select * from products group by id order by price desc`).Rows()
	for rows.Next() {
		var product model.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price)
		//TODO if errs
		if err != nil {
			return err
		}
		products = append(products, product)
	}
	if err != nil {
		fmt.Println("SQL 执行错误", err)
		return err
	}
	return c.JSON(http.StatusOK, products)
}
