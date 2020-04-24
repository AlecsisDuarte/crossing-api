package libs

import (
	"context"
	"log"

	"crossing-api/config"

	firebase "firebase.google.com/go"
	db "firebase.google.com/go/db"
	"google.golang.org/api/option"
)

///Reference towards the parent db
var DBref *db.Ref

// We initialize Firebase database with the parent database
func InitDatabase() *db.Ref {
	ctx := context.Background()
	dbConfig := config.BuildDBConfig()
	conf := &firebase.Config{
		DatabaseURL: dbConfig.DatabaseURL,
	}

	opt := option.WithCredentialsFile(dbConfig.ServiceAccountKeyPath)

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing firebase app:", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	DBref = client.NewRef(dbConfig.DatabaseName)
	return DBref
}
