package models

import (
	"context"
	_ "database/sql"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/logging"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

//redis client
//var redisDB *redis.Client
var mysqlDB *gorm.DB

const redisPrefix = "ynov:"

func init() {
	ctx := context.Background()
	// Sets your Google Cloud Platform project ID.
	projectID := "ynov-api"

	// Creates a client.
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the name of the log to write to.
	logName := "api-log"

	logger := client.Logger(logName).StandardLogger(logging.Info)

	//initializing redis client
	/*redisAddr, redisPassword := viper.GetString("redis.addr"), viper.GetString("redis.password")
	if redisAddr != "" {
		redisDB = redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: redisPassword,                // no password set
			DB:       viper.GetInt("redis.db_idx"), // use default DB
		})
		if pong, err := redisDB.Ping().Result(); err != nil || pong != "PONG" {
			//logrus.WithError(err).Fatal("could not connect to the redis server")
			logger.Println("could not connect to the redis server")
		}
	}*/
	conn := os.Getenv("MYSQL_CONNECTION")
	//init mysql
	/*conn := fmt.Sprintf("%s:%s@cloudsql(%s)/%s?charset=%s&parseTime=True&loc=Local", viper.GetString("mysql.user"),
	viper.GetString("mysql.password"), viper.GetString("mysql.instance_connection_name"), viper.GetString("mysql.database"),
	viper.GetString("mysql.charset"))*/

	if db, err := gorm.Open("mysql", conn); err == nil {
		mysqlDB = db
	} else {
		//logrus.WithError(err).Fatalln("initialize mysql database failed")
		fmt.Sprintf("initialize mysql database failed")
		logger.Println("initialize mysql database failed")
	}
	//enable Gorm mysql log
	if flag := viper.GetBool("app.enable_sql_log"); flag {
		mysqlDB.LogMode(flag)
		//f, err := os.OpenFile("mysql_gorm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		//if err != nil {
		//	logrus.WithError(err).Fatalln("could not create mysql gorm log file")
		//}
		//logger :=  New(f,"", Ldate)
		//mysqlDB.SetLogger(logger)
	}
	//mysqlDB.AutoMigrate()

}

//Close clear db collection
func Close() {
	if mysqlDB != nil {
		mysqlDB.Close()
	}
	/*if redisDB != nil {
		redisDB.Close()
	}*/
}
