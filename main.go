package main

import (
	"encoding/xml"
	"time"

	"github.com/gin-gonic/gin"
)

var recipes []Recipe

type Recipe struct {
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

func main() {
	router := gin.Default()

	router.GET("/", IndexHandler)

	router.Run(":5000")
}

func IndexHandler(c *gin.Context) {
	c.XML(200, Person{
		FirstName: "Cossette",
		LastName:  "Potrait"},
	)
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}
