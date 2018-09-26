package db

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/MihaiLupoiu/money-transferring-simulation/backend/libs/util"
	"github.com/MihaiLupoiu/money-transferring-simulation/backend/models/task"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Init creates a connection to postgresql database and
// migrates any new models
func Init() {

	// GET configuration
	configFilePath := flag.String("configFile", "./config.json", "JSON config file to read.")
	flag.Parse()
	config := util.GetConfigurationFile(*configFilePath)

	// user := getEnv("PG_USER", "")
	// password := getEnv("PG_PASSWORD", "")
	// host := getEnv("PG_HOST", "postgresql")
	// port := getEnv("PG_PORT", "5432")
	// database := getEnv("PG_DB", "tasks")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	fmt.Println(dbinfo)

	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")

	if !db.HasTable(&models.Task{}) {
		err := db.CreateTable(&models.Task{})
		if err != nil {
			log.Println("Table already exists")
		}
	}

	db.AutoMigrate(&models.Task{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
