package handlers

import (
	"fmt"
	"github.com/djukela17/pinjur-lunch/internal/formatters"
	"github.com/djukela17/pinjur-lunch/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MainHandler struct {
	AllDishList     models.DishCollection
	AvailableDishes models.DishCollection
	DishChoices     models.UserChoices

	NameList    map[string]string
	HostAddress string
}

func NewMainHandler(allDishList models.DishCollection, nameList map[string]string, host, port string) MainHandler {
	return MainHandler{
		AllDishList: allDishList,

		NameList: nameList,

		HostAddress: CreateHostAddress(host, port),
	}
}

func (h *MainHandler) AdminCreateForm(c *gin.Context) {
	content := gin.H{
		"hostAddress": h.HostAddress,
		"dishes":      h.AllDishList.GetAll(),
	}
	c.HTML(http.StatusOK, "admin.tmpl.html", content)
}

func (h *MainHandler) CreateTodayMealList(c *gin.Context) {
	fmt.Println("Create today meal list")

	for i := 0; i < len(h.AllDishList.GetAll()); i++ {
		dishName := c.PostForm("dish_" + strconv.Itoa(i))
		if dishName != "" {
			dish, err := h.AllDishList.GetDish(dishName)
			if err == nil {
				h.AvailableDishes.AddDish(dish)
			}
		}
	}

	content := gin.H{
		"dishes": h.AvailableDishes.GetAll(),
	}
	c.HTML(http.StatusOK, "admin-created.tmpl.html", content)
}

func (h *MainHandler) UserForm(c *gin.Context) {
	fmt.Println(c.ClientIP())

	if len(h.AllDishList.GetAll()) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
		return
	}
	data := gin.H{
		"hostAddress": h.HostAddress,
		"dishes":      h.AllDishList.GetAll(),
		"name":        h.NameList[c.ClientIP()],
	}
	c.HTML(http.StatusOK, "user.tmpl.html", data)
}

func (h *MainHandler) UpdateActiveDishList(c *gin.Context) {
	fmt.Println("Updating dish choices")

	chosenDish := c.PostForm("dish")
	name := c.PostForm("name")
	optionalNote := c.PostForm("optional-note")

	fmt.Println("Client IP:", c.ClientIP())
	fmt.Println("Selected dish:", chosenDish)
	fmt.Println("name:", name)
	fmt.Println("extra note:", optionalNote)

	if dish, err := h.AllDishList.GetDish(chosenDish); err == nil {
		h.DishChoices.AddDish(dish, name, optionalNote)
		data := gin.H{
			"name":         name,
			"chosenDish":   chosenDish,
			"optionalNote": optionalNote,
			"orderStatus":  "success",
		}
		c.HTML(http.StatusOK, "user-submitted.tmpl.html", data)
		return
	}

	data := gin.H{
		"name":         name,
		"chosenDish":   chosenDish,
		"optionalNote": optionalNote,
		"orderStatus":  "fail",
	}
	c.HTML(http.StatusNotFound, "user-submitted.tmpl.html", data)
}

func (h *MainHandler) ListActiveChoices(c *gin.Context) {

	fmt.Println(formatters.DisplayPrice(h.DishChoices.CalcTotalPrice()))
	data := gin.H{
		"totalAmount": formatters.DisplayPrice(h.DishChoices.CalcTotalPrice()),
		"dishes":      h.AllDishList.GetAll(),
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
