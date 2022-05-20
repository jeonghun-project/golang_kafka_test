package key

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var store Store

type Store interface {
	Create(*Key) error
	GetBy(bson.M) ([]Key, error)
	GetOneBy(bson.M) (*Key, error)
}

func SetStore(s Store) {
	store = s
}

func GetByUserID(userID primitive.ObjectID) ([]Key, error) {
	return store.GetBy(bson.M{keyUserID: userID})
}

func GetByID(id primitive.ObjectID) (*Key, error) {
	return store.GetOneBy(bson.M{KeyID: id})
}

func DeleteByKeyId(keyId primitive.ObjectID) {
	fmt.Println("delete in MongoDB: ", keyId)
}
