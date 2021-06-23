package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/PutskouDzmitry/be-sd/pkg/api"
	"github.com/PutskouDzmitry/be-sd/pkg/data"

	"github.com/gorilla/mux"
)

var (
	user     = os.Getenv("DB_USERS_USER")
	password = os.Getenv("DB_USERS_PASSWORD")
	host = os.Getenv("DB_USERS_HOST")
	port = os.Getenv("DB_USER_PORT")
)

func initValues() {
	if user == "" {
		user = "root"
	}
	if password == "" {
		password = "example"
	}
	if host == "" {
		host = "mongo"
	}
	if port == "" {
		port = "27017"
	}
}

func initClient(user string, password string, host string, port string) string{
	return fmt.Sprintf("mongodb://%v:%v@%v:%v/?sslmode=disable", user, password, host, port)
}

func main() {
	initValues()
	client, err := mongo.NewClient(options.Client().ApplyURI(initClient(user, password, host, port)))
	if err != nil {
		logrus.Fatal("error with client ", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		logrus.Fatal("error with connect to db ", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logrus.Fatal(err)
	}
	db := client.Database("book")
	collection := db.Collection("book")
	userData := data.NewBookData(collection)
	// 2. create router that allows to set routes
	r := mux.NewRouter()
	// 4. send data layer to api layer
	api.ServeUserResource(r, *userData)
	// 5. cors for making requests from any domain
	r.Use(mux.CORSMethodMiddleware(r))
	// 6. start server
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Server Listen port...", err)
	}
	if err := http.Serve(listener, r); err != nil {
		log.Fatal("Server has been crashed...")
	}
}

