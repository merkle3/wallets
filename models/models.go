package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Wallet struct {
	ID         primitive.ObjectID  `bson:"_id,omitempty"`
	Name       string              `bson:"name"`
	Address    string              `bson:"address"`
	Derivation int                 `bson:"derivation"`
	CreatedAt  primitive.Timestamp `bson:"created_at"`
}

type Transaction struct {
	ID       primitive.ObjectID  `bson:"_id"`
	Address  string              `bson:"address"`
	SignedAt primitive.Timestamp `bson:"signed_at"`
}
