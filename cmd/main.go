package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/hse-revizor/reports-back/cmd/docs"
	"github.com/hse-revizor/reports-back/internal/common/infra/config"
	"github.com/hse-revizor/reports-back/internal/report/app"
	"github.com/hse-revizor/reports-back/internal/report/infra"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

func main() {
	dev := flag.Bool("dev", false, "Is in the dev mode")
	flag.Parse()

	prod := !*dev

	cfg, err := config.Init(prod)
	if err != nil {
		panic(err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
			log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
			log.Fatal(err)
	}

	// Create report service
	reportService := app.CreateReportService(db)

	// Setup Gin router
	r := gin.Default()
	apiBase := r.Group("/api/v1/")

	infra.CreateReportController(apiBase.Group("reports"), reportService)

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"*"},
	}))

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	listenTo := fmt.Sprint(cfg.ServeEnpoint, ":", strconv.Itoa(int(cfg.ServePort)))
	log.Println("Listening to", listenTo)
	r.Run(listenTo)
}