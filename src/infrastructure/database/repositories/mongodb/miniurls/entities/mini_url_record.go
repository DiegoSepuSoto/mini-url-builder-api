package entities

type MiniURLRecord struct {
	OriginalURL string `bson:"original_url,omitempty"`
	NewURL      string `bson:"new_url,omitempty"`
}
