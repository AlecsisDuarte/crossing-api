package dao

import (
	l "crossing-api/libs/log"
	m "crossing-api/models"
)

// UploadMetadata uploads metadata information to the database
func UploadMetadata(metadata *m.Metadata) (err error) {
	l.Info("Trying to upload metadata to the database")
	if err := metadataClient.Set(ctx, metadata); err != nil {
		l.Error("Error while uploading metadata information", err)
		return err
	}
	l.Info("Successfully uploaded metadata information")
	return nil
}
