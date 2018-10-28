package controller

import (
	service "app/lib/domain/services/coin"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/gin-gonic/gin"
)

type timeChangeStruct struct {
	Key   string  // name of symbol
	Value float64 // time change per hour of this symbol
}

// Compare is a function for comparing token symbols based on query of Context.
// The function will be returned as a json represent name of token which has the most price change.
func Compare(c *gin.Context) {
	queries := c.DefaultQuery("symbol", "")
	if queries == "" {
		panic("Insufficient Query(s)")
	}
	params := strings.Split(strings.ToUpper(queries), ",")
	resp := service.Coin{}.GetTickerBySymbols(params)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	jsonParsed, _ := gabs.ParseJSON(body)
	timeChangeObj := timeChangeStruct{}
	for _, ech := range params {
		timeChange, timeChangeOk := jsonParsed.Search("data", ech, "quote", "USD", "percent_change_1h").Data().(float64)
		name, nameOk := jsonParsed.Search("data", ech, "name").Data().(string)
		if !nameOk || !timeChangeOk {
			continue
		}
		isMetConditions := (timeChangeStruct{} == timeChangeObj) || (timeChangeObj.Value < timeChange)
		if isMetConditions {
			timeChangeObj.Key = name
			timeChangeObj.Value = timeChange
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"name": timeChangeObj.Key,
	})
}
