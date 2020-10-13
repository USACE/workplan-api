package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/apex/gateway"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"

	"github.com/USACE/workplan-api/handlers"
	"github.com/USACE/workplan-api/middleware"

	_ "github.com/lib/pq"
)

// Config stores configuration information stored in environment variables
type Config struct {
	SkipJWT       bool
	LambdaContext bool
	DBUser        string
	DBPass        string
	DBName        string
	DBHost        string
	DBSSLMode     string
}

func initDB(connStr string) *sqlx.DB {

	log.Printf("Getting database connection")
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Could not connect to database")
		panic(err)
	}
	if db == nil {
		log.Panicf("database is nil")
	}
	return db
}

// Connection is a database connnection
func Connection(connStr string) *sqlx.DB {
	conn := initDB(connStr)
	return conn
}

func main() {
	//  Here's what would typically be here:
	// lambda.Start(Handler)
	//
	// There were a few options on how to incorporate Echo v4 on Lambda.
	//
	// Landed here for now:
	//
	//     https://github.com/apex/gateway
	//     https://github.com/labstack/echo/issues/1195
	//
	// With this for local development:
	//     https://medium.com/a-man-with-no-server/running-go-aws-lambdas-locally-with-sls-framework-and-sam-af3d648d49cb
	//
	// This looks promising and is from awslabs, but Echo v4 support isn't quite there yet.
	// There is a pull request in progress, Re-evaluate in April 2020.
	//
	//    https://github.com/awslabs/aws-lambda-go-api-proxy
	//
	var cfg Config
	if err := envconfig.Process("workplan", &cfg); err != nil {
		log.Fatal(err.Error())
	}

	db := Connection(
		fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s sslmode=%s binary_parameters=yes",
			cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBHost, cfg.DBSSLMode,
		),
	)

	e := echo.New()
	e.Use(
		middleware.CORS,
		middleware.GZIP,
	)

	// Public Routes
	public := e.Group("")

	// Public Routes
	// NOTE: ALL GET REQUESTS ARE ALLOWED WITHOUT AUTHENTICATION USING JWTConfig Skipper. See appconfig/jwt.go
	public.GET("workplan/commitments", handlers.ListCommitments(db))
	public.POST("workplan/commitments", handlers.CreateCommitment(db))
	public.DELETE("workplan/commitments/:id", handlers.DeleteCommitment(db))
	// Employees
	public.GET("workplan/employees", handlers.ListEmployees(db))
	// Projects
	public.GET("workplan/projects", handlers.ListProjects(db))
	// Timeperiods
	public.GET("workplan/timeperiods", handlers.ListTimeperiods(db))

	if cfg.LambdaContext {
		log.Print("starting server; Running On AWS LAMBDA")
		log.Fatal(gateway.ListenAndServe("localhost:3030", e))
	} else {
		log.Print("starting server")
		log.Fatal(http.ListenAndServe("localhost:3030", e))
	}
}
