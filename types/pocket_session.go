package types

import "time"

type PocketSession struct {
	SessionKey            string    `json:"sessionKey"`
	SessionHeight         int       `json:"sessionHeight"`
	ProtocolApplicationID string    `json:"protocolApplicationID"`
	ProtocolPublicKey     string    `json:"protocolPublicKey"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}
