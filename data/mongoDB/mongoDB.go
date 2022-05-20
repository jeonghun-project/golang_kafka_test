package mongoDB

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"kafka/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CKeys  = "keys"
	DBName = "lot"
)

var client *mongo.Client

const ContextTimeoutSeconds = 15

func Init() {
	log.Infoln("MongoDB Initializing..")
	var err error

	ctx, cancelFunc := context.WithTimeout(context.Background(), ContextTimeoutSeconds*time.Second)
	defer cancelFunc()

	opts := &options.ClientOptions{}
	opts.SetCompressors([]string{"snappy", "zstd"})

	opts.SetReadPreference(readpref.Nearest())
	opts.ApplyURI(config.MongoURL)

	client, err = mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatalln("MongoDB connecting err", err)
	}
}

func GetClient() *mongo.Client {
	if client == nil {
		Init()
	}
	return client
}

func Collection(name string) *mongo.Collection {
	return GetClient().Database(DBName).Collection(name)
}
