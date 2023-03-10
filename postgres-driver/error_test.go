package postgresdriver

import (
	"context"

	"github.com/pokt-foundation/transaction-db/types"
)

func (ts *PGDriverTestSuite) TestPostgresDriver_WriteError() {
	tests := []struct {
		name      string
		poktError types.Error
		err       error
	}{
		{
			name: "Success",
			poktError: types.Error{
				ErrorCode:        404,
				ErrorName:        "not found",
				ErrorDescription: "I guess we didn't find it",
				ErrorType:        types.ErrorTypeRelay,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		ts.Equal(ts.driver.WriteError(context.Background(), tt.poktError), tt.err)
	}
}
