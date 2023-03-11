package postgresdriver

import (
	"context"
	"testing"
	"time"

	"github.com/pokt-foundation/transaction-db/types"

	"github.com/stretchr/testify/suite"
)

const (
	connectionString = "postgres://postgres:pgpassword@localhost:5432/postgres?sslmode=disable" // pragma: allowlist secret
)

type PGDriverTestSuite struct {
	suite.Suite
	connectionString string
	driver           *PostgresDriver

	// Records inserted on setup for testing purposes
	firstRelay types.Relay
}

func Test_RunPGDriverSuite(t *testing.T) {
	testSuite := new(PGDriverTestSuite)
	testSuite.connectionString = connectionString

	suite.Run(t, testSuite)
}

// SetupSuite runs before each test suite run
func (ts *PGDriverTestSuite) SetupSuite() {
	ts.NoError(ts.initPostgresDriver())

	ts.NoError(ts.driver.WriteSession(context.Background(), types.PocketSession{
		SessionKey:            "22",
		SessionHeight:         22,
		ProtocolApplicationID: 22,
	}))

	ts.NoError(ts.driver.WriteRegion(context.Background(), types.PortalRegion{
		PortalRegionName: "La Colombia",
	}))

	ts.NoError(ts.driver.WriteRelay(context.Background(), types.Relay{
		ChainID:                  21,
		EndpointID:               21,
		SessionKey:               "22",
		PoktNodeAddress:          "21",
		RelayStartDatetime:       time.Date(199, time.July, 21, 0, 0, 0, 0, time.Local),
		RelayReturnDatetime:      time.Date(199, time.July, 21, 0, 0, 0, 0, time.Local),
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
		PortalOriginRegionID:     1,
		IsAltruistRelay:          false,
	}))

	firstRelay, err := ts.driver.ReadRelay(context.Background(), 1)
	ts.NoError(err)

	ts.firstRelay = firstRelay
}

// Initializes a real instance of the Postgres driver that connects to the test Postgres Docker container
func (ts *PGDriverTestSuite) initPostgresDriver() error {
	driver, err := NewPostgresDriver(ts.connectionString)
	if err != nil {
		return err
	}
	ts.driver = driver

	return nil
}
