package main

import (
	"flag"
	"fmt"
	"github.com/djukela17/pinjur-lunch/internal/handlers"
	"github.com/djukela17/pinjur-lunch/internal/models"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	port = flag.String("p", "80", "custom port value (default 80)")
	host = flag.String("host", "192.168.190.111", "host address (default: 192.168.190.111")
)

func main() {
	flag.Parse()
	*port = ":" + *port

	//fullDishList, err := models.LoadDishList("data/discounted-prices.json")
	//if err != nil {
	//	log.Fatal("error while loading dishes: ", err)
	//}

	dc, err := models.NewDishCollection("data/discounted-prices.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dc)

	nameList, err := models.LoadUsernameSuggestList("data/ip-username.json")
	if err != nil {
		log.Fatal(err)
	}

	mh := handlers.NewMainHandler(dc, nameList, *host, *port)

	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")

	mhCreateRoutes(router, mh)
	//CreateRoutes(router, fullDishList, nameList, *host, *port)

	if err := router.Run(*port); err != nil {
		log.Fatal(err)
		return
	}
}

func CreateRoutes(router *gin.Engine, fullDishList []models.Dish, nameList map[string]string, host, port string) {

	routeHandler := handlers.MainHandler{FullDishList: fullDishList, NameList: nameList, HostAddress: handlers.CreateHostAddress(host, port)}

	// admin
	router.GET("/admin/create", routeHandler.AdminCreateForm)
	router.POST("/admin/create", routeHandler.CreateTodayMealList)
	router.GET("/admin/list", routeHandler.ListActiveChoices)

	// users
	router.GET("/", routeHandler.UserForm)
	router.POST("/", routeHandler.UpdateActiveDishList)
}

func mhCreateRoutes(router *gin.Engine, routeHandler handlers.MainHandler) {

	// admin
	router.GET("/admin/create", routeHandler.AdminCreateForm)
	router.POST("/admin/create", routeHandler.CreateTodayMealList)
	router.GET("/admin/list", routeHandler.ListActiveChoices)

	// users
	router.GET("/", routeHandler.UserForm)
	router.POST("/", routeHandler.UpdateActiveDishList)
}
