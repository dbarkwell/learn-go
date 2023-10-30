package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "learn-go.barkwell.com/docs"
)

func initDB(dsn string) *sqlx.DB {
	if dsn == "" {
		panic("Missing MySQL connection string")
	}

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db
}

// @title           Learning Go
// @version         1.0
// @description     Project to learn Go.
// @host localhost:8080
// @BasePath /api/v1
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	dsn := flag.String("dsn", "", "MySQL connection string")
	flag.Parse()

	db := initDB(*dsn)
	defer db.Close()

	albumApi := initAlbumAPI(db)

	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000"},
		AllowMethods:  []string{"GET", "POST", "DELETE"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	v1 := r.Group("/api/v1")
	{
		album := v1.Group("/albums")
		{
			album.GET("", albumApi.FindAll)
			album.GET("/:id", albumApi.FindByID)
			album.POST("", albumApi.Add)
			album.DELETE("/:id", albumApi.Remove)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run("localhost:8080")
}
