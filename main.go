package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func PrepareConfig() {
	config := "./config/dev.yml"
	viper.SetConfigType("yaml")
	viper.SetConfigFile(config)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Cannot read a config file: %v\n", err)
	}
}

var R *mux.Router
var onceRouter sync.Once

func GetRouter() *mux.Router {
	onceRouter.Do(func() {
		R = mux.NewRouter()
	})
	return R
}

func main() {
	// var err error
	PrepareConfig()
	// Check for connection
	// defer GetDbAdapter().Close()

	r := GetRouter()
	ConfigureHandlers(r)
	fmt.Println("Server is starting...")
	http.ListenAndServe(":"+viper.GetString("port"), r)
}
