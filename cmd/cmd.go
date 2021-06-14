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

	"github.com/PutskouDzmitry/golang-training-Library/pkg/api"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/data"

	"github.com/gorilla/mux"
)

var (
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
)

func init() {
	if user == "" {
		user = "kvarc"
	}
	if dbname == "" {
		dbname = "myFirstDatabase"
	}
	if password == "" {
		password = "pilubadima"
	}
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprint("mongodb+srv://", user, ":", password, "@clusterkvarc.bdz6v.mongodb.net/", dbname, "?retryWrites=true&w=majority")))
	if err != nil {
		logrus.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		logrus.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logrus.Fatal(err)
	}
	db := client.Database("myFirstDatabase")
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

