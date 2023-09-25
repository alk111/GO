package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient() *mongo.Client {
	connection_string := os.Getenv("CONNECTION_STRING")
	database := os.Getenv("DATABASE")
	if connection_string == "" || database == "" {
		log.Fatalln("Unable to connect to DB. Error: config values cannot be null.")
	}
	clientOptions := options.Client().ApplyURI(connection_string)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	// fmt.Println(connection_string)
	// fmt.Println(database)
	// config := &bongo.Config{
	// 	ConnectionString: connection_string,
	// 	Database:         database,
	// }
	// connection, err := bongo.Connect(config)
	if err != nil {
		log.Fatalln("Unable to connect to database. Error:" + err.Error())
	}
	fmt.Println("Database connected!!!")
	return client
}
