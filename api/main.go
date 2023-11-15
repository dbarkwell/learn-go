package main

import (
	"flag"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	auth "learn-go.barkwell.com/authentication"
	_ "learn-go.barkwell.com/docs"
	"strings"
)

func initCache(cacheCfg string) *memcache.Client {
	if cacheCfg == "" {
		panic("Missing Memcached connection string")
	}

	return memcache.New(cacheCfg)
}

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
	rpID := flag.String("rpID", "", "RP ID")
	rpOrigin := flag.String("rpOrigin", "", "RP Origin")
	cacheCfg := flag.String("memcache", "", "Memcache config")
	flag.Parse()

	db := initDB(*dsn)
	memcache := initCache(*cacheCfg)
	defer db.Close()

	config := auth.AuthnConfig{RPID: *rpID, RPOrigin: *rpOrigin}
	albumAPI := initAlbumAPI(db)
	userAPI := initUserAPI(db, memcache)
	authAPI := initAuthenticationAPI(db, memcache, config)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(static.Serve("/", static.LocalFile("./dist", true)))
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "DELETE"},
		AllowHeaders:  []string{"Origin", "Content-Type", "X-Requested-With", "Content-Length"},
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
			auth.POST("/registerRequest/:username", authAPI.BeginRegistration)
			auth.POST("/registerResponse/:username", authAPI.FinishRegistration)
			auth.GET("/signinRequest/:username", authAPI.BeginLogin)
			auth.POST("/signinResponse/:username", authAPI.FinishLogin)
		}
		user := v1.Group("/users")
		{
			user.POST("/", userAPI.Add)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/api") {
			c.File("./dist/index.html")
		}
	})

	r.Run()
}
