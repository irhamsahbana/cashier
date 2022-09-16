package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/fatih/color"
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMariaDatabase() *sql.DB {
	dbHost := App.Config.GetString(`mariadb.host`)
	dbPort := App.Config.GetString(`mariadb.port`)
	dbUser := App.Config.GetString(`mariadb.user`)
	dbPass := App.Config.GetString(`mariadb.pass`)
	dbName := App.Config.GetString(`mariadb.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, err := sql.Open(`mysql`, dsn)

	dbConn.SetMaxIdleConns(10)
	dbConn.SetMaxOpenConns(100)
	dbConn.SetConnMaxIdleTime(5 * time.Minute)
	dbConn.SetConnMaxLifetime(1 * time.Hour)

	errCheck("MariaDB", err)

	err = dbConn.Ping()
	errCheck("MariaDB", err)

	color.Green(fmt.Sprintf("connected to MariaDB from %s:%s", dbHost, dbPort))
	return dbConn
}

func InitMongoDatabase() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := App.Config.GetString(`mongo.host`)
	dbPort := App.Config.GetString(`mongo.port`)
	dbUser := App.Config.GetString(`mongo.user`)
	dbPass := App.Config.GetString(`mongo.pass`)
	dbName := App.Config.GetString(`mongo.name`)

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}

	var client *mongo.Client
	var err error
	var debugMode bool = App.Config.GetBool("mongodb.monitor_query")

	if debugMode {
		cmdMonitor := &event.CommandMonitor{
			Started: func(_ context.Context, evt *event.CommandStartedEvent) {
				color.Yellow(evt.Command.String())
			},
		}

		client, err = mongo.NewClient(options.Client().ApplyURI(mongodbURI).SetMonitor(cmdMonitor))
	} else {
		client, err = mongo.NewClient(options.Client().ApplyURI(mongodbURI))
	}

	errCheck("MongoDB", err)

	err = client.Connect(ctx)
	errCheck("MongoDB", err)

	err = client.Ping(ctx, readpref.Primary())
	errCheck("MongoDB", err)

	color.Green(fmt.Sprintf("connected to MongoDB from %s:%s", dbHost, dbPort))
	defer initMongoDatabaseIndexes(ctx, client, dbName)

	return client
}

func InitRedis() *redis.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := App.Config.GetString(`redis.host`)
	dbPort := App.Config.GetString(`redis.port`)
	dbPass := App.Config.GetString(`redis.pass`)
	dbName := App.Config.GetInt(`redis.name`)

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", dbHost, dbPort),
		Password: dbPass,
		DB:       dbName,
	})

	_, err := client.Ping(ctx).Result()
	errCheck("Redis", err)

	color.Green(fmt.Sprintf("connected to Redis from %s:%s", dbHost, dbPort))
	return client
}

func errCheck(prefix string, err error) {
	if err != nil {
		color.Red(fmt.Sprintf("%s: %s", prefix, err.Error()))
		log.Fatal(err)
	}
}
