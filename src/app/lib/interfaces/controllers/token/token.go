package controller

import (
	service "app/lib/domain/services/coin"
	"fmt"
	"io/ioutil"
	"net/http"

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
	// coinService := service.Coin{}
	// resp := coinService.GetTickerBySymbols()
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// if resp.StatusCode != http.StatusOK {
	// 	c.Status(http.StatusServiceUnavailable)
	// 	return
	// }
	// var responseObject response
	// json.Unmarshal(body, &responseObject)
	// fmt.Println(responseObject.Data["BTC"].Quotes)
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "pong",
	// })
	params := []string{"ETC", "BTC"}
	resp := service.Coin{}.GetTickerBySymbols(params)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	jsonParsed, _ := gabs.ParseJSON(body)
	fmt.Println(jsonParsed)
	timeChangeObj := timeChangeStruct{}
	for _, ech := range params {
		var timeChange float64
		var ok bool
		fmt.Println(ech)
		timeChange, ok = jsonParsed.Search("data", ech, "quote", "USD", "percent_change_1h").Data().(float64)
		fmt.Println(timeChange, ok)
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
	fmt.Println(timeChangeObj)
	c.JSON(http.StatusOK, gin.H{
		"name": timeChangeObj.Key,
	})
}
