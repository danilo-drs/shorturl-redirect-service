package model

import (
	"database/sql"
	"fmt"
	"meli-redirect-service/repository"
)

type ShortUrl struct {
	Key         string `json:"shortUrlKey"`
	OriginalURL string `json:"url"`
	ShortUrl    string `json:"shortUrl"`
	CreateAt    string `json:"createAt"`
	CreatedBy   string `json:"createdBy"`
	UpdateAt    string `json:"updateAt"`
	UpdatedBy   string `json:"updatedBy"`
}

// FillFromKey fills the short URL from the key
func (s *ShortUrl) FillFromKey() (bool, error) {
	query := "SELECT key, url, short_url, create_at, created_by, update_at, updated_by FROM short_url WHERE key = $1"
	err := repository.DB.QueryRow(query, s.Key).Scan(&s.Key, &s.OriginalURL, &s.ShortUrl, &s.CreateAt, &s.CreatedBy, &s.UpdateAt, &s.UpdatedBy)
	var found bool = false
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return found, nil
		}
		fmt.Println("Error getting short URL with key: ", err)
		return found, err
	}
	return true, nil
}
