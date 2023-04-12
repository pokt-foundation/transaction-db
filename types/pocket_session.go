package types

import "time"

type PocketSession struct {
	SessionKey    string    `json:"sessionKey"`
	SessionHeight int       `json:"sessionHeight"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
