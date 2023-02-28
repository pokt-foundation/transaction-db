package postgresdriver

import (
	"context"
	"time"

	"github.com/pokt-foundation/transaction-db/types"
)

func (ts *PGDriverTestSuite) TestPostgresDriver_WriteRelay() {
	tests := []struct {
		name  string
		relay types.Relay
		err   error
	}{
		{
			name: "Success",
			relay: types.Relay{
				ChainID:                  21,
				EndpointID:               21,
				PocketSessionID:          ts.firstRelay.PocketSessionID,
				PoktNodeAddress:          "21",
				RelayStartDatetime:       time.Now(),
				RelayReturnDatetime:      time.Now(),
				IsError:                  true,
				ErrorID:                  ts.firstRelay.ErrorID,
				RelayRoundtripTime:       1,
				RelayChainMethodID:       21,
				RelayDataSize:            21,
				RelayPortalTripTime:      21,
				RelayNodeTripTime:        21,
				RelayURLIsPublicEndpoint: false,
				PortalOriginRegionID:     ts.firstRelay.PortalOriginRegionID,
				IsAltruistRelay:          false,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		ts.Equal(ts.driver.WriteRelay(context.Background(), tt.relay), tt.err)
	}
}

func (ts *PGDriverTestSuite) TestPostgresDriver_ReadRelay() {
	tests := []struct {
		name     string
		relayID  int
		expRelay types.Relay
		err      error
	}{
		{
			name:     "Success",
			relayID:  int(ts.firstRelay.RelayID),
			expRelay: ts.firstRelay,
			err:      nil,
		},
	}
	for _, tt := range tests {
		relay, err := ts.driver.ReadRelay(context.Background(), tt.relayID)
		ts.Equal(err, tt.err)
		ts.Equal(relay, tt.expRelay)
	}
}
