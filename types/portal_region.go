package types

import "time"

type PortalRegion struct {
	PortalRegionID   int       `json:"portalRegionID"`
	PortalRegionName string    `json:"portalRegionName"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
