package database

import (
	"context"

	"crossing-api/config"
	dao "crossing-api/dao"
	l "crossing-api/libs/log"

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
		l.Fatal("Error initializing firebase app", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		l.Fatal("Error initializing database client", err)
	}

	dbRef = client.NewRef(dbConfig.DatabaseName)
	dao.InitClients(dbRef)
}

// GetDB returns the database reference
func GetDB() *db.Ref {
	return dbRef
}
