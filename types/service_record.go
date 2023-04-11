package types

type ServiceRecord struct {
	ServiceRecordID        int     `json:"serviceRecordID"`
	SessionKey             string  `json:"sessionKey"`
	RelayID                string  `json:"relayID"`
	PortalRegionID         int     `json:"portalRegionID"`
	Latency                float64 `json:"latency"`
	Tickets                int     `json:"tickets"`
	Result                 string  `json:"result"`
	Available              bool    `json:"available"`
	Successes              int     `json:"successes"`
	Failures               int     `json:"failures"`
	P90SuccessLatency      float64 `json:"p90SuccessLatency"`
	MedianSuccessLatency   float64 `json:"medianSuccessLatency"`
	WeightedSuccessLatency float64 `json:"weightedSuccessLatency"`
	SuccessRate            float64 `json:"successRate"`
}
