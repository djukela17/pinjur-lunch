package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/djukela17/pinjur-lunch/internal/formatters"
	"github.com/djukela17/pinjur-lunch/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/http"
	"strconv"
	"time"
)

type MainHandler struct {
	// constants
	HostAddress string
	MongoURI    string

	AllDishCollectionName string
	DatabaseName          string

	MongoClient *mongo.Client

	AllDishList     models.DishCollection
	AvailableDishes models.DishCollection
	DishChoices     models.Orders

	SideDishes models.MealAdditionsCollection
	Extras     models.MealAdditionsCollection

	Extras2 models.MealAdditions2Collection

	NameList map[string]string
}

func NewMainHandler(nameList map[string]string, serveHost, servePort, mongoURI string) MainHandler {
	return MainHandler{
		MongoURI:    mongoURI,
		HostAddress: CreateHostAddress(serveHost, servePort),

		AllDishCollectionName: "all-dishes",
		DatabaseName:          "pinjur",

		NameList: nameList,
	}
}

func (h *MainHandler) Init() error {

	client, err := mongo.NewClient(h.MongoURI)
	if err != nil {
		return err
	}
	// check the mongo connection
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Connect(ctx); err != nil {
		return err
	}

	h.MongoClient = client

	// Get the data from the mongo database
	dishes, err := h.GetAllDishes()
	if err != nil {
		return nil
	}
	// Load from a file
	//dishes, err := models.LoadDishList("data/discounted-prices.json")

	// Load data from mongo db
	//h.SideDishes = models.NewAdditionsCollection(h.DatabaseName, "sideDishes")
	//if err := h.SideDishes.LoadAll(h.MongoClient); err != nil {
	//	return err
	//}
	// Load data from json file
	sideDishes, err := models.NewAdditionsCollectionFromFile("data/side-dishes.json")
	if err != nil {
		return err
	}
	h.SideDishes = sideDishes

	// Load data from mongo db
	//h.Extras = models.NewAdditionsCollection(h.DatabaseName, "extras")
	//if err := h.Extras.LoadAll(h.MongoClient); err != nil {
	//	return err
	//}
	// Load data from json file
	extras, err := models.NewAdditionsCollectionFromFile("data/extras.json")
	if err != nil {
		return err
	}
	h.Extras = extras

	h.AllDishList = models.NewDishCollection(dishes)

	// Replace the old items in mongo db with new ones
	// should be used when loading list from json file
	//if err := h.AllDishList.InsertAll(
	//	h.MongoClient, h.DatabaseName, h.AllDishCollectionName, true);
	//err != nil {
	//	return err
	//}

	// Assign side dishes and extras to dishes
	lst, err := models.LoadFromFile("data/additions.json")
	if err != nil {
		return err
	}
	h.Extras2 = models.NewMealAdditions2Collection(lst)

	// Assign additions to meals
	for _, d := range h.AllDishList.GetAll() {
		switch d.Type {
		case "sendvic":
			d.Extras2 = h.Extras2
			fmt.Println(d.Extras2)
		default:
			fmt.Println("no type")
		}
	}

	fmt.Println(h.GetAllDishes())

	return nil
}

func (h *MainHandler) AdminCreateForm(c *gin.Context) {
	content := gin.H{
		"hostAddress": h.HostAddress,
		"dishes":      h.AllDishList.GetAll(),
	}
	c.HTML(http.StatusOK, "admin.tmpl.html", content)
}

func (h *MainHandler) CreateTodayMealList(c *gin.Context) {

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

	if len(h.AvailableDishes.GetAll()) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
		return
	}

	dj, err := json.Marshal(h.AllDishList.GetAll())
	if err != nil {
		fmt.Println(err)
	}

	data := gin.H{
		"hostAddress": h.HostAddress,
		"dishes":      h.AvailableDishes.GetAll(),
		"dishesJson":  string(dj),
		"name":        h.NameList[c.ClientIP()],
		"sideDishes":  h.SideDishes,
		"extras":      h.Extras,
	}
	c.HTML(http.StatusOK, "user.tmpl.html", data)
}

func (h *MainHandler) UpdateActiveDishList(c *gin.Context) {

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

	data := gin.H{
		"totalAmount": formatters.DisplayPrice(h.DishChoices.CalcTotalPrice()),
		"dishes":      h.AvailableDishes.GetAll(),
		"choices":     h.DishChoices.GetOrders(),
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
