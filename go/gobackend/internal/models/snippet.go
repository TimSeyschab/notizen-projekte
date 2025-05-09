package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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
	oid, err := bson.ObjectIDFromHex(hex)
	if err != nil {
		return Snippet{}, err
	}

	var result Snippet
	err = m.Collection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Snippet{}, ErrNoSnippet
		}
		return Snippet{}, err
	}

	return result, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	opts := options.Find().SetLimit(2)

	cursor, error := m.Collection.Find(context.TODO(), bson.M{}, opts)
	if error != nil {
		return nil, error
	}

	var results []Snippet
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, error
	}

	return results, nil
}
