package handlers

import (
	"fmt"
	"github.com/djukela17/pinjur-lunch/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DishHandler struct {
	FullDishList      []models.Dish
	AvailableDishList []models.Dish
}

func (d *DishHandler) AdminCreateForm(c *gin.Context) {

	content := gin.H{
		"title":  "Mate je Caca",
		"dishes": d.FullDishList,
	}
	c.HTML(http.StatusOK, "admin.tmpl.html", content)
}

//func ScrapMealList(c *gin.Context) {
//	client := &http.Client{
//		Timeout: 30 * time.Second,
//	}
//
//	fmt.Println("Getting shit scraped")
//
//	res, err := client.Get("https://www.dobartek.hr/restoran-pizzeria-pinjur/")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer res.Body.Close()
//
//	// go query stuff
//	document, err := goquery.NewDocumentFromReader(res.Body)
//	if err != nil {
//		log.Fatal("Error loading HTTP response body -", err)
//	}
//	dishes := internal.GetAllDishes(document)
//
//	if err != nil {
//		fmt.Println("error on json marshal: ", err)
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{ "status": http.StatusOK, "data": dishes })
//}

func (d *DishHandler) CreateTodayMealList(c *gin.Context) {
	fmt.Println("Create today meal list")

	// check for password

	for i := 0; i < len(d.FullDishList); i++ {
		dishName := c.PostForm("dish_" + strconv.Itoa(i))
		if dishName != "" {
			// find the corresponding dish from the list
			for _, dish := range d.FullDishList {
				if dish.Name == dishName {
					d.AvailableDishList = append(d.AvailableDishList, dish)
				}
			}
		}
	}

	content := gin.H{
		"dishes": d.AvailableDishList,
	}
	c.HTML(http.StatusOK, "admin-created.tmpl.html", content)

}

func GetAvailableMealList(c *gin.Context) {

}
