package libs

import (
	"context"

	"crossing-api/config"

	"crossing-api/libs/log"

	firebase "firebase.google.com/go"
	db "firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var dbRef *db.Ref

// InitDatabase initializes Firebase database with the parent database
func InitDatabase() *db.Ref {
	ctx := context.Background()
	dbConfig := config.BuildDBConfig()
	conf := &firebase.Config{
		DatabaseURL: dbConfig.DatabaseURL,
	}

	opt := option.WithCredentialsFile(dbConfig.ServiceAccountKeyPath)

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatal("Error initializing firebase app", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal("Error initializing database client", err)
	}

	dbRef = client.NewRef(dbConfig.DatabaseName)
	return dbRef
}
