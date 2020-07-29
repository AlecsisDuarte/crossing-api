package dao

import (
	"context"

	"firebase.google.com/go/db"
)

const (
	portBucket           = "port"
	metadataBucket       = "metadata"
	geographicInfoBucket = "geographic_info"
	countriesBucket      = "countries"
	statesBucket         = "states"
	countiesBucket       = "counties"
	exchangeBucket       = "exchange"
)

var (
	portClient     *db.Ref
	metadataClient *db.Ref
	ctx            context.Context
)

// InitClients initializes all the required database clients
func InitClients(dbRef *db.Ref) {
	portClient = dbRef.Child(portBucket)
	metadataClient = dbRef.Child(metadataBucket)
	ctx = context.Background()
}
