package database

import (
	"context"

	"crossing-api/config"
	"crossing-api/dao"
	"crossing-api/libs/log"
	"crossing-api/models"

	firebase "firebase.google.com/go"
	db "firebase.google.com/go/db"
	"google.golang.org/api/option"
)

//dbRef holds reference to the database
var dbRef *db.Ref

// Init initializes Firebase database with the parent database
func Init() {
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
	dao.InitClients(dbRef)
	models.InitClients(dbRef)
}

// GetDB returns the database reference
func GetDB() *db.Ref {
	return dbRef
}
