package main

import (
	"log"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	//Setting config paths
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	conn, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		logrus.Fatalf("error initializing DB: %s", err.Error())
	}

	repos := repository.NewRepository(conn)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)
	log.Println("server URL localhost:" + viper.GetString("port"))
	srv.Run(viper.GetString("port"), handler.InitRoutes())

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
