package api

import (
	"fmt"

	"github.com/adesokanayo/readcommend/internal/api/router"
	"github.com/adesokanayo/readcommend/internal/pkg/config"
	"github.com/adesokanayo/readcommend/internal/pkg/db"
	"github.com/gin-gonic/gin"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	db.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "data"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()
	err := web.Run(":" + conf.Server.Port)
	if err != nil{
		fmt.Println("==================> API Running on port " + conf.Server.Port)
	} else {
		
	}

}
