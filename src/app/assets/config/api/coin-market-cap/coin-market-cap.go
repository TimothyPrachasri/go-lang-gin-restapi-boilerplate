package config

import (
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Key string
}

func CoinMarketCap() string {
	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		panic(err)
	}
	return configuration.Key
}

func getFileName() string {
	filename := []string{"coin-market-cap", ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))
	return filePath
}
