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
	Maria  *sql.DB
	Mongo  *mongo.Client
	Redis  *redis.Client
	Log    *logrus.Logger
}

func init() {
	AppInit()
}

func AppInit() {
	App = &Application{}
	App.Config = InitConfig()
	App.Maria = InitMariaDatabase()
	App.Mongo = InitMongoDatabase()
	App.Redis = InitRedis()
	App.Log = InitLogger()
}
