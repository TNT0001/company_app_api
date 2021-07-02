package main

import (
	"flag"
	"go-api/internal/api/router"
	"go-api/pkg/infrastructure"
	"go-api/pkg/shared/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	//init : Init flag, not use ENV
	_ = flag.String("dirConfig", "/go-api/configs", "Direction config files")
	_ = flag.String("dirMigration", "/go-api/internal/app/db/migrations", "DIR migration file")
	_ = flag.String("logOutput", "stdout", "Logger Output")
	_ = flag.String("logFormat", "tsv", "Logger Format")
	_ = flag.String("logLocation", "/var/logs/", "Logger Location")
	_ = flag.String("port", "5000", "App port")
	_ = flag.String("env", "dev", "Environment")
	_ = flag.String("workdir", "/go-api/", "Work dir")
	_ = flag.String("dbHost", "db", "Database Host")
	_ = flag.String("dbName", "go_test", "Database Name")
	_ = flag.String("dbPort", "3306", "Database Port")
	_ = flag.String("dbUser", "go_test", "Database User")
	_ = flag.String("dbPass", "go_test", "Database Password")
	_ = flag.String("accessSecret", "access secret", "Access Secret Key")

	flag.Parse()
}

func main() {
	port := utils.GetStringFlag("port")

	ginMode := gin.DebugMode
	if utils.GetStringFlag("env") != "dev" && utils.GetStringFlag("env") != "local" {
		ginMode = gin.ReleaseMode
	}

	err := infrastructure.NewConfig()
	if err != nil {
		panic(err)
	}
	// logConfig := infrastructure.LoggerConfig{
	// 	Level:    utils.GetStringFlag(infrastructure.GetConfigString("flags.logger.level")),
	// 	Output:   utils.GetStringFlag(infrastructure.GetConfigString("flags.logger.output")),
	// 	Format:   utils.GetStringFlag(infrastructure.GetConfigString("flags.logger.format")),
	// 	Location: utils.GetStringFlag(infrastructure.GetConfigString("flags.logger.location")),
	// }

	// loggerHandler, err := infrastructure.NewLogger(&logConfig, infrastructure.LogrusLogger)
	// if err != nil {
	// 	log.Fatalf("Could not instantiate log %s", err.Error())
	// }

	databaseHandler, err := infrastructure.NewDatabase(infrastructure.MySQLDatabase)
	if err != nil {
		log.Fatal(nil, err)
	}

	err = databaseHandler.MigrationDB()
	if err != nil {
		log.Println("Migrations DB : " + err.Error())
	}

	defer func() {
		infrastructure.CloseDB(databaseHandler)
	}()

	gin.SetMode(ginMode)
	engine := gin.New()

	r := &router.Router{
		Engine: engine,
		// LoggerHandler: loggerHandler,
		DBHandler: databaseHandler,
	}
	r.InitializeRouter()
	r.SetupHandler()

	engine.Run(":" + port)
}
