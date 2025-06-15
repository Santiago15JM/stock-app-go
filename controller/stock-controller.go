package controller

import (
	"net/http"
	"stock-app/fetcher"
	"stock-app/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/stock", GetStock)
	r.GET("/stocks", GetStocks)
	r.PUT("/sync", SyncWithApi)
	r.GET("/recommendation", GetRecommendedStock)
	r.GET("/query-stocks", GetQueriedStocks)
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
	limitStr := c.DefaultQuery("limit", "50")
	pageStr := c.DefaultQuery("page", "0")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)

	stocks, err := service.GetAllStock(limit, page)
	if err != nil {
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
	limitStr := c.DefaultQuery("limit", "50")
	pageStr := c.DefaultQuery("page", "0")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	sorting := c.DefaultQuery("sortingType", "TICKER")
	if sorting == "" {
		sorting = "TICKER"
	}

	stocks, err := service.GetQueriedStocks(
		c.Query("search"),
		sorting,
		c.DefaultQuery("ascending", "true") == "true",
		limit,
		page)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, stocks)
	}
}
