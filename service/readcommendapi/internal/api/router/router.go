package router

import (
	"fmt"
	"github.com/adesokanayo/readcommend/internal/api/controllers"
	"github.com/adesokanayo/readcommend/internal/api/middlewares"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Setup() *gin.Engine {
	app := gin.New()

	//==================Logging to file.
	f, _ := os.Create("readcommendapi.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f,os.Stdout)

	//================== Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())
    //================== Routes
	app.GET("/api/v1/books", controllers.GetBooks)
	app.GET("/api/v1/authors", controllers.GetAuthors)
	app.GET("/api/v1/genres", controllers.GetGenres)
	app.GET("/api/v1/sizes", controllers.GetSizes)
	app.GET("/api/v1/eras", controllers.GetEras)
	return app
}
