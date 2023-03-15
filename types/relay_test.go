package types

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRelay_ValidateStruct(t *testing.T) {
	c := require.New(t)

	tests := []struct {
		name  string
		relay Relay
		err   error
	}{
		{
			name: "Success no error relay",
			relay: Relay{
				ChainID:                  21,
				EndpointID:               21,
				SessionKey:               "21",
				PoktNodeAddress:          "21",
				RelayStartDatetime:       time.Now(),
				RelayReturnDatetime:      time.Now(),
				RelayRoundtripTime:       1,
				RelayChainMethodID:       21,
				RelayDataSize:            21,
				RelayPortalTripTime:      21,
				RelayNodeTripTime:        21,
				RelayURLIsPublicEndpoint: false,
				PortalOriginRegionID:     12,
				IsAltruistRelay:          false,
			},
			err: nil,
		},
		{
			name: "Success no error relay",
			relay: Relay{
				ChainID:                  21,
				EndpointID:               21,
				SessionKey:               "21",
				PoktNodeAddress:          "21",
				RelayStartDatetime:       time.Now(),
				RelayReturnDatetime:      time.Now(),
				IsError:                  true,
				ErrorCode:                21,
				ErrorName:                "favorite number",
				ErrorMessage:             "just Pablo can use it",
				ErrorType:                ErrorTypeChainCheck,
				RelayRoundtripTime:       1,
				RelayChainMethodID:       21,
				RelayDataSize:            21,
				RelayPortalTripTime:      21,
				RelayNodeTripTime:        21,
				RelayURLIsPublicEndpoint: false,
				PortalOriginRegionID:     12,
				IsAltruistRelay:          false,
			},
			err: nil,
		},
		{
			name: "Failure no erro relay with error fields",
			relay: Relay{
				ChainID:                  21,
				EndpointID:               21,
				SessionKey:               "21",
				PoktNodeAddress:          "21",
				RelayStartDatetime:       time.Now(),
				RelayReturnDatetime:      time.Now(),
				IsError:                  false,
				ErrorCode:                21,
				ErrorName:                "favorite number",
				ErrorMessage:             "just Pablo can use it",
				ErrorType:                ErrorTypeChainCheck,
				RelayRoundtripTime:       1,
				RelayChainMethodID:       21,
				RelayDataSize:            21,
				RelayPortalTripTime:      21,
				RelayNodeTripTime:        21,
				RelayURLIsPublicEndpoint: false,
				PortalOriginRegionID:     12,
				IsAltruistRelay:          false,
			},
			err: errors.New("ErrorCode should not be set"),
		},
		{
			name: "Failure invalid error type",
			relay: Relay{
				ChainID:                  21,
				EndpointID:               21,
				SessionKey:               "21",
				PoktNodeAddress:          "21",
				RelayStartDatetime:       time.Now(),
				RelayReturnDatetime:      time.Now(),
				IsError:                  true,
				ErrorCode:                21,
				ErrorName:                "favorite number",
				ErrorMessage:             "just Pablo can use it",
				ErrorType:                "pablito",
				RelayRoundtripTime:       1,
				RelayChainMethodID:       21,
				RelayDataSize:            21,
				RelayPortalTripTime:      21,
				RelayNodeTripTime:        21,
				RelayURLIsPublicEndpoint: false,
				PortalOriginRegionID:     12,
				IsAltruistRelay:          false,
			},
			err: errors.New("ErrorType is not valid"),
		},
		{
			name: "Failure error relay without error fields",
			relay: Relay{
				ChainID:                  21,
				EndpointID:               21,
				SessionKey:               "21",
				PoktNodeAddress:          "21",
				RelayStartDatetime:       time.Now(),
				RelayReturnDatetime:      time.Now(),
				IsError:                  true,
				ErrorName:                "favorite number",
				ErrorMessage:             "just Pablo can use it",
				ErrorType:                "pablito",
				RelayRoundtripTime:       1,
				RelayChainMethodID:       21,
				RelayDataSize:            21,
				RelayPortalTripTime:      21,
				RelayNodeTripTime:        21,
				RelayURLIsPublicEndpoint: false,
				PortalOriginRegionID:     12,
				IsAltruistRelay:          false,
			},
			err: errors.New("ErrorCode is not set"),
		},
		{
			name: "Failure not set field",
			relay: Relay{
				EndpointID:               21,
				SessionKey:               "21",
				PoktNodeAddress:          "21",
				RelayStartDatetime:       time.Now(),
				RelayReturnDatetime:      time.Now(),
				RelayRoundtripTime:       1,
				RelayChainMethodID:       21,
				RelayDataSize:            21,
				RelayPortalTripTime:      21,
				RelayNodeTripTime:        21,
				RelayURLIsPublicEndpoint: false,
				PortalOriginRegionID:     12,
				IsAltruistRelay:          false,
			},
			err: errors.New("ChainID is not set"),
		},
		{
			name: "Failure set field should not be set",
			relay: Relay{
				RelayID:                  21,
				EndpointID:               21,
				SessionKey:               "21",
				PoktNodeAddress:          "21",
				RelayStartDatetime:       time.Now(),
				RelayReturnDatetime:      time.Now(),
				RelayRoundtripTime:       1,
				RelayChainMethodID:       21,
				RelayDataSize:            21,
				RelayPortalTripTime:      21,
				RelayNodeTripTime:        21,
				RelayURLIsPublicEndpoint: false,
				PortalOriginRegionID:     12,
				IsAltruistRelay:          false,
			},
			err: errors.New("RelayID should not be set"),
		},
	}

	for _, tt := range tests {
		c.Equal(tt.err, tt.relay.Validate())
	}
}
