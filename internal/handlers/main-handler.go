package handlers

import (
	"fmt"
	"github.com/djukela17/pinjur-lunch/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MainHandler struct {
	FullDishList      []models.Dish
	AvailableDishList []models.Dish
	DishChoices       models.UserChoices
	NameList          map[string]string
	HostAddress       string
}

func (h *MainHandler) AdminCreateForm(c *gin.Context) {
	content := gin.H{
		"hostAddress": h.HostAddress,
		"dishes":      h.FullDishList,
	}
	c.HTML(http.StatusOK, "admin.tmpl.html", content)
}

func (h *MainHandler) CreateTodayMealList(c *gin.Context) {
	fmt.Println("Create today meal list")

	// check for password

	for i := 0; i < len(h.FullDishList); i++ {
		dishName := c.PostForm("dish_" + strconv.Itoa(i))
		if dishName != "" {
			for _, dish := range h.FullDishList {
				if dish.Name == dishName {
					h.AvailableDishList = append(h.AvailableDishList, dish)
				}
			}
		}
	}

	content := gin.H{
		"dishes": h.AvailableDishList,
	}
	c.HTML(http.StatusOK, "admin-created.tmpl.html", content)
}

func (h *MainHandler) UserForm(c *gin.Context) {
	fmt.Println(c.ClientIP())

	if len(h.AvailableDishList) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
		return
	}
	data := gin.H{
		"hostAddress": h.HostAddress,
		"dishes":      h.AvailableDishList,
		"name":        h.NameList[c.ClientIP()],
	}
	c.HTML(http.StatusOK, "user.tmpl.html", data)
}

func (h *MainHandler) UpdateActiveDishList(c *gin.Context) {
	fmt.Println("Updating dish choices")

	chosenDish := c.PostForm("dish")
	username := c.PostForm("username")
	optionalNote := c.PostForm("optional-note")

	fmt.Println("Client IP:", c.ClientIP())
	fmt.Println("Selected dish:", chosenDish)
	fmt.Println("username:", username)
	fmt.Println("extra note:", optionalNote)

	if err := h.DishChoices.AddDish(h.AvailableDishList, chosenDish, username, optionalNote); err != nil {
		fmt.Println(err)
		return
	}
	data := gin.H{
		"name":         h.NameList[c.ClientIP()],
		"chosenDish":   chosenDish,
		"optionalNote": optionalNote,
		"orderStatus":  "success",
	}
	c.HTML(http.StatusOK, "user-submitted.tmpl.html", data)
}

func (h *MainHandler) ListActiveChoices(c *gin.Context) {
	fmt.Println(h.DishChoices)

	data := gin.H{
		"choices":     h.DishChoices.GetUserChoices(),
		"stackedList": h.DishChoices.CreateCompressedList(),
	}
	c.HTML(http.StatusOK, "admin-list.tmpl.html", data)
}

func CreateHostAddress(host, port string) string {
	if port != ":80" {
		return host + port
	}
	return host
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
