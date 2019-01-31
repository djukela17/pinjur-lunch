package main

import (
	"flag"
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

	dishes, err := models.NewDishCollection("data/discounted-prices.json")
	if err != nil {
		log.Fatal(err)
	}

	nameList, err := models.LoadUsernameSuggestList("data/ip-username.json")
	if err != nil {
		log.Fatal(err)
	}

	mh := handlers.NewMainHandler(dishes, nameList, *host, *port)

	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")

	CreateRoutes(router, mh)

	if err := router.Run(*port); err != nil {
		log.Fatal(err)
		return
	}
}

func CreateRoutes(router *gin.Engine, routeHandler handlers.MainHandler) {

	// admin
	router.GET("/admin/create", routeHandler.AdminCreateForm)
	router.POST("/admin/create", routeHandler.CreateTodayMealList)
	router.GET("/admin/list", routeHandler.ListActiveChoices)

	// users
	router.GET("/", routeHandler.UserForm)
	router.POST("/", routeHandler.UpdateActiveDishList)

	router.StaticFile("/deer", "./web/images/deer.png")

}
