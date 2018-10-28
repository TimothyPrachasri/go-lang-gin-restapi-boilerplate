package controller

import (
	service "app/lib/domain/services/coin"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/gin-gonic/gin"
)

type response struct {
	Data responseData `json:"data"`
}

type responseData map[string]responseQuotes

type responseQuotes struct {
	Quotes map[string]responseRepesentativeCurrency
}

type responseRepesentativeCurrency struct {
	PercentChange float32 `json:"percent_change_1h"`
}

type timeChangeStruct struct {
	Key   string
	Value float64
}

func Compare(c *gin.Context) {
	queries := c.DefaultQuery("symbol", "")
	if queries == "" {
		panic("Insufficient Query(s)")
	}
	params := strings.Split(queries, ",")
	resp := service.Coin{}.GetTickerBySymbols(params)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	jsonParsed, _ := gabs.ParseJSON(body)
	timeChangeObj := timeChangeStruct{}
	for _, ech := range params {
		var timeChange float64
		var ok bool
		timeChange, ok = jsonParsed.Search("data", ech, "quote", "USD", "percent_change_1h").Data().(float64)
		if !ok {
			continue
		}
		if ok {
			isMetConditions := (timeChangeStruct{} == timeChangeObj) || (timeChangeObj.Value < timeChange)
			if isMetConditions {
				timeChangeObj.Key = ech
				timeChangeObj.Value = timeChange
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"name": timeChangeObj.Key,
	})
}
