package key

import (
	"github.com/mervick/aes-everywhere/go/aes256"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kafka/config"
	"time"
)

const (
	keyUserID    = "user_id"
	keyExchange  = "exchange"
	KeyID        = "_id"
	KeyName      = "name"
	KeyAccessKey = "access_key"
	KeySecretKey = "secret_key"
)

type Key struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Exchange  string             `bson:"exchange" json:"exchange,omitempty"`
	AccessKey string             `bson:"access_key" json:"access_key,omitempty"`
	SecretKey string             `bson:"secret_key,omitempty" json:"secret_key,omitempty"`

	//For Okex, Bitget
	Passphrase     string     `bson:"pass_phrase,omitempty" json:"pass_phrase,omitempty"`
	RefreshToken   string     `bson:"refresh_token,omitempty" json:"-"`
	TokenRefreshAt *time.Time `bson:"token_refresh_at,omitempty" json:"token_refresh_at,omitempty"`

	//Used for FTX
	SubAccount  string `bson:"sub_account,omitempty" json:"sub_account,omitempty"`
	Hash        string `bson:"hash" json:"-"`
	IsFirstTime bool   `bson:"is_first_time,omitempty" json:"is_first_time,omitempty"`

	IsActive        bool   `bson:"is_active" json:"is_active"`
	CreatedAt       int64  `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt       int64  `bson:"updated_at" json:"-"`
	UpdatedStatusAt int64  `bson:"updated_status_at" json:"-"` //only when there's a status update
	AccountID       string `bson:"account_id" json:"account_id"`

	//For binance
	SubEnabled int    `bson:"sub_enabled,omitempty" json:"-"`
	Type       string `bson:"type,omitempty" json:"-"`
	MainID     string `bson:"main_id,omitempty" json:"-"`

	Deleted   bool  `bson:"deleted" json:"-"`
	DeletedAt int64 `bson:"deleted_at" json:"-"`
}

func New() *Key {
	return &Key{}
}

func GetDecrypted(encryptedText string) string {
	return aes256.Decrypt(encryptedText, config.EncryptionSecret)
}

//DecryptAll - decrypts all sensitive info, don't use anywhere else, the database already decrypts when you fetch in all methods
func (k *Key) DecryptAll() {
	if k.SecretKey != "" {
		k.SecretKey = GetDecrypted(k.SecretKey)
	}
	if k.Passphrase != "" {
		k.Passphrase = GetDecrypted(k.Passphrase)
	}
	if k.RefreshToken != "" {
		k.RefreshToken = GetDecrypted(k.RefreshToken)
	}
}
