package main

import (
	"context"
	"flag"
	"gobackend/internal/models"
	"log/slog"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	mongodb_uri := flag.String("dburi", "mongodb://root:example@localhost:27017", "Mongo Database URI")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	client, err := openDB(*mongodb_uri)
	defer client.Disconnect(context.TODO())

	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{Collection: client.Database("snippetbox").Collection("snippets")},
	}

	logger.Info("starting server", "addr", *addr)
	err = http.ListenAndServe(*addr, app.routes())

	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(mongodb_uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}

	return client, err
}
