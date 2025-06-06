package controller

import (
	"log"
	"net/http"
	"stock-app/fetcher"
	"stock-app/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/stock", GetStock)
	r.GET("/stocks", GetStocks)
	r.PUT("/sync", SyncWithApi)
	r.GET("/recommendation", GetRecommendedStock)
	r.GET("/filtered-stocks", GetQueriedStocks)
}

func SyncWithApi(c *gin.Context) {
	err := fetcher.Sync()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, "Successfully synced")
	}
}

func GetStocks(c *gin.Context) {
	stocks, err := service.GetAllStock()
	if err != nil {
		log.Println("ERROR: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, stocks)
	}
}

func GetRecommendedStock(c *gin.Context) {
	recommendation := service.FindBestPick()

	c.JSON(http.StatusOK, recommendation)
}

func GetStock(c *gin.Context) {
	stock, err := service.GetStock(c.Query("ticker"))
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, stock)
	}
}

func GetQueriedStocks(c *gin.Context) {
	stocks, err := service.GetQueriedStocks(c.Query("search"), c.Query("sortingType"), c.Query("ascending") == "true")
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, stocks)
	}
}
