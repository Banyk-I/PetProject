package mongo

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectMongoDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI(viper.GetString("mongodb.uri"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logrus.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("Connected to MongoDB!")
	database := client.Database(viper.GetString("mongodb.database"))
	return database
}
