package bootstrap

import (
	"database/sql"

	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	App *Application
)

type Application struct {
	Config *viper.Viper
	Log    *logrus.Logger
	Maria  *sql.DB
	Mongo  *mongo.Client
	Redis  *redis.Client
}

func init() {
	AppInit()
}

func AppInit() {
	App = &Application{}
	App.Config = InitConfig()
	App.Log = InitLogger()
	// App.Maria = InitMariaDatabase()
	App.Mongo = InitMongoDatabase()
	// App.Redis = InitRedis()
}
