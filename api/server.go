package api

import (
	"fmt"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jbrissier/gobring/model"
)

func CreateBring(c *gin.Context) {

	var bring model.Bring

	c.BindJSON(&bring)
	fmt.Printf("Got Bring %v", bring)
	bring.Save()

}

func DeleteBring(c *gin.Context) {

	id := c.Param("id")
	idStr, _ := strconv.ParseUint(id, 10, 32)

	err := model.DeleteBring((uint)(idStr))
	if err != nil {
		c.JSON(404, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Success": "Deleted"})

}

func AllBrings(c *gin.Context) {

	brings, _ := model.GetAllBrings()

	c.JSON(200, brings)

}

func GetSingelBring(id string) (*model.Bring, error) {
	idStr, _ := strconv.ParseUint(id, 10, 32)

	return model.FindBring((uint)(idStr))
}

func GetBring(c *gin.Context) {

	bring, err := GetSingelBring(c.Param("id"))

	if err != nil {
		c.JSON(404, gin.H{"Error": "Not found or DB Error"})
		return
	}

	c.JSON(200, bring)

}

func CreateBringItem(c *gin.Context) {
	bring, err := GetSingelBring(c.Param("id"))
	bringItem := model.BringItem{}
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not found or DB Error"})
		return
	}

	err = c.BindJSON(&bringItem)

	if err != nil {
		c.JSON(500, gin.H{"Error": "Can't parse JSON"})
		return
	}

	err = model.AddBrintItem(bring, &bringItem)

	if err != nil {
		c.JSON(500, gin.H{"Error": "Can't add item"})
		return
	}
	c.JSON(200, gin.H{"Success": "Added"})

}

func GetBringItems(c *gin.Context) {
	bring, err := GetSingelBring(c.Param("id"))

	if err != nil {
		c.JSON(404, gin.H{"Error": "Not found or DB Error"})
		return
	}

	items, err := bring.GetItems()

	if err != nil {
		c.JSON(500, gin.H{"Error": "DB Error"})
		return
	}

	c.JSON(200, items)
}

func DeleteBringItem(c *gin.Context) {

	idStr, _ := strconv.ParseUint(c.Param("itemId"), 10, 32)
	err := model.DeleteBringItem((uint)(idStr))
	if err != nil {
		c.JSON(404, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Success": "Deleted"})
}

func CorsMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Request-Method", "POST, GET, PUT, OPTIONS, DELETE")
	}
}

func StarServer() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("svelte-app/public", false)))
	r.Use(CorsMiddelware())

	api := r.Group("/api")

	api.GET("bring", AllBrings)
	api.POST("bring", CreateBring)
	api.GET("bring/:id", GetBring)
	api.DELETE("bring/:id", DeleteBring)

	// add and remove items to bring

	api.POST("bring/:id", CreateBringItem)
	api.GET("bring/:id/items", GetBringItems)
	api.DELETE("bring/:id/items/:itemid", DeleteBringItem)
	//api.DELETE("bring/:id", DeleteBringItem)

	r.Run()
}