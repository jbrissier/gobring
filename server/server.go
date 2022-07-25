package server

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
		c.JSON(404, gin.H{"Error": "Not found or DB Error"})
		return
	}

	c.JSON(200, gin.H{"Success": "Deleted"})

}

func AllBrings(c *gin.Context) {

	brings, _ := model.GetAllBrings()

	c.JSON(200, brings)

}

func GetBring(c *gin.Context) {

	id := c.Param("id")
	idStr, _ := strconv.ParseUint(id, 10, 32)

	bring, err := model.FindBring((uint)(idStr))
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not found or DB Error"})
		return
	}

	c.JSON(200, bring)

}

func CorsMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Request-Method", "POST, GET, PUT, OPTIONS, DELETE")
	}
}

func StarServer() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("static", true)))
	r.Use(static.Serve("/public", static.LocalFile("svelte-app/public", false)))
	r.Use(CorsMiddelware())
	r.GET("api/bring", AllBrings)
	r.GET("api/bring/:id", GetBring)
	r.POST("api/bring", CreateBring)
	r.DELETE("api/bring", DeleteBring)

	r.Run()
}
