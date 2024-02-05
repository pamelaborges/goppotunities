package main

import (
	"github.com/pamelaborges/goppotunities/configs"
	"github.com/pamelaborges/goppotunities/router"
)

var (
	logger *configs.Logger
)

func init() {
	logger = configs.GetLogger()

}

func main() {
	router.Init()
	error := configs.Init()
	if error != nil {
		logger.Errorf("Erro ao iniciar DB %s", error)
		return
	}
}
