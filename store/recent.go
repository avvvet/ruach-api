package store

import (
	"encoding/json"
	"time"

	bolt "go.etcd.io/bbolt"
)

type Transcription struct {
	ID             string    `json:"id"`
	Text           string    `json:"text"`
	Duration       float64   `json:"duration"`
	ProcessingTime float64   `json:"processing_time"`
	CreatedAt      time.Time `json:"created_at"`
}

// Add saves a transcription and keeps only the last 10
func Add(t Transcription) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketRecent))

		// get existing
		existing := getAll(b)

		// prepend new item
		items := append([]Transcription{t}, existing...)

		// trim to 10
		if len(items) > 10 {
			items = items[:10]
		}

		// save back
		data, err := json.Marshal(items)
		if err != nil {
			return err
		}

		return b.Put([]byte("list"), data)
	})
}

// GetAll returns all recent transcriptions
func GetAll() ([]Transcription, error) {
	var items []Transcription
	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketRecent))
		items = getAll(b)
		return nil
	})
	return items, err
}

func getAll(b *bolt.Bucket) []Transcription {
	var items []Transcription
	data := b.Get([]byte("list"))
	if data != nil {
		json.Unmarshal(data, &items)
	}
	return items
}
