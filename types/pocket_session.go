package types

type PocketSession struct {
	PocketSessionID       int32  `json:"pocketSessionID"`
	SessionKey            string `json:"sessionKey"`
	SessionHeight         int32  `json:"sessionHeight"`
	ProtocolApplicationID int32  `json:"protocolApplicationID"`
}
