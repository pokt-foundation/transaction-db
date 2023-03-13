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
				SessionKey:               ts.firstRelay.SessionKey,
				PoktNodeAddress:          "21",
				RelayStartDatetime:       time.Now(),
				RelayReturnDatetime:      time.Now(),
				IsError:                  true,
				ErrorCode:                21,
				ErrorName:                "favorite number",
				ErrorMessage:             "just Pablo can use it",
				ErrorType:                types.ErrorTypeChainCheck,
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

func (ts *PGDriverTestSuite) TestPostgresDriver_WriteRelays() {
	var relays []types.Relay
	for i := 0; i < 1000; i++ {
		relays = append(relays, types.Relay{
			ChainID:                  21,
			EndpointID:               21,
			SessionKey:               ts.firstRelay.SessionKey,
			PoktNodeAddress:          "21",
			RelayStartDatetime:       time.Now(),
			RelayReturnDatetime:      time.Now(),
			IsError:                  true,
			ErrorCode:                21,
			ErrorName:                "favorite number",
			ErrorMessage:             "just Pablo can use it",
			ErrorType:                types.ErrorTypeChainCheck,
			RelayRoundtripTime:       1,
			RelayChainMethodID:       21,
			RelayDataSize:            21,
			RelayPortalTripTime:      21,
			RelayNodeTripTime:        21,
			RelayURLIsPublicEndpoint: false,
			PortalOriginRegionID:     ts.firstRelay.PortalOriginRegionID,
			IsAltruistRelay:          false,
		})
	}

	tests := []struct {
		name   string
		relays []types.Relay
		err    error
	}{
		{
			name:   "Success",
			relays: relays,
			err:    nil,
		},
	}
	for _, tt := range tests {
		ts.Equal(ts.driver.WriteRelays(context.Background(), tt.relays), tt.err)
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
