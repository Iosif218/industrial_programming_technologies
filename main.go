package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type tea struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Country     string `json:"country"`
	Description string `json:"description"`
	Cost        string `json:"cost"`
}

var teas = []tea{
	{ID: "1", Name: "Lapsang Souchong Black Tea", Type: "Black", Country: "China", Description: "From the Fujian province of China comes this very distinctive tea with its exotic smoky flavor. After plucking, the leaves are withered over cypress or pine wood fires. After the rolling process, they are placed into wooden barrels until they begin to emit their own pleasant aroma. As a final step they are placed in bamboo baskets and hung on racks over smoky pine fires where they dry and absorb the essence of the smoke.", Cost: "$7.95"},
	{ID: "2", Name: "Glenburn Second Flush Darjeeling", Type: "Black", Country: "India", Description: "A rare, fragrant second flush Darjeeling from the foothills of the Himalayas in India brews to a lovely amber color. The taste is exquisite and complex, with a refined malty and muscatel character, and a fruity, floral finish.", Cost: "$14.95"},
	{ID: "3", Name: "Dragon Phoenix Pearl Jasmine Tea", Type: "Green", Country: "China", Description: "Finest downy green leaves are hand-rolled into small spheres and scented with fragrant jasmine blossoms. It brews up an exquisite cup with an alluring aroma and delicate sweet floral taste.", Cost: "$18.95"},
}

func main() {
	router := gin.Default()

	router.GET("/teas", getTeas)

	router.GET("/teas/:id", getTeaByID)

	router.POST("/teas", createTea)

	router.PUT("/teas/:id", updateTea)

	router.DELETE("/teas/:id", deleteTea)

	router.Run(":8080")
}

func getTeas(c *gin.Context) {
	c.JSON(http.StatusOK, teas)
}

func getTeaByID(c *gin.Context) {
	id := c.Param("id")

	for _, tea := range teas {
		if tea.ID == id {
			c.JSON(http.StatusOK, tea)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "tea not found"})
}

func createTea(c *gin.Context) {
	var newTea tea

	if err := c.BindJSON(&newTea); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	teas = append(teas, newTea)
	c.JSON(http.StatusCreated, newTea)
}

func updateTea(c *gin.Context) {
	id := c.Param("id")
	var updatedTea tea

	if err := c.BindJSON(&updatedTea); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, tea := range teas {
		if tea.ID == id {
			teas[i] = updatedTea
			c.JSON(http.StatusOK, updatedTea)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func deleteTea(c *gin.Context) {
	id := c.Param("id")

	for i, tea := range teas {
		if tea.ID == id {
			teas = append(teas[:i], teas[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "tea deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "tea not found"})
}
