//go:generate go run github.com/swaggo/swag/cmd/swag init -g ./api.go
package api

import (
	"context"
	"fmt"
	"github.com/dreamvo/gilfoyle/api/db"
	"github.com/dreamvo/gilfoyle/api/docs"
	"github.com/dreamvo/gilfoyle/api/v1"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
)

// @license.name GNU General Public License v3.0
// @license.url https://github.com/dreamvo/gilfoyle/blob/master/LICENSE

// Serve runs a REST API web server
func Serve(port int) {
	client, err := db.NewClient()
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	docs.SwaggerInfo.Title = "Gilfoyle server"
	docs.SwaggerInfo.Description = " Video streaming server backed by decentralized filesystem."
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", port)
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()

	r.GET("/healthcheck", healthcheckHandler)

	v1.Router(r)

	// register swagger docs handler
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// launch web server
	_ = r.Run(fmt.Sprintf(":%d", port))
}

// Healthcheck godoc
// @Summary Check service status
// @Description get string by ID
// @Success 200
// @Router /healthcheck [get]
func healthcheckHandler(ctx *gin.Context) {
	ctx.Status(200)
}
