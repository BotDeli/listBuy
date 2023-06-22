package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"listBuy/logger"
)

var ctx = context.TODO()

func ConnectDB(uri string) (*mongo.Client, *mongo.Collection) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	logger.CheckError(err)
	err = client.Connect(ctx)
	checkConnection(err)
	collection := client.Database("mydb").Collection("users")
	return client, collection
}

func checkConnection(err error) {
	if err == nil {
		logger.Logger.Info("Successful connect MongoDB!")
	} else {
		logger.Logger.Fatal("Dont connection MongoDB...")
	}
}

func DisconnectDB(client *mongo.Client) {
	err := client.Disconnect(ctx)
	logger.CheckWarning(err)
}
