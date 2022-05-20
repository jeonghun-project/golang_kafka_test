package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"kafka/data/mongoDB"
	"kafka/model/key"
	"time"
)

type KeyStore struct {
	collection *mongo.Collection
	ctx        context.Context
	close      context.CancelFunc
}

func (ks *KeyStore) setContext() {
	ks.ctx, ks.close = context.WithTimeout(context.Background(), mongoDB.ContextTimeoutSeconds*time.Second)
	ks.collection = mongoDB.Collection(mongoDB.CKeys)
}

func NewKeyStore() *KeyStore {
	return &KeyStore{}
}

func (ks KeyStore) GetOneBy(query bson.M) (*key.Key, error) {
	ks.setContext()
	defer ks.close()

	r := key.New()
	err := ks.collection.FindOne(ks.ctx, query).Decode(r)
	if err != nil {
		return nil, err
	}
	r.DecryptAll()

	return r, err
}

func (ks KeyStore) GetBy(by bson.M) ([]key.Key, error) {
	ks.setContext()
	defer ks.close()

	results := make([]key.Key, 0)

	keys := make([]key.Key, 0)
	cursor, err := ks.collection.Find(ks.ctx, by)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ks.ctx, &results)
	if err != nil {
		return nil, err
	}

	for i := range results {
		r := results[i]
		if r.Deleted {
			continue
		}
		r.DecryptAll()
		keys = append(keys, r)
	}

	return keys, err
}
