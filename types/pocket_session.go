package types

import "time"

type PocketSession struct {
	SessionKey       string    `json:"sessionKey"`
	SessionHeight    int       `json:"sessionHeight"`
	PortalRegionName string    `json:"portalRegionName"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
