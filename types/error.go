package types

type Error struct {
	ErrorID          int32  `json:"errorID"`
	ErrorCode        int32  `json:"errorCode"`
	ErrorName        string `json:"errorName"`
	ErrorDescription string `json:"errorDescription"`
}
