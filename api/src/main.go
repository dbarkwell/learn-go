package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
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

	albumAPI := initAlbumAPI(db)
	//userApi := initUserAPI(db)
	authAPI := initAuthenticationAPI(db)

	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(static.Serve("/", static.LocalFile("../../ui/build", true)))
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"https://aa0c-99-228-180-110.ngrok-free.app"},
		AllowMethods:  []string{"GET", "POST", "DELETE"},
		AllowHeaders:  []string{"Origin", "Content-Type", "X-Requested-With"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	v1 := r.Group("/api/v1")
	{
		album := v1.Group("/albums")
		{
			album.GET("", albumAPI.FindAll)
			album.GET("/:id", albumAPI.FindByID)
			album.POST("", albumAPI.Add)
			album.DELETE("/:id", albumAPI.Remove)
		}
		auth := v1.Group("/auth")
		{
			auth.GET("/registerRequest", authAPI.BeginRegistration)
			auth.POST("/registerResponse", authAPI.FinishRegistration)
			auth.GET("/signinRequest", authAPI.BeginLogin)
			auth.POST("/signinResponse", authAPI.FinishLogin)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
