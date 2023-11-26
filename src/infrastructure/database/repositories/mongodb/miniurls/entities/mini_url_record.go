package entities

import "time"

type MiniURLRecord struct {
	OriginalURL string    `bson:"original_url,omitempty"`
	NewURL      string    `bson:"new_url,omitempty"`
	CreatedAt   time.Time `bson:"created_at,omitempty"`
}
