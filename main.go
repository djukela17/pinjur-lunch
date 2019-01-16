package main

import (
	"fmt"
	"github.com/djukela17/pinjur-lunch/internal/handlers"
	"github.com/djukela17/pinjur-lunch/internal/models"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	fullDishList, err := models.LoadDishList("data/discounted-prices.json")
	if err != nil {
		log.Fatal("error while loading dishes: ", err)
	}
	//fmt.Println(fullDishList)

	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")
	CreateRoutes(router, fullDishList)
	go router.Run(":8000")

	getListCmd := make(chan interface{})

	select {
	case cmd := <-getListCmd:
		fmt.Println("Get the fucking list", cmd)
	}

}

func CreateRoutes(router *gin.Engine, fullDishList []models.Dish) {

	routeHandler := handlers.DishHandler{FullDishList: fullDishList}

	//router.GET("/api/scrap", handlers.ScrapMealList)
	router.GET("/admin", routeHandler.AdminCF)

	router.POST("/admin/create-meal-form", routeHandler.CreateTodayMealList)
	//router.GET("/admin/create-form", handlers.AdminCreateForm)
}
