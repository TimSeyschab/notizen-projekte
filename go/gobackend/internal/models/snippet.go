package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Snippet struct {
	ID      bson.ObjectID `bson:"_id,omitempty"`
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	Collection *mongo.Collection
}

func (m *SnippetModel) Insert(title string, content string, expires int) (string, error) {
	snippet := Snippet{ID: bson.NewObjectID(), Title: title, Content: content, Created: time.Now(), Expires: time.Now().Add(time.Duration(expires) * time.Hour)}
	result, err := m.Collection.InsertOne(context.TODO(), snippet)
	if err != nil {
		return "", err
	}

	oid, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return "", err
	}

	return oid.Hex(), nil
}

func (m *SnippetModel) Get(hex string) (Snippet, error) {
	return Snippet{}, nil
}
