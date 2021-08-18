package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/hussar_company/vstazherstve_back/dbcontext"
	"github.com/hussar_company/vstazherstve_back/user"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func buildHandler(r *mux.Router, db *dbcontext.DB) {
	user.RegisterHandlers(r, user.NewService(user.NewRepository(db)))
}

func init() {
	viper.SetConfigName("appconfig")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("env", "local")
	viper.SetDefault("port", "8080")

	err := viper.ReadInConfig()
	if err != nil {
		log.WithError(err).Error("Ошибка при чтении файла конфигурации")
	}

	environment := viper.GetString("env")

	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)

	if environment == "dev" {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	log.Info("Application starting.")

	connectionString := viper.GetString("connectionString")
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.WithError(err).Fatal("Ошибка при создании клиент")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	defer client.Disconnect(ctx)
	if err != nil {
		log.WithError(err).Fatal("Ошибка при подключении к БД")
	}

	db := client.Database(viper.GetString("database"))
	router := mux.NewRouter()
	buildHandler(router, dbcontext.New(db))

	httpError := http.ListenAndServe(":"+viper.GetString("PORT"), router)
	log.WithError(httpError).Error("Ошибка во время работы сервера")
}
