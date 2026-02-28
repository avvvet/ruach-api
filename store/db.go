package store

import (
	"log"

	bolt "go.etcd.io/bbolt"
)

var DB *bolt.DB

const BucketRecent = "recent"

func Init(path string) {
	var err error
	DB, err = bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatalf("store: failed to open db: %v", err)
	}

	// create bucket if not exists
	err = DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BucketRecent))
		return err
	})
	if err != nil {
		log.Fatalf("store: failed to create bucket: %v", err)
	}

	log.Println("store: db initialized at", path)
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
