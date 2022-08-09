package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type market struct{
	ID			string `json:"id"`
	Title		string `json:title`
	Factory		string `json:factory`
	Quantity	int `json:quantity`
}

var markets = []market{
	{ID: "1", Title: "Area A", Factory: "Asia", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Factory: "America", Quantity: 5},
	{ID: "3", Title: "Thunder Tropic", Factory: "Russia", Quantity: 6},
}


func getMarkets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, markets)
}

func marketById(c *gin.Context) {
	id := c.Param("id")
	market, err := getMarketById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Not Found for {id}"})
		return
	}

	c.IndentedJSON(http.StatusOK, market)
}


func addQuantity(c *gin.Context){
	id := c.Param("id")

	market, err := getMarketById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Parameters Error"})
		return	
	}

	if market.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Sold Out"})
		return	
	}

	market.Quantity += 1
	c.IndentedJSON(http.StatusOK, market)
}

func deleteQuantity(c *gin.Context) {
	id := c.Param("id")

	market, err := getMarketById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Parameters Error"})
		return	
	}

	if market.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Sold Out"})
		return	
	}

	market.Quantity -= 1
	c.IndentedJSON(http.StatusOK, market)
}

func getMarketById(id string) (*market, error) {
	for i, b := range markets {
		if b.ID == id {
			return &markets[i], nil
		}
	}

	return nil, errors.New("Not Found")
}

func createBook(c * gin.Context) {
	var newMarket market

	if err:= c.BindJSON(&newMarket); err != nil {
		return
	}

	markets = append(markets, newMarket)
	c.IndentedJSON(http.StatusCreated, newMarket)
}

func main() {
	router := gin.Default()
	router.GET("/markets", getMarkets)
	router.GET("/markets/:id", marketById)
	router.POST("/markets", createBook)
	router.PATCH("/delete/:id", deleteQuantity)
	router.PATCH("/add/:id", addQuantity)
	router.Run("localhost:6666")

}