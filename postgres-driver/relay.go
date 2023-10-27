package postgresdriver

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pokt-foundation/transaction-db/types"
)

const chainMethodIDSeparator = ","

func (d *PostgresDriver) WriteRelay(ctx context.Context, relay types.Relay) error {
	now := time.Now()

	return d.InsertRelay(ctx, InsertRelayParams{
		PoktChainID:              relay.PoktChainID,
		EndpointID:               relay.EndpointID,
		SessionKey:               relay.SessionKey,
		ProtocolAppPublicKey:     relay.ProtocolAppPublicKey,
		RelaySourceUrl:           newSQLNullString(relay.RelaySourceURL),
		PoktNodeAddress:          newSQLNullString(relay.PoktNodeAddress),
		PoktNodeDomain:           newSQLNullString(relay.PoktNodeDomain),
		PoktNodePublicKey:        newSQLNullString(relay.PoktNodePublicKey),
		RelayStartDatetime:       relay.RelayStartDatetime,
		RelayReturnDatetime:      relay.RelayReturnDatetime,
		IsError:                  relay.IsError,
		ErrorCode:                newSQLNullInt32(int32(relay.ErrorCode)),
		ErrorName:                newSQLNullString(relay.ErrorName),
		ErrorMessage:             newSQLNullString(relay.ErrorMessage),
		ErrorType:                newSQLNullString(relay.ErrorType),
		ErrorSource:              newSQLNullErrorSource(ErrorSourcesEnum(relay.ErrorSource)),
		RelayRoundtripTime:       relay.RelayRoundtripTime,
		RelayChainMethodIds:      strings.Join(relay.RelayChainMethodIDs, chainMethodIDSeparator),
		RelayDataSize:            int32(relay.RelayDataSize),
		RelayPortalTripTime:      relay.RelayPortalTripTime,
		RelayNodeTripTime:        relay.RelayNodeTripTime,
		RelayUrlIsPublicEndpoint: relay.RelayURLIsPublicEndpoint,
		PortalRegionName:         relay.PortalRegionName,
		IsAltruistRelay:          relay.IsAltruistRelay,
		RequestID:                relay.RequestID,
		PoktTxID:                 newSQLNullString(relay.PoktTxID),
		IsUserRelay:              relay.IsUserRelay,
		GigastakeAppID:           newSQLNullString(relay.GigastakeAppID),
		CreatedAt:                now,
		UpdatedAt:                now,
		BlockingPlugin:           newSQLNullString(relay.BlockingPlugin),
	})
}

// TODO: use CopyFrom postgres method to do batch inserts more efficiently.
// 	https://docs.sqlc.dev/en/stable/howto/insert.html#using-copyfrom
func (d *PostgresDriver) WriteRelays(ctx context.Context, relays []*types.Relay) error {
	var errors []error
	for _, relay := range relays {
		if err := d.WriteRelay(ctx, *relay); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("error writing relay batch: %d insert errors: %w", len(errors), errors[0])
	}

	return nil
}

func (d *PostgresDriver) ReadRelay(ctx context.Context, relayID int) (types.Relay, error) {
	relay, err := d.SelectRelay(ctx, int64(relayID))
	if err != nil {
		return types.Relay{}, err
	}

	return types.Relay{
		RelayID:                  int(relay.ID),
		PoktChainID:              relay.PoktChainID,
		EndpointID:               relay.EndpointID,
		SessionKey:               relay.SessionKey,
		ProtocolAppPublicKey:     relay.ProtocolAppPublicKey,
		RelaySourceURL:           relay.RelaySourceUrl.String,
		PoktNodeAddress:          relay.PoktNodeAddress.String,
		PoktNodeDomain:           relay.PoktNodeDomain.String,
		PoktNodePublicKey:        relay.PoktNodePublicKey.String,
		RelayStartDatetime:       relay.RelayStartDatetime,
		RelayReturnDatetime:      relay.RelayReturnDatetime,
		IsError:                  relay.IsError,
		ErrorCode:                int(relay.ErrorCode.Int32),
		ErrorName:                relay.ErrorName.String,
		ErrorMessage:             relay.ErrorMessage.String,
		ErrorType:                relay.ErrorType.String,
		ErrorSource:              types.ErrorSource(relay.ErrorSource.ErrorSourcesEnum),
		RelayRoundtripTime:       relay.RelayRoundtripTime,
		RelayChainMethodIDs:      strings.Split(relay.RelayChainMethodIds, ","),
		RelayDataSize:            int(relay.RelayDataSize),
		RelayPortalTripTime:      relay.RelayPortalTripTime,
		RelayNodeTripTime:        relay.RelayNodeTripTime,
		RelayURLIsPublicEndpoint: relay.RelayUrlIsPublicEndpoint,
		PortalRegionName:         relay.PortalRegionName,
		IsAltruistRelay:          relay.IsAltruistRelay,
		RequestID:                relay.RequestID,
		IsUserRelay:              relay.IsUserRelay,
		PoktTxID:                 relay.PoktTxID.String,
		GigastakeAppID:           relay.GigastakeAppID.String,
		CreatedAt:                relay.CreatedAt,
		UpdatedAt:                relay.UpdatedAt,
		Session: types.PocketSession{
			SessionKey:    relay.SessionKey,
			SessionHeight: int(relay.SessionHeight),
			CreatedAt:     relay.CreatedAt_2,
			UpdatedAt:     relay.UpdatedAt_2,
		},
		Region: types.PortalRegion{
			PortalRegionName: relay.PortalRegionName,
		},
	}, nil
}
