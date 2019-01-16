package internal

//func GetAllDishes(document *goquery.Document) []models.Dish {
//
//	dishes := make([]models.Dish, 0, 0)
//
//	mealTypes := []string{
//		"dnevna-ponuda",
//		"pizza",
//		"jela-sa-zara",
//		"jestenine",
//		"lazanje",
//		"palacinke",
//		"prilozi",
//	}
//
//	for _, mealType := range mealTypes {
//		foundDishes := getDishesByType(document, mealType)
//
//		for _, dish := range foundDishes {
//			dishes = append(dishes, dish)
//		}
//
//	}
//
//	return dishes
//}

//func getDishesByType(document *goquery.Document, mealType string) []models.Dish {
//
//	queryString := "#" + mealType
//	var dishes []models.Dish
//
//	dailyDishNodes := document.Find(queryString).Next().Next()
//
//	dailyDishNodes.Children().Each(func(index int, element *goquery.Selection) {
//
//		element.Children().First().Children().Each(func(index int, elem *goquery.Selection) {
//
//			if elem.HasClass("dish-title") {
//				// get the dish name
//				dishName := strings.TrimSpace(elem.Children().Last().Contents().Text())
//
//				// getting the price
//				priceText := strings.TrimSpace(elem.Children().First().Children().Last().Text())
//
//				priceNormal, _ := ParsePrice(priceText)
//
//				dishes = append(dishes, models.Dish{ Name: dishName, Type: mealType, PriceText: priceText, PriceNormal: priceNormal, })
//			}
//		})
//	})
//	return dishes
//}
